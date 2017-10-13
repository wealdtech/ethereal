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
	"github.com/ethereum/go-ethereum/core/types"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

// transactionInfoCmd represents the transaction info command
var transactionInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a transaction",
	Long: `Obtain information about a transaction.  For example:

    ethereal transaction info --transaction=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the transaction exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		txHash := common.HexToHash(transactionStr)
		tx, pending, err := client.TransactionByHash(context.Background(), txHash)

		cli.ErrCheck(err, quiet, "Failed to obtain transaction")

		if quiet {
			os.Exit(0)
		}

		var receipt *types.Receipt
		if pending {
			if tx.To() == nil {
				fmt.Printf("Type:\t\t\tPending contract creation\n")
			} else {
				fmt.Printf("Type:\t\t\tPending transaction\n")
			}
		} else {
			if tx.To() == nil {
				fmt.Printf("Type:\t\t\tMined contract creation\n")
			} else {
				fmt.Printf("Type:\t\t\tMined transaction\n")
			}
			receipt, err = client.TransactionReceipt(context.Background(), txHash)
		}

		// TODO: From

		// To
		if tx.To() == nil {
			if receipt != nil {
				contractAddress := receipt.ContractAddress
				to, err := ens.ReverseResolve(client, &contractAddress)
				if err == nil {
					fmt.Printf("Contract address:\t%v (%s)\n", to, contractAddress.Hex())
				} else {
					fmt.Printf("Contract address:\t%v\n", contractAddress.Hex())
				}
			}
		} else {
			to, err := ens.ReverseResolve(client, tx.To())
			if err == nil {
				fmt.Printf("To:\t\t\t%v (%s)\n", to, tx.To().Hex())
			} else {
				fmt.Printf("To:\t\t\t%v\n", tx.To().Hex())
			}
		}

		fmt.Printf("Nonce:\t\t\t%v\n", tx.Nonce())
		fmt.Printf("Gas limit:\t\t%v\n", tx.Gas())
		if receipt != nil {
			fmt.Printf("Gas used:\t\t%v\n", receipt.GasUsed)
		}
		fmt.Printf("Gas price:\t\t%v\n", etherutils.WeiToString(tx.GasPrice(), true))
		fmt.Printf("Value:\t\t\t%v\n", etherutils.WeiToString(tx.Value(), true))
	},
}

func init() {
	transactionFlags(transactionInfoCmd)
	transactionCmd.AddCommand(transactionInfoCmd)
}
