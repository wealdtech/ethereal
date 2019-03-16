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
	"fmt"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/util/funcparser"
	ens "github.com/wealdtech/go-ens"
)

var contractCallFromAddress string
var contractCallCall string
var contractCallReturns string

// contractCallCmd represents the contract call command
var contractCallCmd = &cobra.Command{
	Use:   "call",
	Short: "Call a contract method",
	Long: `Call a contract method against a local node (not transmitting to the blockchain).  For example:

   ethereal contract call --contract=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07 --abi="./erc20.abi" --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --call="totalSupply()"

   ethereal contract call --contract=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07 --signature="balanceOf(address)" --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --call="balanceOf(@wealdtech.eth)"

In quiet mode this will return 0 if the contract is successfully called, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {

		cli.Assert(contractCallFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractCallFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractCallFromAddress))

		// We need to have 'call'
		cli.Assert(contractCallCall != "", quiet, "--call is required")

		contract := parseContract("")
		method, methodArgs, err := funcparser.ParseCall(client, contract, contractCallCall)
		cli.ErrCheck(err, quiet, "Failed to parse call")

		data, err := contract.Abi.Pack(method.Name, methodArgs...)
		cli.ErrCheck(err, quiet, "Failed to convert arguments")
		outputIf(verbose, fmt.Sprintf("Data is %x", data))

		cli.Assert(contractStr != "", quiet, "--contract is required")
		contractAddress, err := ens.Resolve(client, contractStr)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve contract address %s", contractStr))

		// Make the call
		msg := ethereum.CallMsg{
			From: fromAddress,
			To:   &contractAddress,
			Data: data,
		}
		ctx, cancel := localContext()
		defer cancel()
		result, err := client.CallContract(ctx, msg, nil)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to call %s", method.Name))
		if len(method.Outputs) == 0 {
			// No output
			os.Exit(0)
		}
		cli.Assert(len(result) > 0, quiet, fmt.Sprintf("Call to %s did not return expected data", method.Name))

		if quiet {
			os.Exit(0)
		}

		outputIf(verbose, fmt.Sprintf("Result is %x", []byte(result)))

		var tmp interface{}
		err = contract.Abi.Unpack(&tmp, method.Name, result)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse output of %s: %v", method.Name, err))

		outputs := make([]interface{}, len(method.Outputs))
		if len(method.Outputs) == 1 {
			outputs[0] = tmp
		} else {
			for i, x := range tmp.([]interface{}) {
				outputs[i] = x
			}
		}

		results := []string{}
		for i := range outputs {
			val, err := contractValueToString(method.Outputs[i].Type, outputs[i])
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to turn value %v in to suitable output", outputs[i]))
			results = append(results, val)
		}

		// Output the result
		fmt.Printf("%s\n", strings.Join(results, ","))
	},
}

func init() {
	contractCmd.AddCommand(contractCallCmd)
	contractFlags(contractCallCmd)
	contractCallCmd.Flags().StringVar(&contractCallFromAddress, "from", "", "Address from which to call the contract method")
	contractCallCmd.Flags().StringVar(&contractCallCall, "call", "", "Contract method to call")
	contractCallCmd.Flags().StringVar(&contractCallReturns, "returns", "", "Comma-separated return types")
}
