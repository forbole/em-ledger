package keeper

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/params"
	db "github.com/tendermint/tm-db"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/e-money/em-ledger/x/slashing/types"
)

// Keeper of the slashing store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.Marshaler
	sk         types.StakingKeeper
	paramspace types.ParamSubspace

	feeModuleName string
	supplyKeeper  types.SupplyKeeper

	// Alternative to IAVL KV storage. For data that should not be part of consensus.
	database types.ReadOnlyDB
}

const (
	dbKeyMissedByVal      = "%v.missedBlocks"
	dbKeyPendingPenalties = "activePenalties"
	dbKeyBlockTimes       = "blocktimes"
)

// NewKeeper creates a slashing keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, sk types.StakingKeeper, supplyKeeper types.SupplyKeeper, feeModuleName string, paramspace params.Subspace, database db.DB) Keeper {
	// set KeyTable if it has not already been set
	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		sk:            sk,
		paramspace:    paramspace,
		feeModuleName: feeModuleName,
		supplyKeeper:  supplyKeeper,
		database:      database,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// AddPubkey sets a address-pubkey relation
func (k Keeper) AddPubkey(ctx sdk.Context, pubkey crypto.PubKey) {
	addr := pubkey.Address()

	pkStr, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pubkey)
	if err != nil {
		panic(fmt.Errorf("error while setting address-pubkey relation: %s", addr))
	}

	k.setAddrPubkeyRelation(ctx, addr, pkStr)
}

// GetPubkey returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetPubkey(ctx sdk.Context, address crypto.Address) (crypto.PubKey, error) {
	store := ctx.KVStore(k.storeKey)

	var pubkey gogotypes.StringValue
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(types.GetAddrPubkeyRelationKey(address)), &pubkey)
	if err != nil {
		return nil, fmt.Errorf("address %s not found", sdk.ConsAddress(address))
	}

	pkStr, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, pubkey.Value)
	if err != nil {
		return pkStr, err
	}

	return pkStr, nil
}

// Slash attempts to slash a validator. The slash is delegated to the staking
// module to make the necessary validator changes.
func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, fraction sdk.Dec, power, distributionHeight int64) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSlash,
			sdk.NewAttribute(types.AttributeKeyAddress, consAddr.String()),
			sdk.NewAttribute(types.AttributeKeyPower, fmt.Sprintf("%d", power)),
			sdk.NewAttribute(types.AttributeKeyReason, types.AttributeValueDoubleSign),
		),
	)

	k.sk.Slash(ctx, consAddr, distributionHeight, power, fraction)
}

// Jail attempts to jail a validator. The slash is delegated to the staking module
// to make the necessary validator changes.
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSlash,
			sdk.NewAttribute(types.AttributeKeyJailed, consAddr.String()),
		),
	)

	k.sk.Jail(ctx, consAddr)
}

func (k Keeper) setAddrPubkeyRelation(ctx sdk.Context, addr crypto.Address, pubkey string) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&gogotypes.StringValue{Value: pubkey})
	store.Set(types.GetAddrPubkeyRelationKey(addr), bz)
}

func (k Keeper) deleteAddrPubkeyRelation(ctx sdk.Context, addr crypto.Address) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAddrPubkeyRelationKey(addr))
}

func (k Keeper) deleteMissingBlocksForValidator(batch db.Batch, address sdk.ConsAddress) {
	key := []byte(fmt.Sprintf(dbKeyMissedByVal, address.String()))
	batch.Delete(key)
}

func (k Keeper) slashValidator(ctx sdk.Context, batch db.Batch, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec) {
	k.sk.Slash(ctx, consAddr, infractionHeight, power, slashFactor)

	// Mint the slashed coins and assign them to the distribution pool.
	slashAmount := calculateSlashingAmount(power, slashFactor)
	stakingDenom := k.sk.BondDenom(ctx)
	coins := sdk.NewCoins(sdk.NewCoin(stakingDenom, slashAmount))
	err := k.supplyKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		panic(err)
	}

	err = k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.PenaltyAccount, coins)
	if err != nil {
		panic(err)
	}

	activePenalties := k.getPendingPenalties()
	activePenalties[consAddr.String()] = coins

	k.setPendingPenalties(batch, activePenalties)
}

// Adopted f6rom cosmos-sdk/x/staking/keeper/slash.go
func calculateSlashingAmount(power int64, slashFactor sdk.Dec) sdk.Int {
	amount := sdk.TokensFromConsensusPower(power)
	slashAmountDec := amount.ToDec().Mul(slashFactor)
	slashAmount := slashAmountDec.TruncateInt()
	return slashAmount
}

func (k Keeper) getMissingBlocksForValidator(address sdk.ConsAddress) []time.Time {
	key := []byte(fmt.Sprintf(dbKeyMissedByVal, address.String()))
	bz, err := k.database.Get(key)
	if err != nil {
		panic(err) // TODO Better handling
	}

	if len(bz) == 0 {
		return nil
	}

	b := bytes.NewBuffer(bz)
	dec := gob.NewDecoder(b)

	res := make([]time.Time, 0)
	err = dec.Decode(&res)
	if err != nil {
		panic(err)
	}

	return res
}

func (k Keeper) setMissingBlocksForValidator(batch db.Batch, address sdk.ConsAddress, missingBlocks []time.Time) {
	bz := new(bytes.Buffer)
	enc := gob.NewEncoder(bz)
	err := enc.Encode(missingBlocks)
	if err != nil {
		panic(err)
	}

	key := []byte(fmt.Sprintf(dbKeyMissedByVal, address.String()))
	batch.Set(key, bz.Bytes())
}

// Downtime slashing threshold
func (k Keeper) MinSignedPerWindow(ctx sdk.Context) (res sdk.Dec) {
	k.paramspace.Get(ctx, types.KeyMinSignedPerWindow, &res)
	return
}
