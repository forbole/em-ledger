// This software is Copyright (c) 2019-2020 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

var _ bank.SendKeeper = (*ProxyKeeper)(nil)

type BalanceListener func(ctx sdk.Context, addr sdk.AccAddress)

type ProxyKeeper struct {
	bk        bank.Keeper
	rk        RestrictedKeeper
	listeners []BalanceListener
}

func Wrap(bk bank.Keeper, rk RestrictedKeeper) *ProxyKeeper {
	return &ProxyKeeper{bk, rk, nil}
}

func (pk *ProxyKeeper) AddBalanceListener(l BalanceListener) {
	pk.listeners = append(pk.listeners, l)
}

func (pk ProxyKeeper) notifyListeners(ctx sdk.Context, addr sdk.AccAddress) {
	for _, l := range pk.listeners {
		l(ctx, addr)
	}
}

func (pk ProxyKeeper) SetBalance(ctx sdk.Context, addr sdk.AccAddress, balance sdk.Coin) (err error) {
	if err = pk.bk.SetBalance(ctx, addr, balance); err == nil {
		pk.notifyListeners(ctx, addr)
	}
	return
}

func (pk ProxyKeeper) SetBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) (err error) {
	if err = pk.bk.SetBalances(ctx, addr, balances); err == nil {
		pk.notifyListeners(ctx, addr)
	}

	return
}

func (pk ProxyKeeper) InputOutputCoins(ctx sdk.Context, inputs []bank.Input, outputs []bank.Output) error {
	restrictedDenoms := pk.rk.GetRestrictedDenoms(ctx)
	// Multisend does not support restricted denominations.
	for _, input := range inputs {
		for _, coin := range input.Coins {
			if _, found := restrictedDenoms.Find(coin.Denom); found {
				return sdkerrors.Wrap(ErrRestrictedDenominationUsed, coin.Denom)
			}
		}
	}

	err := pk.bk.InputOutputCoins(ctx, inputs, outputs)
	if err == nil {
		for _, input := range inputs {
			pk.notifyListeners(ctx, input.Address)
		}

		for _, output := range outputs {
			pk.notifyListeners(ctx, output.Address)
		}
	}
	return err
}

func (pk ProxyKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	restrictedDenoms := pk.rk.GetRestrictedDenoms(ctx)
	for _, c := range amt {
		if denom, found := restrictedDenoms.Find(c.Denom); found {
			if !denom.IsAnyAllowed(fromAddr, toAddr) {
				return sdkerrors.Wrap(ErrRestrictedDenominationUsed, c.Denom)
			}
		}
	}

	err := pk.bk.SendCoins(ctx, fromAddr, toAddr, amt)
	if err == nil {
		pk.notifyListeners(ctx, fromAddr)
		pk.notifyListeners(ctx, toAddr)
	}
	return err
}

func (pk ProxyKeeper) SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error) {
	c, err := pk.bk.SubtractCoins(ctx, addr, amt)
	if err == nil {
		pk.notifyListeners(ctx, addr)
	}
	return c, err
}

func (pk ProxyKeeper) AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error) {
	c, err := pk.bk.AddCoins(ctx, addr, amt)
	if err == nil {
		pk.notifyListeners(ctx, addr)
	}
	return c, err
}

func (pk ProxyKeeper) DelegateCoins(ctx sdk.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error {
	err := pk.bk.DelegateCoins(ctx, delegatorAddr, moduleAccAddr, amt)
	if err == nil {
		pk.notifyListeners(ctx, delegatorAddr)
	}
	return err
}

func (pk ProxyKeeper) UndelegateCoins(ctx sdk.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error {
	err := pk.bk.UndelegateCoins(ctx, moduleAccAddr, delegatorAddr, amt)
	if err == nil {
		pk.notifyListeners(ctx, delegatorAddr)
	}
	return err
}

func (pk ProxyKeeper) ValidateBalance(ctx sdk.Context, addr sdk.AccAddress) error {
	return pk.bk.ValidateBalance(ctx, addr)
}

func (pk ProxyKeeper) HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool {
	return pk.bk.HasBalance(ctx, addr, amt)
}

func (pk ProxyKeeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return pk.bk.GetAllBalances(ctx, addr)
}

func (pk ProxyKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return pk.bk.GetBalance(ctx, addr, denom)
}

func (pk ProxyKeeper) LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return pk.bk.LockedCoins(ctx, addr)
}

func (pk ProxyKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return pk.bk.SpendableCoins(ctx, addr)
}

func (pk ProxyKeeper) IterateAccountBalances(ctx sdk.Context, addr sdk.AccAddress, cb func(coin sdk.Coin) (stop bool)) {
	pk.bk.IterateAccountBalances(ctx, addr, cb)
}

func (pk ProxyKeeper) IterateAllBalances(ctx sdk.Context, cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool)) {
	pk.bk.IterateAllBalances(ctx, cb)
}

func (pk ProxyKeeper) GetSendEnabled(ctx sdk.Context) bool {
	return pk.bk.GetSendEnabled(ctx)
}

func (pk ProxyKeeper) SetSendEnabled(ctx sdk.Context, enabled bool) {
	pk.bk.SetSendEnabled(ctx, enabled)
}

func (pk ProxyKeeper) BlacklistedAddr(addr sdk.AccAddress) bool {
	return pk.bk.BlacklistedAddr(addr)
}
