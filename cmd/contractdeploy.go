// Copyright © 2017-2019 Weald Technology Trading
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

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	"github.com/wealdtech/ethereal/v2/util/funcparser"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	contractDeployFromAddress string
	contractDeployConstructor string
	contractDeployData        string
	contractDeployAmount      string
	contractDeployRepeat      int
)

// contractDeployCmd represents the contract deploy command.
var contractDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a contract",
	Long: `Deploy a contract.  For example:

   ethereal contract deploy --data=0x606060...430029 --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

where data is the hex string of the contract binary.  If the contract constructor requires arguments the both the ABI and the constructor are required, for example:

   ethereal contract deploy --data=0x606060...430029 --abi='./MyContract.abi' --constructor='constructor(1,2,3)' --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

Usually the easiest way to deploy a contract is to use the combined JSON output of solc.  In this situation deploying the contract might be:

   ethereal contract deploy --json='./MyContract.json' --constructor='constructor(1,2,3') --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(contractDeployFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(contractDeployFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractDeployFromAddress))
		cli.Assert(contractDeployData != "" || contractJSON != "", quiet, "either --data or --json is required")

		contract := parseContract(contractDeployData)
		cli.Assert(len(contract.Binary) > 0, quiet, "failed to obtain contract binary data")
		if contractDeployConstructor != "" {
			_, constructorArgs, err := funcparser.ParseCall(c.Client(), contract, contractDeployConstructor)
			cli.ErrCheck(err, quiet, "Failed to parse constructor")

			argData, err := contract.Abi.Pack("", constructorArgs...)
			cli.ErrCheck(err, quiet, "Failed to convert arguments")
			outputIf(verbose, fmt.Sprintf("Constructor data is %x", argData))
			contract.Binary = append(contract.Binary, argData...)
		}

		amount := big.NewInt(0)
		if contractDeployAmount != "" {
			amount, err = string2eth.StringToWei(contractDeployAmount)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid amount %s", contractDeployAmount))
		}

		var gasLimit *uint64
		limit := uint64(viper.GetInt64("gaslimit"))
		if limit > 0 {
			gasLimit = &limit
		}

		var signedTx *types.Transaction
		for i := 0; i < contractDeployRepeat; i++ {
			// Create and sign the transaction.
			signedTx, err = c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
				From:     fromAddress,
				Value:    amount,
				GasLimit: gasLimit,
				Data:     contract.Binary,
			})
			cli.ErrCheck(err, quiet, "Failed to create contract deployment transaction")
			outputIf(verbose, fmt.Sprintf("Transaction data is %x", signedTx.Data()))
			outputIf(verbose, fmt.Sprintf("Transaction data size is %d", len(signedTx.Data())))

			if offline {
				if !quiet {
					buf := new(bytes.Buffer)
					cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
					fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
				}
				os.Exit(exitSuccess)
			}
			err = c.SendTransaction(context.Background(), signedTx)
			cli.ErrCheck(err, quiet, "Failed to send transaction")
			logTransaction(signedTx, log.Fields{
				"group":   "contract",
				"command": "deploy",
			})
		}

		// Wait for the last transaction if requested.
		handleSubmittedTransaction(signedTx, nil, false)
	},
}

func init() {
	contractCmd.AddCommand(contractDeployCmd)
	contractFlags(contractDeployCmd)
	contractDeployCmd.Flags().StringVar(&contractDeployAmount, "amount", "", "Amount of Ether to send with the contract deployment")
	contractDeployCmd.Flags().StringVar(&contractDeployConstructor, "constructor", "", "Constructor invocation (if required)")
	contractDeployCmd.Flags().StringVar(&contractDeployData, "data", "", "Contract data (as a hex string)")
	contractDeployCmd.Flags().StringVar(&contractDeployFromAddress, "from", "", "Address from which to deploy the contract")
	contractDeployCmd.Flags().IntVar(&contractDeployRepeat, "repeat", 1, "Number of times to repeat sending the transaction (incrementing the nonce each time)")
	addTransactionFlags(contractDeployCmd, "Passphrase for the address from which to deploy the conract")
}
