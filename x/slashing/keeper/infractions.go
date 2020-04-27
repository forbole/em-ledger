package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/e-money/em-ledger/x/slashing/types"
	"github.com/tendermint/tendermint/crypto"
	db "github.com/tendermint/tm-db"
	"sort"
	"time"
)

func (k Keeper) HandleValidatorSignature(ctx sdk.Context, batch db.Batch, addr crypto.Address, power int64, signed bool, blockCount int64) {
	logger := k.Logger(ctx)
	height := ctx.BlockHeight()
	consAddr := sdk.ConsAddress(addr)

	missedBlocks := k.getMissingBlocksForValidator(consAddr)
	missedBlocks = truncateByWindow(ctx.BlockTime(), missedBlocks, k.SignedBlocksWindowDuration(ctx))

	if !signed {
		missedBlocks = append(missedBlocks, ctx.BlockTime())
	}

	k.setMissingBlocksForValidator(batch, consAddr, missedBlocks)
	missedBlockCount := sdk.NewInt(int64(len(missedBlocks))).ToDec()

	missedRatio := missedBlockCount.QuoInt64(blockCount)

	// TODO Only do this if missing is true?
	minSignedPerWindow := k.MinSignedPerWindow(ctx)

	// if we are past the minimum height and the validator has missed too many blocks, punish them
	if sdk.OneDec().Sub(minSignedPerWindow).LT(missedRatio) {
		validator := k.sk.ValidatorByConsAddr(ctx, consAddr)
		if validator != nil && !validator.IsJailed() {

			// Downtime confirmed: slash and jail the validator
			logger.Info(fmt.Sprintf("Validator %s is below signed blocks threshold of %d during the last %d",
				consAddr, k.MinSignedPerWindow(ctx), k.SignedBlocksWindowDuration(ctx)))

			// We need to retrieve the stake distribution which signed the block, so we subtract ValidatorUpdateDelay from the evidence height,
			// and subtract an additional 1 since this is the LastCommit.
			// Note that this *can* result in a negative "distributionHeight" up to -ValidatorUpdateDelay-1,
			// i.e. at the end of the pre-genesis block (none) = at the beginning of the genesis block.
			// That's fine since this is just used to filter unbonding delegations & redelegations.
			distributionHeight := height - sdk.ValidatorUpdateDelay - 1

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeSlash,
					sdk.NewAttribute(types.AttributeKeyAddress, consAddr.String()),
					sdk.NewAttribute(types.AttributeKeyPower, fmt.Sprintf("%d", power)),
					sdk.NewAttribute(types.AttributeKeyReason, types.AttributeValueMissingSignature),
					sdk.NewAttribute(types.AttributeKeyJailed, consAddr.String()),
				),
			)

			k.slashValidator(ctx, batch, consAddr, distributionHeight, power, k.SlashFractionDowntime(ctx))
			k.sk.Jail(ctx, consAddr)

			// fetch signing info
			signInfo, found := k.GetValidatorSigningInfo(ctx, consAddr)
			if !found {
				panic(fmt.Sprintf("Expected signing info for validator %s but not found", consAddr))
			}

			signInfo.JailedUntil = ctx.BlockHeader().Time.Add(k.DowntimeJailDuration(ctx))

			// Reset number of blocks missed.
			k.deleteMissingBlocksForValidator(batch, consAddr)
			k.SetValidatorSigningInfo(ctx, consAddr, signInfo)
		} else {
			// Validator was (a) not found or (b) already jailed, don't slash
			logger.Info(
				fmt.Sprintf("Validator %s would have been slashed for downtime, but was either not found in store or already jailed", consAddr),
			)
		}
	}
}

func truncateByWindow(blockTime time.Time, times []time.Time, signedBlocksWindow time.Duration) []time.Time {
	if len(times) == 0 {
		return times
	}

	// Remove timestamps outside of the time window we are watching
	threshold := blockTime.Add(-1 * signedBlocksWindow)

	index := sort.Search(len(times), func(i int) bool {
		return times[i].After(threshold)
	})

	return times[index:]
}
