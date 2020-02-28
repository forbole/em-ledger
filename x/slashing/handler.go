// This software is Copyright (c) 2019 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package slashing

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/e-money/em-ledger/x/slashing/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgUnjail:
			return handleMsgUnjail(ctx, msg, k)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized slashing message type: %T", msg)

			//errMsg := fmt.Sprintf("unrecognized slashing message type: %T", msg)
			//return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Validators must submit a transaction to unjail itself after
// having been jailed (and thus unbonded) for downtime
func handleMsgUnjail(ctx sdk.Context, msg MsgUnjail, k Keeper) (*sdk.Result, error) {
	validator := k.sk.Validator(ctx, msg.ValidatorAddr)
	if validator == nil {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorForAddress, msg.ValidatorAddr.String())
		//return ErrNoValidatorForAddress(k.codespace).Result()
	}

	// cannot be unjailed if no self-delegation exists
	selfDel := k.sk.Delegation(ctx, sdk.AccAddress(msg.ValidatorAddr), msg.ValidatorAddr)
	if selfDel == nil {
		return nil, sdkerrors.Wrap(types.ErrMissingSelfDelegation, msg.ValidatorAddr.String())
		//return ErrMissingSelfDelegation(k.codespace).Result()
	}

	if validator.TokensFromShares(selfDel.GetShares()).TruncateInt().LT(validator.GetMinSelfDelegation()) {
		return nil, sdkerrors.Wrap(types.ErrSelfDelegationTooLowToUnjail, validator.GetMinSelfDelegation().String())
		//return ErrSelfDelegationTooLowToUnjail(k.codespace).Result()
	}

	// cannot be unjailed if not jailed
	if !validator.IsJailed() {
		return nil, sdkerrors.Wrap(types.ErrValidatorNotJailed, validator.GetConsAddr().String())
		//return ErrValidatorNotJailed(k.codespace).Result()
	}

	consAddr := sdk.ConsAddress(validator.GetConsPubKey().Address())

	info, found := k.getValidatorSigningInfo(ctx, consAddr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorForAddress, consAddr.String())
		//return ErrNoValidatorForAddress(k.codespace).Result()
	}

	// cannot be unjailed if tombstoned
	if info.Tombstoned {
		return nil, sdkerrors.Wrap(types.ErrValidatorJailed, "validator is tombstoned")
		//return ErrValidatorJailed(k.codespace).Result()
	}

	// cannot be unjailed until out of jail
	if ctx.BlockHeader().Time.Before(info.JailedUntil) {
		return nil, sdkerrors.Wrap(types.ErrValidatorJailed, info.JailedUntil.String())
		//return ErrValidatorJailed(k.codespace).Result()
	}

	k.sk.Unjail(ctx, consAddr)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
