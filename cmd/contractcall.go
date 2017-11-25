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
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
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

In quiet mode this will return 0 if the contract is successfully called, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {

		cli.Assert(contractCallFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractCallFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractCallFromAddress))

		// We need to have 'call' and 'abi'
		cli.Assert(contractCallCall != "", quiet, "--call is required")

		var abi abi.ABI
		if contractAbi == "" {
			// TODO See if we can fetch the ABI from ENS
			cli.Err(quiet, "--abi is required")
		} else {
			cli.Assert(contractAbi != "", quiet, "--abi is required (if not present in ENS)")
			abi, err = contractParseAbi(contractAbi)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse ABI %s", contractAbi))
		}

		openBracketPos := strings.Index(contractCallCall, "(")
		cli.Assert(openBracketPos != -1, quiet, fmt.Sprintf("Missing open bracket in call %s", contractCallCall))
		closeBracketPos := strings.LastIndex(contractCallCall, ")")
		cli.Assert(closeBracketPos != -1, quiet, fmt.Sprintf("Missing close bracket in call %s", contractCallCall))

		methodName := contractCallCall[0:openBracketPos]

		var contractCallArgs []string
		if openBracketPos+1 != closeBracketPos {
			parser := csv.NewReader(strings.NewReader(contractCallCall[openBracketPos+1 : closeBracketPos]))
			contractCallArgs, err = parser.Read()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse arguments for %s", contractCallCall))
		}

		method, exists := abi.Methods[methodName]
		cli.Assert(exists, quiet, fmt.Sprintf("Method %s is unknown", methodName))

		var methodArgs []interface{}
		for i, input := range method.Inputs {
			val, err := contractStringToValue(input.Type, contractCallArgs[i])
			cli.ErrCheck(err, quiet, "Failed to decode argument")
			fmt.Printf("val is %v\n", val)
			fmt.Printf("val type is %v\n", reflect.TypeOf(val))
			methodArgs = append(methodArgs, val)
		}

		fmt.Printf("methodName is %v\n", methodName)
		fmt.Printf("methodArgs is %v\n", methodArgs)
		data, err := abi.Pack(methodName, methodArgs...)
		cli.ErrCheck(err, quiet, "Failed to convert arguments")
		fmt.Printf("data is %v\n", data)

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
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to call contract %s", methodName))
		cli.Assert(len(result) > 0, quiet, fmt.Sprintf("Call to %s did not return any data", methodName))

		if quiet {
			os.Exit(0)
		}

		abiOutput, err := contractUnpack(abi, methodName, []byte(result))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid ABI for %s in ABI", methodName))
		results := []string{}
		for i, _ := range *abiOutput {
			val, err := contractValueToString(method.Outputs[i].Type, *((*abiOutput)[i]))
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to turn value %v in to suitable output", *((*abiOutput)[i])))
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
