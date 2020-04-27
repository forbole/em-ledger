package keeper

import (
	"bytes"
	"encoding/gob"
	db "github.com/tendermint/tm-db"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/e-money/em-ledger/x/slashing/types"
)

// GetValidatorSigningInfo retruns the ValidatorSigningInfo for a specific validator
// ConsAddress
func (k Keeper) GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (info types.ValidatorSigningInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorSigningInfoKey(address))
	if bz == nil {
		found = false
		return
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &info)
	found = true
	return
}

// HasValidatorSigningInfo returns if a given validator has signing information
// persited.
func (k Keeper) HasValidatorSigningInfo(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	_, ok := k.GetValidatorSigningInfo(ctx, consAddr)
	return ok
}

// SetValidatorSigningInfo sets the validator signing info to a consensus address key
func (k Keeper) SetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress, info types.ValidatorSigningInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&info)
	store.Set(types.GetValidatorSigningInfoKey(address), bz)
}

// IterateValidatorSigningInfos iterates over the stored ValidatorSigningInfo
func (k Keeper) IterateValidatorSigningInfos(ctx sdk.Context,
	handler func(address sdk.ConsAddress, info types.ValidatorSigningInfo) (stop bool)) {

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorSigningInfoKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		address := types.GetValidatorSigningInfoAddress(iter.Key())
		var info types.ValidatorSigningInfo
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &info)
		if handler(address, info) {
			break
		}
	}
}

// GetValidatorMissedBlockBitArray gets the bit for the missed blocks array
func (k Keeper) GetValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress, index int64) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorMissedBlockBitArrayKey(address, index))
	var missed gogotypes.BoolValue
	if bz == nil {
		// lazy: treat empty key as not missed
		return false
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &missed)

	return missed.Value
}

// IterateValidatorMissedBlockBitArray iterates over the signed blocks window
// and performs a callback function
//func (k Keeper) IterateValidatorMissedBlockBitArray(ctx sdk.Context,
//	address sdk.ConsAddress, handler func(index int64, missed bool) (stop bool)) {
//
//	store := ctx.KVStore(k.storeKey)
//	index := int64(0)
//	// Array may be sparse
//	for ; index < k.SignedBlocksWindow(ctx); index++ {
//		var missed gogotypes.BoolValue
//		bz := store.Get(types.GetValidatorMissedBlockBitArrayKey(address, index))
//		if bz == nil {
//			continue
//		}
//
//		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &missed)
//		if handler(index, missed.Value) {
//			break
//		}
//	}
//}

// JailUntil attempts to set a validator's JailedUntil attribute in its signing
// info. It will panic if the signing info does not exist for the validator.
func (k Keeper) JailUntil(ctx sdk.Context, consAddr sdk.ConsAddress, jailTime time.Time) {
	signInfo, ok := k.GetValidatorSigningInfo(ctx, consAddr)
	if !ok {
		panic("cannot jail validator that does not have any signing information")
	}

	signInfo.JailedUntil = jailTime
	k.SetValidatorSigningInfo(ctx, consAddr, signInfo)
}

// Tombstone attempts to tombstone a validator. It will panic if signing info for
// the given validator does not exist.
func (k Keeper) Tombstone(ctx sdk.Context, consAddr sdk.ConsAddress) {
	signInfo, ok := k.GetValidatorSigningInfo(ctx, consAddr)
	if !ok {
		panic("cannot tombstone validator that does not have any signing information")
	}

	if signInfo.Tombstoned {
		panic("cannot tombstone validator that is already tombstoned")
	}

	signInfo.Tombstoned = true
	k.SetValidatorSigningInfo(ctx, consAddr, signInfo)
}

// IsTombstoned returns if a given validator by consensus address is tombstoned.
func (k Keeper) IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	signInfo, ok := k.GetValidatorSigningInfo(ctx, consAddr)
	if !ok {
		return false
	}

	return signInfo.Tombstoned
}

// SetValidatorMissedBlockBitArray sets the bit that checks if the validator has
// missed a block in the current window
func (k Keeper) SetValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress, index int64, missed bool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&gogotypes.BoolValue{Value: missed})
	store.Set(types.GetValidatorMissedBlockBitArrayKey(address, index), bz)
}

// clearValidatorMissedBlockBitArray deletes every instance of ValidatorMissedBlockBitArray in the store
func (k Keeper) clearValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetValidatorMissedBlockBitArrayPrefixKey(address))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

func (k Keeper) HandlePendingPenalties(ctx sdk.Context, batch db.Batch, vfn func() map[string]bool) {
	activePenalties := k.getPendingPenalties()
	if len(activePenalties) == 0 {
		return
	}

	validatorSet := vfn()
	for val, coins := range activePenalties {
		if _, present := validatorSet[val]; present {
			// Penalized validator is still in the validator set. Do not pay out slashing fine.
			continue
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypePenaltyPayout,
				sdk.NewAttribute(types.AttributeKeyAmount, coins.String()),
				sdk.NewAttribute(types.AttributeKeyAddress, val),
			),
		)

		delete(activePenalties, val)

		err := k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.PenaltyAccount, k.feeModuleName, coins)
		if err != nil {
			panic(err)
		}
	}

	k.setPendingPenalties(batch, activePenalties)
}

func (k Keeper) setPendingPenalties(batch db.Batch, penalties map[string]sdk.Coins) {
	if len(penalties) == 0 {
		batch.Delete([]byte(dbKeyPendingPenalties))
		return
	}

	bz := k.cdc.MustMarshalJSON(penalties)
	batch.Set([]byte(dbKeyPendingPenalties), bz)
}

func (k Keeper) getPendingPenalties() map[string]sdk.Coins {
	bz, err := k.database.Get([]byte(dbKeyPendingPenalties))
	if err != nil {
		panic(err) // TODO Find a better way to handle this?
	}

	if len(bz) == 0 {
		return make(map[string]sdk.Coins)
	}

	activePenalties := make(map[string]sdk.Coins)
	k.cdc.MustUnmarshalJSON(bz, &activePenalties)

	return activePenalties
}

func (k Keeper) GetBlockTimes() []time.Time {
	bz, err := k.database.Get([]byte(dbKeyBlockTimes))
	if err != nil {
		panic(err) // TODO Find a better way to handle this?
	}

	if len(bz) == 0 {
		return make([]time.Time, 0)
	}

	b := bytes.NewBuffer(bz)
	blockTimes := make([]time.Time, 0)
	dec := gob.NewDecoder(b)
	_ = dec.Decode(&blockTimes)
	return blockTimes
}

func (k Keeper) SetBlockTimes(batch db.Batch, blockTimes []time.Time) {
	bz := new(bytes.Buffer)
	enc := gob.NewEncoder(bz)
	_ = enc.Encode(blockTimes)
	batch.Set([]byte(dbKeyBlockTimes), bz.Bytes())
}
