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
	"math/big"
	"os"

	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

var etherBalanceAddress string
var etherBalanceWei bool

// etherBalanceCmd represents the ether balance command
var etherBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Obtain the balance for an address",
	Long: `Obtain the Ether balance for an address.  For example:

    ethereal ether balance --address=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the balance is greater than 0, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(etherBalanceAddress != "", quiet, "--address is required")
		address, err := ens.Resolve(client, etherBalanceAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain address")

		balance, err := client.BalanceAt(context.Background(), address, nil)
		cli.ErrCheck(err, quiet, "Failed to obtain balance")

		if quiet {
			if balance.Cmp(big.NewInt(0)) == 0 {
				os.Exit(1)
			} else {
				os.Exit(0)
			}
		}

		if etherBalanceWei {
			fmt.Printf("%s\n", balance.String())
		} else {
			fmt.Printf("%s\n", etherutils.WeiToString(balance, true))
		}
	},
}

func init() {
	etherCmd.AddCommand(etherBalanceCmd)
	etherBalanceCmd.Flags().BoolVar(&etherBalanceWei, "wei", false, "Dispay output in number of Wei")
	etherBalanceCmd.Flags().StringVar(&etherBalanceAddress, "address", "", "Address to show Ether balance")
}
