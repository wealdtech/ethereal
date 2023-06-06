// Copyright © 2017-2022 Weald Technology Trading
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
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util/txdata"
	ens "github.com/wealdtech/go-ens/v3"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	transactionInfoRaw        bool
	transactionInfoJSON       bool
	transactionInfoSignatures string
)

// transactionInfoCmd represents the transaction info command.
var transactionInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a transaction",
	Long: `Obtain information about a transaction.  For example:

    ethereal transaction info --transaction=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the transaction exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(transactionStr != "", quiet, "--transaction is required")
		var txHash common.Hash
		var pending bool
		var tx *types.Transaction
		if !strings.HasPrefix(transactionStr, "0x") {
			// Read from file.
			fileBytes, err := os.ReadFile(transactionStr)
			cli.ErrCheck(err, quiet, "Failed to read transaction from filesystem")
			transactionStr = strings.TrimSpace(string(fileBytes))
		}
		if len(transactionStr) > 66 {
			// Assume input is a raw transaction.
			data, err := hex.DecodeString(strings.TrimPrefix(transactionStr, "0x"))
			cli.ErrCheck(err, quiet, "Failed to decode data")
			tx = &types.Transaction{}
			stream := rlp.NewStream(bytes.NewReader(data), 0)
			err = tx.DecodeRLP(stream)
			cli.ErrCheck(err, quiet, "Failed to decode raw transaction")
			txHash = tx.Hash()
			// Assume pending.
			pending = true
		} else {
			// Assume input is a transaction ID.
			txHash = common.HexToHash(transactionStr)
			ctx, cancel := localContext()
			defer cancel()
			var err error
			tx, pending, err = c.Client().TransactionByHash(ctx, txHash)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain transaction %s", txHash.Hex()))
		}

		if quiet {
			os.Exit(exitSuccess)
		}

		if transactionInfoRaw {
			buf := new(bytes.Buffer)
			cli.ErrCheck(tx.EncodeRLP(buf), quiet, "failed to encode transaction")
			fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			os.Exit(exitSuccess)
		}

		if transactionInfoJSON {
			json, err := tx.MarshalJSON()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain JSON for transaction %s", txHash.Hex()))
			fmt.Printf("%s\n", string(json))
			os.Exit(exitSuccess)
		}

		txdata.InitFunctionMap()
		if transactionInfoSignatures != "" {
			for _, signature := range strings.Split(transactionInfoSignatures, ";") {
				txdata.AddFunctionSignature(signature)
			}
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
			receipt, err = c.Client().TransactionReceipt(ctx, txHash)
			if receipt != nil {
				if receipt.Status == 0 {
					fmt.Printf("Result:\t\t\tFailed\n")
				} else {
					fmt.Printf("Result:\t\t\tSucceeded\n")
				}
			}
		}

		if receipt != nil && len(receipt.Logs) > 0 {
			// We can obtain the block number from the log.
			fmt.Printf("Block:\t\t\t%d\n", receipt.Logs[0].BlockNumber)
		}

		fromAddress, err := types.Sender(signer, tx)
		if err == nil {
			fmt.Printf("From:\t\t\t%v\n", ens.Format(c.Client(), fromAddress))
		}

		// To
		if tx.To() == nil {
			if receipt != nil {
				fmt.Printf("Contract address:\t%v\n", ens.Format(c.Client(), receipt.ContractAddress))
			}
		} else {
			fmt.Printf("To:\t\t\t%v\n", ens.Format(c.Client(), *tx.To()))
		}

		if verbose {
			switch tx.Type() {
			case types.LegacyTxType:
				fmt.Println("Transaction type:\tLegacy")
			case types.DynamicFeeTxType:
				fmt.Println("Transaction type:\tDynamic")
			case types.AccessListTxType:
				fmt.Println("Transaction type:\tAccess list")
			default:
				fmt.Println("Transaction type:\tUnknown")
			}
			fmt.Printf("Nonce:\t\t\t%v\n", tx.Nonce())
			fmt.Printf("Gas limit:\t\t%v\n", tx.Gas())
		}
		if receipt != nil {
			fmt.Printf("Gas used:\t\t%v\n", receipt.GasUsed)
		}
		switch tx.Type() {
		case types.LegacyTxType, types.AccessListTxType:
			fmt.Printf("Gas price:\t\t%v\n", string2eth.WeiToString(tx.GasPrice(), true))
		case types.DynamicFeeTxType:
			fmt.Printf("Max fee per gas:\t%v\n", string2eth.WeiToString(tx.GasFeeCap(), true))
		}

		var block *types.Block
		if receipt != nil {
			block, err = c.Client().BlockByHash(context.Background(), receipt.BlockHash)
			if err != nil {
				// We can carry on without it.
				block = nil
			}
		}

		if tx.Type() == types.DynamicFeeTxType {
			if receipt != nil && block != nil {
				fmt.Printf("Actual fee per gas:\t%v\n", string2eth.WeiToString(block.BaseFee(), true))
			}
			fmt.Printf("Tip per gas:\t\t%v\n", string2eth.WeiToString(tx.GasTipCap(), true))
		}

		if receipt != nil {
			gasUsed := big.NewInt(int64(receipt.GasUsed))
			switch tx.Type() {
			case types.LegacyTxType, types.AccessListTxType:
				fmt.Printf("Total fee:\t\t%v", string2eth.WeiToString(new(big.Int).Mul(tx.GasPrice(), gasUsed), true))
				if verbose {
					fmt.Printf(" (%v * %v)\n", string2eth.WeiToString(tx.GasPrice(), true), gasUsed)
				} else {
					fmt.Println()
				}
			case types.DynamicFeeTxType:
				if block != nil {
					fmt.Printf("Total fee:\t\t%v", string2eth.WeiToString(new(big.Int).Mul(new(big.Int).Add(block.BaseFee(), tx.GasTipCap()), gasUsed), true))
					if verbose {
						fmt.Printf(" ((%v + %v) * %v)\n", string2eth.WeiToString(block.BaseFee(), true), string2eth.WeiToString(tx.GasTipCap(), true), gasUsed)
					} else {
						fmt.Println()
					}
				}
			}
		}
		fmt.Printf("Value:\t\t\t%v\n", string2eth.WeiToString(tx.Value(), true))

		if tx.To() != nil && len(tx.Data()) > 0 {
			fmt.Printf("Data:\t\t\t%v\n", txdata.DataToString(c.Client(), tx.Data()))
		}

		if verbose && receipt != nil && len(receipt.Logs) > 0 {
			fmt.Printf("Logs:\n")
			for i, log := range receipt.Logs {
				fmt.Printf("\t%d:\n", i)
				fmt.Printf("\t\tFrom:\t%v\n", ens.Format(c.Client(), log.Address))
				// Try to obtain decoded log.
				decoded := txdata.EventToString(c.Client(), log)
				if decoded != "" {
					fmt.Printf("\t\tEvent:\t%s\n", decoded)
				} else {
					if len(log.Topics) > 0 {
						fmt.Printf("\t\tTopics:\n")
						for j, topic := range log.Topics {
							fmt.Printf("\t\t\t%d:\t%v\n", j, topic.Hex())
						}
					}
					if len(log.Data) > 0 {
						fmt.Printf("\t\tData:\n")
						for j := 0; j*32 < len(log.Data); j++ {
							fmt.Printf("\t\t\t%d:\t0x%s\n", j, hex.EncodeToString(log.Data[j*32:(j+1)*32]))
						}
					}
				}
			}
		}
	},
}

func init() {
	transactionCmd.AddCommand(transactionInfoCmd)
	transactionFlags(transactionInfoCmd)
	transactionInfoCmd.Flags().BoolVar(&transactionInfoRaw, "raw", false, "Output the transaction as raw hex")
	transactionInfoCmd.Flags().BoolVar(&transactionInfoJSON, "json", false, "Output the transaction as json")
	transactionInfoCmd.Flags().StringVar(&transactionInfoSignatures, "signatures", "", "Semicolon-separated list of custom transaction signatures (e.g. myFunc(address,bytes32);myFunc2(bool)")
}
