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
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

var contractSendAmount string
var contractSendFromAddress string
var contractSendCall string
var contractSendReturns string

// contractSendCmd represents the contract call command
var contractSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a contract method",
	Long: `Send a contract method to the blockchain.  For example:

   ethereal contract send --contract=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07 --abi="./erc20.abi" --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --call="transfer(0x5FfC014343cd971B7eb70732021E26C35B744cc4, 10)" --passphrase=secret

In quiet mode this will return 0 if the transaction is successfully sent, otherwise 1.`,
	Aliases: []string{"transaction", "transmit"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(contractSendFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractSendFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractSendFromAddress))

		// We need to have 'call' and 'abi'
		cli.Assert(contractSendCall != "", quiet, "--call is required")

		var abi abi.ABI
		if contractAbi == "" {
			// TODO See if we can fetch the ABI from ENS
			cli.Err(quiet, "--abi is required")
		} else {
			cli.Assert(contractAbi != "", quiet, "--abi is required (if not present in ENS)")
			abi, err = contractParseAbi(contractAbi)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse ABI %s", contractAbi))
		}

		openBracketPos := strings.Index(contractSendCall, "(")
		cli.Assert(openBracketPos != -1, quiet, fmt.Sprintf("Missing open bracket in call %s", contractSendCall))
		closeBracketPos := strings.LastIndex(contractSendCall, ")")
		cli.Assert(closeBracketPos != -1, quiet, fmt.Sprintf("Missing close bracket in call %s", contractSendCall))

		methodName := contractSendCall[0:openBracketPos]

		var contractSendArgs []string
		if openBracketPos+1 != closeBracketPos {
			parser := csv.NewReader(strings.NewReader(contractSendCall[openBracketPos+1 : closeBracketPos]))
			contractSendArgs, err = parser.Read()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse arguments for %s", contractSendCall))
		}

		method, exists := abi.Methods[methodName]
		cli.Assert(exists, quiet, fmt.Sprintf("Method %s is unknown", methodName))

		var methodArgs []interface{}
		for i, input := range method.Inputs {
			val, err := contractStringToValue(input.Type, contractSendArgs[i])
			cli.ErrCheck(err, quiet, "Failed to decode argument")
			methodArgs = append(methodArgs, val)
		}

		data, err := abi.Pack(methodName, methodArgs...)
		cli.ErrCheck(err, quiet, "Failed to convert arguments")

		cli.Assert(contractStr != "", quiet, "--contract is required")
		contractAddress, err := ens.Resolve(client, contractStr)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve contract address %s", contractStr))

		amount := big.NewInt(0)
		if contractDeployAmount != "" {
			amount, err = etherutils.StringToWei(contractSendAmount)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid amount %s", contractSendAmount))
		}

		// Create and sign the transaction
		signedTx, err := createSignedTransaction(fromAddress, &contractAddress, amount, gasLimit, data)
		cli.ErrCheck(err, quiet, "Failed to create contract method transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			ctx, cancel := localContext()
			defer cancel()
			err = client.SendTransaction(ctx, signedTx)
			cli.ErrCheck(err, quiet, "Failed to send contract method transaction")

			log.WithFields(log.Fields{
				"group":         "contract",
				"command":       "send",
				"data":          "0x" + hex.EncodeToString(data),
				"from":          fromAddress.Hex(),
				"contract":      contractAddress.Hex(),
				"amount":        amount.String(),
				"networkid":     chainID,
				"gas":           signedTx.Gas().String(),
				"gasprice":      signedTx.GasPrice().String(),
				"transactionid": signedTx.Hash().Hex(),
			}).Info("success")

			if quiet {
				os.Exit(0)
			}
			fmt.Println(signedTx.Hash().Hex())
		}
	},
}

func init() {
	contractCmd.AddCommand(contractSendCmd)
	contractFlags(contractSendCmd)
	contractSendCmd.Flags().StringVar(&contractSendAmount, "amount", "", "Amount of Ether to send with the contract method")
	contractSendCmd.Flags().StringVar(&contractSendFromAddress, "from", "", "Address from which to call the contract function")
	contractSendCmd.Flags().StringVar(&contractSendCall, "call", "", "Contract function to call")
	contractSendCmd.Flags().StringVar(&contractSendReturns, "returns", "", "Comma-separated return types")
	addTransactionFlags(contractSendCmd, "Passphrase for the address from which to send the contract transaction")
}
