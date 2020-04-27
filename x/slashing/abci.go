package slashing

import (
	abci "github.com/tendermint/tendermint/abci/types"
	db "github.com/tendermint/tm-db"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, sk Keeper, batch db.Batch) {
	signedBlocksWindow := sk.SignedBlocksWindowDuration(ctx)

	blockTimes := sk.GetBlockTimes()
	blockTimes = append(blockTimes, ctx.BlockTime())
	blockTimes = truncateByWindow(ctx.BlockTime(), blockTimes, signedBlocksWindow)
	sk.SetBlockTimes(batch, blockTimes)

	sk.HandlePendingPenalties(ctx, batch, validatorset(req.LastCommitInfo.Votes))

	// Iterate over all the validators which *should* have signed this block
	// store whether or not they have actually signed it and slash/unbond any
	// which have missed too many blocks in a row (downtime slashing)
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		sk.HandleValidatorSignature(ctx, batch, voteInfo.Validator.Address, voteInfo.Validator.Power, voteInfo.SignedLastBlock, int64(len(blockTimes)))
	}
}

// Make a set containing all validators that are part of the set
func validatorset(validators []abci.VoteInfo) func() map[string]bool {
	return func() map[string]bool {
		res := make(map[string]bool)
		for _, v := range validators {
			res[sdk.ConsAddress(v.Validator.Address).String()] = true
		}

		return res
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
