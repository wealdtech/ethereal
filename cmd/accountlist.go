// Copyright © 2017 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"os"

	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

// accountListCmd represents the account list command
var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "List visible accounts",
	Long: `List accounts that are visible to Ethereal.  For example:

    ethereal account list

In quiet mode this will return 0 if any accounts are found, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := cli.ObtainWallets(chainID)
		foundAccounts := false
		if err == nil {
			for _, wallet := range wallets {
				for _, account := range wallet.Accounts() {
					foundAccounts = true
					if !quiet {
						fmt.Printf("%s", account.Address.Hex())
						if verbose {
							name, err := ens.ReverseResolve(client, &account.Address)
							if err == nil {
								fmt.Printf(" (%s)", name)
							}
							balance, err := client.BalanceAt(context.Background(), account.Address, nil)
							if err == nil {
								fmt.Printf(" %s", etherutils.WeiToString(balance, true))
							}
						}
						fmt.Printf("\n")
					}
				}
			}
		}

		if quiet {
			if foundAccounts {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	},
}

func init() {
	accountCmd.AddCommand(accountListCmd)
}
