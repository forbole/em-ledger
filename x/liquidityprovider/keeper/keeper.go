// This software is Copyright (c) 2019 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/e-money/em-ledger/x/liquidityprovider/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	authKeeper   types.AccountKeeper
	bankKeeper   types.BankKeeper
	supplyKeeper supply.Keeper
}

func NewKeeper(ak types.AccountKeeper, bk types.BankKeeper, sk supply.Keeper) Keeper {
	return Keeper{
		authKeeper:   ak,
		bankKeeper:   bk,
		supplyKeeper: sk,
	}
}

func (k Keeper) CreateLiquidityProvider(ctx sdk.Context, address sdk.AccAddress, mintable sdk.Coins) (*sdk.Result, error) {
	logger := k.Logger(ctx)

	account := k.authKeeper.GetAccount(ctx, address)
	if account == nil {
		logger.Info("Account not found", "account", address)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, address.String())
		//return types.ErrAccountDoesNotExist(address).Result()
	}

	baseAccount, ok := account.(*auth.BaseAccount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, address.String())
	}

	lpAcc := types.NewLiquidityProviderAccount(*baseAccount, mintable)
	k.authKeeper.SetAccount(ctx, lpAcc)

	logger.Info("Created liquidity provider account.", "account", lpAcc.GetAddress())
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func (k Keeper) BurnTokensFromBalance(ctx sdk.Context, liquidityProvider sdk.AccAddress, amount sdk.Coins) (*sdk.Result, error) {
	account := k.GetLiquidityProviderAccount(ctx, liquidityProvider)
	if account == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s is not a liquidity provider or does not exist", liquidityProvider.String())
		//return sdk.ErrUnknownAddress(fmt.Sprintf("account %s is not a liquidity provider or does not exist", liquidityProvider.String())).Result()
	}

	spendable := k.bankKeeper.SpendableCoins(ctx, account.Address)
	_, anynegative := spendable.SafeSub(amount)
	//_, anynegative := account.BaseAccount.SpendableCoins(ctx.BlockTime()).SafeSub(amount)
	if anynegative {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Insufficient balance for burn operation: %s < %s", spendable, amount)
		//return sdk.ErrInsufficientCoins(fmt.Sprintf("Insufficient balance for burn operation: %s < %s", account.Account.GetCoins(), amount)).Result()
	}

	err := k.supplyKeeper.SendCoinsFromAccountToModule(ctx, liquidityProvider, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	err = k.supplyKeeper.BurnCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	account = k.GetLiquidityProviderAccount(ctx, liquidityProvider)
	account.Mintable = account.Mintable.Add(amount...)
	k.SetLiquidityProviderAccount(ctx, account)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func (k Keeper) MintTokens(ctx sdk.Context, liquidityProvider sdk.AccAddress, amount sdk.Coins) (*sdk.Result, error) {
	logger := k.Logger(ctx)

	account := k.GetLiquidityProviderAccount(ctx, liquidityProvider)
	if account == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s is not a liquidity provider or does not exist", liquidityProvider.String())
		//return sdk.ErrUnknownAddress(fmt.Sprintf("account %s is not a liquidity provider or does not exist", liquidityProvider.String())).Result()
	}

	updatedMintableAmount, anyNegative := account.Mintable.SafeSub(amount)
	if anyNegative {
		logger.Debug(fmt.Sprintf("Insufficient mintable amount for minting operation"), "requested", amount, "available", account.Mintable)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient liquidity provider mintable amount: %s < %s", account.Mintable, amount)
		//return sdk.ErrInsufficientCoins(fmt.Sprintf("insufficient liquidity provider mintable amount: %s < %s", account.Mintable, amount)).Result()
	}

	err := k.supplyKeeper.MintCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return nil, err
		//return err.Result()
	}

	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidityProvider, amount)
	if err != nil {
		return nil, err
		//return err.Result()
	}

	account = k.GetLiquidityProviderAccount(ctx, liquidityProvider)
	account.Mintable = updatedMintableAmount
	k.SetLiquidityProviderAccount(ctx, account)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func (k Keeper) SetLiquidityProviderAccount(ctx sdk.Context, account *types.LiquidityProviderAccount) {
	k.authKeeper.SetAccount(ctx, account)
}

func (k Keeper) RevokeLiquidityProviderAccount(ctx sdk.Context, account exported.Account) bool {
	if lpAcc, isLpAcc := account.(*types.LiquidityProviderAccount); isLpAcc {
		account = &lpAcc.BaseAccount
		k.authKeeper.SetAccount(ctx, account)
		return true
	}

	return false
}

func (k Keeper) GetLiquidityProviderAccount(ctx sdk.Context, liquidityProvider sdk.AccAddress) *types.LiquidityProviderAccount {
	logger := k.Logger(ctx)

	a := k.authKeeper.GetAccount(ctx, liquidityProvider)
	account, ok := a.(*types.LiquidityProviderAccount)
	if !ok {
		logger.Debug(fmt.Sprintf("Account is not a liquidity provider"), "address", liquidityProvider)
		return nil
	}

	return account
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
