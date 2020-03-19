// This software is Copyright (c) 2019 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package slashing

import (
	db "github.com/tendermint/tm-db"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func TestCannotUnjailUnlessJailed(t *testing.T) {
	// initial setup
	ctx, ck, sk, _, keeper, _, ak := createTestInput(t, DefaultParams(), db.NewMemDB())
	slh := NewHandler(keeper)
	amt := sdk.TokensFromConsensusPower(100)
	addr, val := addrs[0], pks[0]
	ak.SetAccount(ctx, ak.NewAccountWithAddress(ctx, sdk.AccAddress(addr)))

	msg := NewTestMsgCreateValidator(addr, val, amt)
	_, err := staking.NewHandler(sk)(ctx, msg)
	require.NoError(t, err, "%v", err)
	staking.EndBlocker(ctx, sk)

	require.Equal(
		t, ck.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.Coins{sdk.NewCoin(sk.GetParams(ctx).BondDenom, initTokens.Sub(amt))},
	)
	require.Equal(t, amt, sk.Validator(ctx, addr).GetBondedTokens())

	// assert non-jailed validator can't be unjailed
	_, err = slh(ctx, NewMsgUnjail(addr))
	require.Error(t, err, "allowed unjail of non-jailed validator")
}

func TestCannotUnjailUnlessMeetMinSelfDelegation(t *testing.T) {
	// initial setup
	ctx, ck, sk, _, keeper, _, ak := createTestInput(t, DefaultParams(), db.NewMemDB())
	slh := NewHandler(keeper)
	amtInt := int64(100)
	addr, val, amt := addrs[0], pks[0], sdk.TokensFromConsensusPower(amtInt)
	ak.SetAccount(ctx, ak.NewAccountWithAddress(ctx, sdk.AccAddress(addr)))

	msg := NewTestMsgCreateValidator(addr, val, amt)
	msg.MinSelfDelegation = amt
	_, err := staking.NewHandler(sk)(ctx, msg)
	require.NoError(t, err)
	staking.EndBlocker(ctx, sk)

	require.Equal(
		t, ck.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.Coins{sdk.NewCoin(sk.GetParams(ctx).BondDenom, initTokens.Sub(amt))},
	)

	unbondAmt := sdk.NewCoin(sk.GetParams(ctx).BondDenom, sdk.OneInt())
	undelegateMsg := staking.NewMsgUndelegate(sdk.AccAddress(addr), addr, unbondAmt)
	_, err = staking.NewHandler(sk)(ctx, undelegateMsg)

	require.True(t, sk.Validator(ctx, addr).IsJailed())

	// assert non-jailed validator can't be unjailed
	_, err = slh(ctx, NewMsgUnjail(addr))
	require.Error(t, err, "allowed unjail of validator with less than MinSelfDelegation")
	//require.EqualValues(t, CodeValidatorNotJailed, got.Code)
	//require.EqualValues(t, DefaultCodespace, got.Codespace)
}

func TestJailedValidatorDelegations(t *testing.T) {
	ctx, _, stakingKeeper, _, slashingKeeper, _, ak := createTestInput(t, DefaultParams(), db.NewMemDB())

	stakingParams := stakingKeeper.GetParams(ctx)
	stakingParams.UnbondingTime = 1
	stakingKeeper.SetParams(ctx, stakingParams)

	// create a validator
	bondAmount := sdk.TokensFromConsensusPower(10)
	valPubKey := pks[0]
	valAddr, consAddr := addrs[1], sdk.ConsAddress(addrs[0])
	ak.SetAccount(ctx, ak.NewAccountWithAddress(ctx, sdk.AccAddress(valAddr)))

	msgCreateVal := NewTestMsgCreateValidator(valAddr, valPubKey, bondAmount)
	_, err := staking.NewHandler(stakingKeeper)(ctx, msgCreateVal)
	require.NoError(t, err, "expected create validator msg to be ok, got: %v", err)

	// end block
	staking.EndBlocker(ctx, stakingKeeper)

	// set dummy signing info
	newInfo := NewValidatorSigningInfo(consAddr, time.Unix(0, 0), false)
	slashingKeeper.SetValidatorSigningInfo(ctx, consAddr, newInfo)

	// delegate tokens to the validator
	delAddr := sdk.AccAddress(addrs[2])
	ak.SetAccount(ctx, ak.NewAccountWithAddress(ctx, delAddr))
	msgDelegate := newTestMsgDelegate(delAddr, valAddr, bondAmount)
	_, err = staking.NewHandler(stakingKeeper)(ctx, msgDelegate)
	require.NoError(t, err, "expected delegation to be ok, got %v", err)

	unbondAmt := sdk.NewCoin(stakingKeeper.GetParams(ctx).BondDenom, bondAmount)

	// unbond validator total self-delegations (which should jail the validator)
	msgUndelegate := staking.NewMsgUndelegate(sdk.AccAddress(valAddr), valAddr, unbondAmt)
	_, err = staking.NewHandler(stakingKeeper)(ctx, msgUndelegate)
	require.NoError(t, err, "expected begin unbonding validator msg to be ok, got: %v", err)

	err = stakingKeeper.CompleteUnbonding(ctx, sdk.AccAddress(valAddr), valAddr)
	require.Nil(t, err, "expected complete unbonding validator to be ok, got: %v", err)

	// verify validator still exists and is jailed
	validator, found := stakingKeeper.GetValidator(ctx, valAddr)
	require.True(t, found)
	require.True(t, validator.IsJailed())

	// verify the validator cannot unjail itself
	_, err = NewHandler(slashingKeeper)(ctx, NewMsgUnjail(valAddr))
	require.Error(t, err, "expected jailed validator to not be able to unjail.")

	// self-delegate to validator
	msgSelfDelegate := newTestMsgDelegate(sdk.AccAddress(valAddr), valAddr, bondAmount)
	_, err = staking.NewHandler(stakingKeeper)(ctx, msgSelfDelegate)
	require.NoError(t, err, "expected delegation to not be ok, got %v", err)

	// verify the validator can now unjail itself
	_, err = NewHandler(slashingKeeper)(ctx, NewMsgUnjail(valAddr))
	require.NoError(t, err, "expected jailed validator to be able to unjail, got: %v", err)
}

func TestInvalidMsg(t *testing.T) {
	k := Keeper{}
	h := NewHandler(k)

	_, err := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), "unrecognized slashing message type"))
}
