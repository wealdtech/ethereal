// Copyright © 2017-2025 Weald Technology Trading
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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	Run: func(_ *cobra.Command, _ []string) {
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
			cli.ErrCheck(tx.UnmarshalBinary(data), quiet, "Failed to decode raw transaction")
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
			fmt.Fprintf(os.Stdout, "0x%s\n", hex.EncodeToString(buf.Bytes()))
			os.Exit(exitSuccess)
		}

		if transactionInfoJSON {
			json, err := tx.MarshalJSON()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain JSON for transaction %s", txHash.Hex()))
			fmt.Fprintf(os.Stdout, "%s\n", string(json))
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
				fmt.Fprintf(os.Stdout, "Type: Pending contract creation\n")
			} else {
				fmt.Fprintf(os.Stdout, "Type: Pending transaction\n")
			}
		} else {
			if tx.To() == nil {
				fmt.Fprintf(os.Stdout, "Type: Mined contract creation\n")
			} else {
				fmt.Fprintf(os.Stdout, "Type: Mined transaction\n")
			}
			ctx, cancel := localContext()
			defer cancel()
			receipt, err = c.Client().TransactionReceipt(ctx, txHash)
			if receipt != nil {
				fmt.Fprintf(os.Stdout, "Block: %d\n", receipt.BlockNumber)
				if receipt.Status == 0 {
					fmt.Fprintf(os.Stdout, "Result: Failed\n")
					revertReason := obtainRevertReason(tx, receipt)
					if revertReason != "" {
						fmt.Fprintf(os.Stdout, "Reason: %s\n", revertReason)
					}
				} else {
					fmt.Fprintf(os.Stdout, "Result: Succeeded\n")
				}
			}
		}

		if receipt != nil && len(receipt.Logs) > 0 {
			// We can obtain the block number from the log.
			fmt.Fprintf(os.Stdout, "Block: %d\n", receipt.Logs[0].BlockNumber)
		}

		fromAddress, err := types.Sender(signer, tx)
		if err == nil {
			fmt.Fprintf(os.Stdout, "From: %v\n", ens.Format(c.Client(), fromAddress))
		}

		// To
		if tx.To() == nil {
			if receipt != nil {
				fmt.Fprintf(os.Stdout, "Contract address: %v\n", ens.Format(c.Client(), receipt.ContractAddress))
			}
		} else {
			fmt.Fprintf(os.Stdout, "To: %v\n", ens.Format(c.Client(), *tx.To()))
		}

		if verbose {
			switch tx.Type() {
			case types.LegacyTxType:
				fmt.Println("Transaction type: Legacy (type 0)")
			case types.AccessListTxType:
				fmt.Println("Transaction type: Access list (type 1)")
			case types.DynamicFeeTxType:
				fmt.Println("Transaction type: Dynamic (type 2)")
			case types.BlobTxType:
				fmt.Println("Transaction type: Blob (type 3)")
			case types.SetCodeTxType:
				fmt.Println("Transaction type: Set code (type 4)")
			default:
				fmt.Println("Transaction type: Unknown")
			}
			fmt.Fprintf(os.Stdout, "Nonce: %v\n", tx.Nonce())
			fmt.Fprintf(os.Stdout, "Gas limit: %v\n", tx.Gas())
		}
		if receipt != nil {
			fmt.Fprintf(os.Stdout, "Gas used: %v\n", receipt.GasUsed)
		}
		switch tx.Type() {
		case types.LegacyTxType, types.AccessListTxType:
			fmt.Fprintf(os.Stdout, "Gas price: %v\n", string2eth.WeiToString(tx.GasPrice(), true))
		case types.DynamicFeeTxType:
			fmt.Fprintf(os.Stdout, "Max fee per gas: %v\n", string2eth.WeiToString(tx.GasFeeCap(), true))
		case types.BlobTxType:
			fmt.Fprintf(os.Stdout, "Max fee per gas: %v\n", string2eth.WeiToString(tx.GasFeeCap(), true))
			fmt.Fprintf(os.Stdout, "Max fee per blob gas: %v\n", string2eth.WeiToString(tx.BlobGasFeeCap(), true))
			blobSidecars := tx.BlobTxSidecar()
			if blobSidecars != nil {
				for i := range blobSidecars.Blobs {
					fmt.Fprintf(os.Stdout, "Blob %d:\n", i)
					fmt.Fprintf(os.Stdout, "  Blob data: %#x\n", blobSidecars.Blobs[i])
					fmt.Fprintf(os.Stdout, "  Blob commitment: %#x\n", blobSidecars.Commitments[i])
					fmt.Fprintf(os.Stdout, "  Blob proofs: %#x\n", blobSidecars.Proofs[i])
				}
			}
		case types.SetCodeTxType:
			fmt.Fprintf(os.Stdout, "Max fee per gas: %v\n", string2eth.WeiToString(tx.GasFeeCap(), true))
			authorities := tx.SetCodeAuthorities()
			authorizations := tx.SetCodeAuthorizations()
			for i := range authorities {
				fmt.Fprintf(os.Stdout, "Authorization %d:\n", i)
				fmt.Fprintf(os.Stdout, "  Authority: %s\n", authorities[i].Hex())
				fmt.Fprintf(os.Stdout, "  Chain ID: %s\n", authorizations[i].ChainID.Hex())
				fmt.Fprintf(os.Stdout, "  Authorization: %s\n", authorizations[i].Address.Hex())
			}
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
				fmt.Fprintf(os.Stdout, "Actual fee per gas: %v\n", string2eth.WeiToString(block.BaseFee(), true))
			}
			fmt.Fprintf(os.Stdout, "Tip per gas: %v\n", string2eth.WeiToString(tx.GasTipCap(), true))
		}

		if receipt != nil {
			gasUsed := big.NewInt(int64(receipt.GasUsed))
			switch tx.Type() {
			case types.LegacyTxType, types.AccessListTxType:
				fmt.Fprintf(os.Stdout, "Total fee: %v", string2eth.WeiToString(new(big.Int).Mul(tx.GasPrice(), gasUsed), true))
				if verbose {
					fmt.Fprintf(os.Stdout, " (%v * %v)\n", string2eth.WeiToString(tx.GasPrice(), true), gasUsed)
				} else {
					fmt.Println()
				}
			case types.DynamicFeeTxType:
				if block != nil {
					fmt.Fprintf(os.Stdout, "Total fee: %v", string2eth.WeiToString(new(big.Int).Mul(new(big.Int).Add(block.BaseFee(), tx.GasTipCap()), gasUsed), true))
					if verbose {
						fmt.Fprintf(os.Stdout, " ((%v + %v) * %v)\n", string2eth.WeiToString(block.BaseFee(), true), string2eth.WeiToString(tx.GasTipCap(), true), gasUsed)
					} else {
						fmt.Println()
					}
				}
			}
		}
		fmt.Fprintf(os.Stdout, "Value: %v\n", string2eth.WeiToString(tx.Value(), true))

		if verbose {
			if tx.Type() == types.AccessListTxType || tx.Type() == types.DynamicFeeTxType || tx.Type() == types.SetCodeTxType {
				accessList := tx.AccessList()
				if len(accessList) > 0 {
					fmt.Fprintf(os.Stdout, "Access list:\n")
				}
				for _, tuple := range accessList {
					fmt.Fprintf(os.Stdout, "  Address: %s\n", tuple.Address.String())
					for _, key := range tuple.StorageKeys {
						fmt.Fprintf(os.Stdout, "    Storage key: %s\n", key.Hex())
					}
				}
			}
		}

		if tx.To() != nil && len(tx.Data()) > 0 {
			fmt.Fprintf(os.Stdout, "Data: %v\n", txdata.DataToString(c.Client(), tx.Data()))
		}

		if verbose && receipt != nil && len(receipt.Logs) > 0 {
			fmt.Fprintf(os.Stdout, "Logs:\n")
			for i, log := range receipt.Logs {
				fmt.Fprintf(os.Stdout, "  %2d:\n", i)
				fmt.Fprintf(os.Stdout, "    From: %v\n", ens.Format(c.Client(), log.Address))
				// Try to obtain decoded log.
				decoded := txdata.EventToString(c.Client(), log)
				if decoded != "" {
					fmt.Fprintf(os.Stdout, "    Event: %s\n", decoded)
				} else {
					if len(log.Topics) > 0 {
						fmt.Fprintf(os.Stdout, "    Topics:\n")
						for j, topic := range log.Topics {
							fmt.Fprintf(os.Stdout, "        %d: %v\n", j, topic.Hex())
						}
					}
					if len(log.Data) > 0 {
						fmt.Fprintf(os.Stdout, "    Data:\n")
						for j := 0; j*32 < len(log.Data); j++ {
							end := min((j+1)*32, len(log.Data))
							fmt.Fprintf(os.Stdout, "        %d: 0x%s\n", j, hex.EncodeToString(log.Data[j*32:end]))
						}
					}
				}
			}
		}
	},
}

func obtainRevertReason(tx *types.Transaction, receipt *types.Receipt) string {
	// To obtain the revert reason we rerun the transaction as a call.
	from, err := types.Sender(types.NewCancunSigner(tx.ChainId()), tx)
	if err != nil {
		// Failed to obtain sender; not a transaction error.
		return ""
	}

	msg := ethereum.CallMsg{
		From:  from,
		To:    tx.To(),
		Gas:   tx.Gas(),
		Value: tx.Value(),
		Data:  tx.Data(),
	}
	switch tx.Type() {
	case types.LegacyTxType:
		msg.GasPrice = tx.GasPrice()
	case types.AccessListTxType:
		msg.AccessList = tx.AccessList()
	case types.DynamicFeeTxType:
		msg.GasFeeCap = tx.GasFeeCap()
		msg.GasTipCap = tx.GasTipCap()
		msg.AccessList = tx.AccessList()
	case types.BlobTxType:
		msg.GasFeeCap = tx.GasFeeCap()
		msg.GasTipCap = tx.GasTipCap()
		msg.AccessList = tx.AccessList()
	}

	_, err = c.Client().CallContract(context.Background(), msg, receipt.BlockNumber)
	if err == nil {
		// Did not fail.
		return ""
	}

	revertReason := err.Error()
	if !strings.Contains(revertReason, "reverted") {
		// Did not revert.
		return ""
	}

	return revertReason
}

func init() {
	transactionCmd.AddCommand(transactionInfoCmd)
	transactionFlags(transactionInfoCmd)
	transactionInfoCmd.Flags().BoolVar(&transactionInfoRaw, "raw", false, "Output the transaction as raw hex")
	transactionInfoCmd.Flags().BoolVar(&transactionInfoJSON, "json", false, "Output the transaction as json")
	transactionInfoCmd.Flags().StringVar(&transactionInfoSignatures, "signatures", "", "Semicolon-separated list of custom transaction signatures (e.g. myFunc(address,bytes32);myFunc2(bool)")
}
