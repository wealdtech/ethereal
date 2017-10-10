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

	"github.com/ethereum/go-ethereum/common"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

// txInfoCmd represents the tx info command
var txInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a transaction",
	Long: `Obtain information about a transaction.  For example:

    ethereal tx info --transaction=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the transaction exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(txStr)
		transaction, pending, err := client.TransactionByHash(context.Background(), txHash)

		cli.ErrCheck(err, quiet, "Failed to obtain transaction")

		if quiet {
			os.Exit(0)
		}

		if pending {
			fmt.Printf("Type:\t\tPending transaction\n")
		} else {
			fmt.Printf("Type:\t\tMined transaction\n")
		}

		// TODO: From

		// To
		if transaction.To() == nil {
			// Contract creation; calculate the address of the contract
			fmt.Printf("Type:\t\tContract creation\n")
			// TODO need From address to work this out
			//			contractAddress := common.StringToAddress("0x01")
			//			to, err := ens.ReverseResolve(client, &contractAddress)
			//			if err == nil {
			//				fmt.Printf("Contract creation:\t%v (%s)\n", to, contractAddress)
			//			} else {
			//				fmt.Printf("Contract creation:\t%v\n", contractAddress)
			//			}
		} else {
			to, err := ens.ReverseResolve(client, transaction.To())
			if err == nil {
				fmt.Printf("To:\t\t%v (%s)\n", to, transaction.To().Hex())
			} else {
				fmt.Printf("To:\t\t%v\n", transaction.To().Hex())
			}
		}

		fmt.Printf("Nonce:\t\t%v\n", transaction.Nonce())
		fmt.Printf("Gas limit:\t%v\n", transaction.Gas())
		fmt.Printf("Gas price:\t%v\n", etherutils.WeiToString(transaction.GasPrice(), true))
		fmt.Printf("Value:\t\t%v\n", etherutils.WeiToString(transaction.Value(), true))
	},
}

func init() {
	txFlags(txInfoCmd)
	txCmd.AddCommand(txInfoCmd)
}
