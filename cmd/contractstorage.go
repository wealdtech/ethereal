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
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
)

var contractStorageFromAddress string
var contractStorageCall string
var contractStorageReturns string
var contractStorageKey string

// contractStorageCmd represents the contract storage command
var contractStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Obtain details about a contract's storage",
	Long: `Call a contract method against a local node (not transmitting to the blockchain).  For example:

   ethereal contract storage --contract=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07 --abi="./erc20.abi" --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --call="totalSupply()"

In quiet mode this will return 0 if the contract is successfully called, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {

		cli.Assert(contractStr != "", quiet, "--contract is required")
		contractAddress, err := ens.Resolve(client, contractStr)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve contract address %s", contractStr))

		ctx, cancel := localContext()
		defer cancel()
		storage, err := client.StorageAt(ctx, contractAddress, common.HexToHash(contractStorageKey), nil)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain storage for contract %s", contractStr))
		fmt.Println(hex.EncodeToString(storage))

		os.Exit(0)

		cli.Assert(contractStorageFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractStorageFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractStorageFromAddress))

		// We need to have 'call' and 'abi'
		cli.Assert(contractStorageCall != "", quiet, "--call is required")

		var abi abi.ABI
		if contractAbi == "" {
			// TODO See if we can fetch the ABI from ENS
			cli.Err(quiet, "--abi is required")
		} else {
			cli.Assert(contractAbi != "", quiet, "--abi is required (if not present in ENS)")
			abi, err = contractParseAbi(contractAbi)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse ABI %s", contractAbi))
		}

		openBracketPos := strings.Index(contractStorageCall, "(")
		cli.Assert(openBracketPos != -1, quiet, fmt.Sprintf("Missing open bracket in call %s", contractStorageCall))
		closeBracketPos := strings.LastIndex(contractStorageCall, ")")
		cli.Assert(closeBracketPos != -1, quiet, fmt.Sprintf("Missing close bracket in call %s", contractStorageCall))

		methodName := contractStorageCall[0:openBracketPos]

		var contractStorageArgs []string
		if openBracketPos+1 != closeBracketPos {
			parser := csv.NewReader(strings.NewReader(contractStorageCall[openBracketPos+1 : closeBracketPos]))
			contractStorageArgs, err = parser.Read()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse arguments for %s", contractStorageCall))
		}

		method, exists := abi.Methods[methodName]
		cli.Assert(exists, quiet, fmt.Sprintf("Method %s is unknown", methodName))

		var methodArgs []interface{}
		for i, input := range method.Inputs {
			val, err := contractStringToValue(input.Type, contractStorageArgs[i])
			cli.ErrCheck(err, quiet, "Failed to decode argument")
			outputIf(verbose, fmt.Sprintf("input %d is %v (%v)", i, val, reflect.TypeOf(val)))
			methodArgs = append(methodArgs, val)
		}

		data, err := abi.Pack(methodName, methodArgs...)
		cli.ErrCheck(err, quiet, "Failed to convert arguments")

		// Make the call
		msg := ethereum.CallMsg{
			From: fromAddress,
			To:   &contractAddress,
			Data: data,
		}
		ctx, cancel = localContext()
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
	contractCmd.AddCommand(contractStorageCmd)
	contractFlags(contractStorageCmd)
	contractStorageCmd.Flags().StringVar(&contractStorageFromAddress, "from", "", "Address from which to call the contract method")
	contractStorageCmd.Flags().StringVar(&contractStorageCall, "call", "", "Contract method to call")
	contractStorageCmd.Flags().StringVar(&contractStorageReturns, "returns", "", "Comma-separated return types")
	contractStorageCmd.Flags().StringVar(&contractStorageKey, "key", "", "Storage key")
}
