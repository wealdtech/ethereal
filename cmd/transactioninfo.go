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
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
	"github.com/wealdtech/ethereal/util"
)

// transactionInfoCmd represents the transaction info command
var transactionInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a transaction",
	Long: `Obtain information about a transaction.  For example:

    ethereal transaction info --transaction=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the transaction exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(transactionStr != "", quiet, "--transaction is required")
		txHash := common.HexToHash(transactionStr)
		ctx, cancel := localContext()
		defer cancel()
		tx, pending, err := client.TransactionByHash(ctx, txHash)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain transaction %s", txHash.Hex()))

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
			ctx, cancel := localContext()
			defer cancel()
			receipt, err = client.TransactionReceipt(ctx, txHash)
			// TODO need the block number to calculate the status?
			// Any other way to recognise non-Byzantium receipts?
			if receipt.Status == 0 {
				fmt.Printf("Result:\t\t\tFailed\n")
			} else {
				fmt.Printf("Result:\t\t\tSucceeded\n")
			}
		}

		fromAddress, err := txFrom(tx)
		if err == nil {
			to, err := ens.ReverseResolve(client, &fromAddress)
			if err == nil {
				fmt.Printf("From:\t\t\t%v (%s)\n", to, fromAddress.Hex())
			} else {
				fmt.Printf("From:\t\t\t%v\n", fromAddress.Hex())
			}
		}

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
		fmt.Printf("Gas price:\t\t%v\n", etherutils.WeiToString(tx.GasPrice(), true))
		fmt.Printf("Gas limit:\t\t%v\n", tx.Gas())
		if receipt != nil {
			fmt.Printf("Gas used:\t\t%v\n", receipt.GasUsed)
			fmt.Printf("Gas cost:\t\t%v\n", etherutils.WeiToString(big.NewInt(0).Mul(tx.GasPrice(), receipt.GasUsed), true))
		}
		fmt.Printf("Value:\t\t\t%v\n", etherutils.WeiToString(tx.Value(), true))

		if receipt != nil && len(receipt.Logs) != 0 {
			if !verbose {
				fmt.Printf("Log entries:\t\t%v\n", len(receipt.Logs))
			} else {
				fmt.Println("Logs")
				util.InitLogDefinitions()
				for i, log := range receipt.Logs {
					logDefinition := util.LogDefinitions[log.Topics[0]]
					if logDefinition == nil {
						fmt.Printf(" %3d:\t\t\tUnknown(", i)
						for j := 1; j < len(log.Topics); j++ {
							if j > 1 {
								fmt.Printf(",")
							}
							fmt.Printf("%s", log.Topics[j].Hex())
						}
						fmt.Printf(")\n")
					} else {
						fmt.Println(logDefinition.Signature)
					}
				}
			}
		}
	},
}

func init() {
	transactionCmd.AddCommand(transactionInfoCmd)
	transactionFlags(transactionInfoCmd)
}
