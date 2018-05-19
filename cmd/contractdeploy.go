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

	etherutils "github.com/orinocopay/go-etherutils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
)

var contractDeployFromAddress string
var contractDeployConstructor string
var contractDeployData string
var contractDeployAmount string

// contractDeployCmd represents the contract deploy command
var contractDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a contract",
	Long: `Deploy a contract.  For example:

   ethereal contract deploy --data=0x606060...430029 --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

where data is the hex string of the contract binary.  If the contract constructor requires arguments the both the ABI and the constructor are required, for example:

   ethereal contract deploy --data=0x606060...430029 --abi='./MyContract.abi' --constructor='constructor(1,2,3)' --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

In quiet mode this will return 0 if the contract creation transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(contractDeployFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractDeployFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractDeployFromAddress))

		cli.Assert(contractDeployData != "", quiet, "--data is required")

		var data []byte
		data, err = hex.DecodeString(strings.TrimPrefix(contractDeployData, "0x"))
		cli.ErrCheck(err, quiet, "Failed to decode data")
		// Add construcor arguments if present
		if contractAbi != "" {
			cli.Assert(contractDeployConstructor != "", quiet, "Constructor required if ABI is present")

			abi, err := contractParseAbi(contractAbi)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse ABI %s", contractAbi))

			openBracketPos := strings.Index(contractDeployConstructor, "(")
			cli.Assert(openBracketPos != -1, quiet, fmt.Sprintf("Missing open bracket in call %s", contractDeployConstructor))
			closeBracketPos := strings.LastIndex(contractDeployConstructor, ")")
			cli.Assert(closeBracketPos != -1, quiet, fmt.Sprintf("Missing close bracket in call %s", contractDeployConstructor))

			var contractDeployArgs []string
			if openBracketPos+1 != closeBracketPos {
				parser := csv.NewReader(strings.NewReader(contractDeployConstructor[openBracketPos+1 : closeBracketPos]))
				contractDeployArgs, err = parser.Read()
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse arguments for %s", contractDeployConstructor))
			}

			method := abi.Constructor

			var methodArgs []interface{}
			for i, input := range method.Inputs {
				val, err := contractStringToValue(input.Type, contractDeployArgs[i])
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to decode argument %s", contractDeployArgs[i]))
				methodArgs = append(methodArgs, val)
			}

			argData, err := abi.Pack("", methodArgs...)
			cli.ErrCheck(err, quiet, "Failed to pack arguments")
			data = append(data, argData...)
		}

		amount := big.NewInt(0)
		if contractDeployAmount != "" {
			amount, err = etherutils.StringToWei(contractDeployAmount)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid amount %s", contractDeployAmount))
		}

		// Create and sign the transaction
		signedTx, err := createSignedTransaction(fromAddress, nil, amount, gasLimit, data)
		cli.ErrCheck(err, quiet, "Failed to create contract deployment transaction")

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
			cli.ErrCheck(err, quiet, "Failed to send contract deployment transaction")

			log.WithFields(log.Fields{
				"group":         "contract",
				"command":       "deploy",
				"data":          "0x" + hex.EncodeToString(data),
				"from":          fromAddress.Hex(),
				"amount":        amount.String(),
				"networkid":     chainID,
				"gas":           signedTx.Gas(),
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
	contractCmd.AddCommand(contractDeployCmd)
	contractFlags(contractDeployCmd)
	contractDeployCmd.Flags().StringVar(&contractDeployAmount, "amount", "", "Amount of Ether to send with the contract deployment")
	contractDeployCmd.Flags().StringVar(&contractDeployConstructor, "constructor", "", "Constructor invocation (if required)")
	contractDeployCmd.Flags().StringVar(&contractDeployData, "data", "", "Contract data (as a hex string)")
	contractDeployCmd.Flags().StringVar(&contractDeployFromAddress, "from", "", "Address from which to deploy the contract")
	addTransactionFlags(contractDeployCmd, "Passphrase for the address from which to deploy the conract")
}
