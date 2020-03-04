// This software is Copyright (c) 2019-2020 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package bank

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	appcodec "github.com/e-money/em-ledger/app/codec"
	"github.com/e-money/em-ledger/types"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	emauth "github.com/e-money/em-ledger/hooks/auth"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func TestProxySendCoins(t *testing.T) {
	ctx, ak, bk := createTestComponents(t)

	var (
		acc1 = createAccount(ctx, ak, bk, "acc1", "150000gbp, 150000usd, 150000sek")
		acc2 = createAccount(ctx, ak, bk, "acc2", "150000gbp, 150000usd, 150000sek")
		dest = sdk.AccAddress([]byte("dest"))
	)

	bk.rk = restrictedKeeper{
		RestrictedDenoms: []types.RestrictedDenom{
			{"gbp", []sdk.AccAddress{}},
			{"usd", []sdk.AccAddress{acc1.GetAddress()}},
		},
	}

	var testdata = []struct {
		denom string
		acc   sdk.AccAddress
		valid bool
	}{
		{"gbp", acc2.GetAddress(), false},
		{"usd", acc2.GetAddress(), false},
		{"gbp", acc1.GetAddress(), false},
		{"usd", acc1.GetAddress(), true},
		{"sek", acc1.GetAddress(), true},
		{"sek", acc2.GetAddress(), true},
	}

	for _, d := range testdata {
		c := fmt.Sprintf("1000%s", d.denom)
		err := bk.SendCoins(ctx, d.acc, dest, coins(c))
		if d.valid {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			require.True(t, strings.HasPrefix(err.Error(), ErrRestrictedDenominationUsed.Error()))
		}
	}
}

func TestInputOutputCoins(t *testing.T) {
	ctx, ak, bk := createTestComponents(t)

	var (
		acc1 = createAccount(ctx, ak, bk, "acc1", "150000gbp, 150000usd, 150000sek")
		acc2 = createAccount(ctx, ak, bk, "acc2", "150000gbp, 150000usd, 150000sek")
		acc3 = createAccount(ctx, ak, bk, "acc3", "")
	)

	// For simplicity's sake, inputoutput will reject any transaction that includes restricted denominations.

	bk.rk = restrictedKeeper{
		RestrictedDenoms: []types.RestrictedDenom{
			{"gbp", []sdk.AccAddress{}},
			{"usd", []sdk.AccAddress{acc1.GetAddress()}},
		},
	}

	var testdata = []struct {
		inputs  []bank.Input
		outputs []bank.Output
		valid   bool
	}{
		{[]bank.Input{}, []bank.Output{}, true},
		{
			inputs: []bank.Input{
				{acc1.GetAddress(), coins("1000sek")},
			},
			outputs: []bank.Output{
				{acc2.GetAddress(), coins("500sek")},
				{acc3.GetAddress(), coins("500sek")},
			},
			valid: true,
		},
		{
			inputs: []bank.Input{
				{acc1.GetAddress(), coins("500sek, 1000gbp")},
			},
			outputs: []bank.Output{
				{acc2.GetAddress(), coins("500sek, 500gbp")},
				{acc3.GetAddress(), coins("500gbp")},
			},
			valid: false,
		},
		{
			inputs: []bank.Input{
				{acc1.GetAddress(), coins("1000usd")},
			},
			outputs: []bank.Output{
				{acc2.GetAddress(), coins("1000usd")},
			},
			valid: false,
		},
	}

	for _, d := range testdata {
		err := bk.InputOutputCoins(ctx, d.inputs, d.outputs)
		if d.valid {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			require.True(t, strings.HasPrefix(err.Error(), ErrRestrictedDenominationUsed.Error()))
		}
	}
}

func createTestComponents(t *testing.T) (sdk.Context, auth.AccountKeeper, ProxyKeeper) {
	var (
		authCapKey = sdk.NewKVStoreKey("authCapKey")
		keyBank    = sdk.NewKVStoreKey(bank.StoreKey)
		keyParams  = sdk.NewKVStoreKey("params")
		tkeyParams = sdk.NewTransientStoreKey("transient_params")

		blacklistedAddrs = make(map[string]bool)
	)

	cdc := createCodec()
	appCodec := appcodec.NewAppCodec(cdc)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)

	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	pk := params.NewKeeper(appCodec, keyParams, tkeyParams)

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain"}, true, log.NewNopLogger())
	accountKeeper := auth.NewAccountKeeper(appCodec, authCapKey, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	accountKeeperWrapped := emauth.Wrap(accountKeeper)

	bankKeeper := bank.NewBaseKeeper(appCodec, keyBank, accountKeeperWrapped, pk.Subspace(bank.DefaultParamspace), blacklistedAddrs)

	wrappedBK := Wrap(bankKeeper, restrictedKeeper{})

	return ctx, accountKeeper, wrappedBK
}

type restrictedKeeper struct {
	RestrictedDenoms types.RestrictedDenoms
}

func (rk restrictedKeeper) GetRestrictedDenoms(sdk.Context) types.RestrictedDenoms {
	return rk.RestrictedDenoms
}

func createAccount(ctx sdk.Context, ak auth.AccountKeeper, bk bank.Keeper, address, balance string) exported.Account {
	acc := ak.NewAccountWithAddress(ctx, sdk.AccAddress([]byte(address)))

	//acc.SetCoins(coins(balance))
	ak.SetAccount(ctx, acc)
	bk.SetBalances(ctx, acc.GetAddress(), coins(balance))

	return acc
}

func coins(s string) sdk.Coins {
	coins, err := sdk.ParseCoins(s)
	if err != nil {
		panic(err)
	}
	return coins
}

func createCodec() *codec.Codec {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)

	return cdc
}
