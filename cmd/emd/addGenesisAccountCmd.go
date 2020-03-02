// This software is Copyright (c) 2019 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/genutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	appcodec "github.com/e-money/em-ledger/app/codec"

	"github.com/tendermint/tendermint/libs/cli"
)

const (
	flagClientHome = "home-client"
)

// Adapted from Cosmos-sdk : x/genaccounts/client/cli/genesis_accts.go
func addGenesisAccountCmd(ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account [address_or_key_name] [coin][,[coin]]",
		Short: "Add genesis account to genesis.json",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			var pubkey crypto.PubKey
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				kb, err := keys.NewKeyBaseFromDir(viper.GetString(flagClientHome))
				if err != nil {
					return err
				}

				info, err := kb.Get(args[0])
				if err != nil {
					return err
				}

				addr = info.GetAddress()
				pubkey = info.GetPubKey()
			}

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			genAcc := auth.NewBaseAccount(addr, pubkey, 0, 0)
			//genAcc := genaccounts.NewGenesisAccountRaw(addr, coins, sdk.NewCoins(), 0, 0, "")
			if err := genAcc.Validate(); err != nil {
				return err
			}

			// retrieve the app state
			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
			if err != nil {
				return err
			}

			// add genesis account to the app state
			appCodec := appcodec.NewAppCodec(cdc)
			authGenState := auth.GetGenesisStateFromAppState(appCodec, appState)
			bankGenState := bank.GetGenesisStateFromAppState(appCodec.Amino, appState)

			if authGenState.Accounts.Contains(addr) {
				return fmt.Errorf("cannot add account at existing address %s", addr)
			}

			balance := bank.Balance{Address: addr, Coins: coins}

			bankGenState.Balances = append(bankGenState.Balances, balance)
			bankGenState.Balances = bank.SanitizeGenesisBalances(bankGenState.Balances)

			// Add the new account to the set of genesis accounts and sanitize the
			// accounts afterwards.
			authGenState.Accounts = append(authGenState.Accounts, genAcc)
			authGenState.Accounts = auth.SanitizeGenesisAccounts(authGenState.Accounts)

			authGenStateBz, err := cdc.MarshalJSON(authGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal auth genesis state: %w", err)
			}

			bankGenStateBz, err := cdc.MarshalJSON(bankGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal bank genesis state: %w", err)
			}

			appState[auth.ModuleName] = authGenStateBz
			appState[bank.ModuleName] = bankGenStateBz

			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			// export app state
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")
	return cmd
}
