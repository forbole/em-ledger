// This software is Copyright (c) 2019-2020 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	supply "github.com/cosmos/cosmos-sdk/x/supply/exported"
	embank "github.com/e-money/em-ledger/hooks/bank"

	"github.com/e-money/em-ledger/types"
)

type (
	// TODO This could be unnecessary now. Remove.
	AccountKeeper interface {
		GetAccount(sdk.Context, sdk.AccAddress) auth.Account
	}

	BankKeeper interface {
		InputOutputCoins(ctx sdk.Context, inputs []bank.Input, outputs []bank.Output) error
		SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

		AddBalanceListener(l embank.BalanceListener)
	}

	SupplyKeeper interface {
		GetSupply(ctx sdk.Context) (supply supply.SupplyI)
	}

	RestrictedKeeper interface {
		GetRestrictedDenoms(sdk.Context) types.RestrictedDenoms
	}
)
