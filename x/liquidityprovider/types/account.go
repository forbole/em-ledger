// This software is Copyright (c) 2019 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

var _ exported.Account = new(LiquidityProviderAccount)

//type LiquidityProviderAccount struct {
//	auth.BaseAccount
//
//	Mintable sdk.Coins `json:"mintable" yaml:"mintable"`
//}

func NewLiquidityProviderAccount(baseAccount *auth.BaseAccount, mintable sdk.Coins) *LiquidityProviderAccount {
	return &LiquidityProviderAccount{
		BaseAccount: baseAccount,
		Mintable:    mintable,
	}
}

func (acc *LiquidityProviderAccount) IncreaseMintableAmount(increase sdk.Coins) {
	// TODO Get protobuf to declare as Coins: https://github.com/gogo/protobuf/pull/658
	acc.Mintable = acc.Mintable.Add(increase...)
}

// Function panics if resulting mintable amount is negative. Should be checked prior to invocation for cleaner handling.
func (acc *LiquidityProviderAccount) DecreaseMintableAmount(decrease sdk.Coins) {
	if mintable, anyNegative := sdk.Coins(acc.Mintable).SafeSub(decrease); !anyNegative {
		acc.Mintable = mintable
		return
	}

	panic(fmt.Errorf("mintable amount cannot be negative"))
}

func (acc LiquidityProviderAccount) String() string {
	var pubkey string

	if acc.GetPubKey() != nil {
		pubkey = sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, acc.GetPubKey())
	}

	return fmt.Sprintf(`Account:
  Address:       %s
  Pubkey:        %s
  Mintable:      %s
  AccountNumber: %d
  Sequence:      %d`,
		acc.GetAddress(), pubkey, acc.Mintable, acc.GetAccountNumber(), acc.GetSequence(),
	)
}
