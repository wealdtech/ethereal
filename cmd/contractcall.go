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
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

var contractCallFromAddress string
var contractCallSignature string
var contractCallReturns string

// contractCallCmd represents the contract send command
var contractCallCmd = &cobra.Command{
	Use:   "call",
	Short: "Call a contract method",
	Long: `Call a contract method.  For example:

    ethereal contract call --contract=0x2ab7150Bba7D5F181b3aF5623e52b15bB1054845 --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --method="totalSupply()"

In quiet mode this will return 0 if the contract is successfully called, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {

		cli.Assert(contractCallFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractCallFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractCallFromAddress))

		data := make([]byte, 4)
		if contractCallSignature != "" {
			sig := make([]byte, 32)
			sha := sha3.NewKeccak256()
			sha.Write([]byte(contractCallSignature))
			sha.Sum(sig[:0])
			copy(data, sig)
		}

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
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to call contract %s", contractCallSignature))
		cli.Assert(len(result) > 0, quiet, fmt.Sprintf("Call to %s did not return any data", contractCallSignature))

		if quiet {
			os.Exit(0)
		}

		// Format the result
		cli.Assert(contractCallReturns != "", quiet, "--returns is required")
		returnTypes := strings.Split(contractCallReturns, ",")
		results := []string{}
		position := 0
		for _, returnType := range returnTypes {
			cli.Assert(position < len(result), quiet, "Not enough data returned for all expected return values")
			if strings.HasPrefix(returnType, "uint") {
				var bytes int
				if returnType == "uint" {
					// Alias for uint256
					bytes = 32
				} else {
					bytes, err = strconv.Atoi(returnType[4:])
					cli.ErrCheck(err, quiet, fmt.Sprintf("Unknown return type %s", returnType))
					bytes /= 8
				}
				item := big.NewInt(0)
				item.SetBytes(result[position+32-bytes : position+32])
				position += 32
				results = append(results, item.String())
			} else if strings.HasPrefix(returnType, "int") {
				// TODO how to handle negative numbers
				var bytes int
				if returnType == "int" {
					// Alias for int256
					bytes = 32
				} else {
					bytes, err = strconv.Atoi(returnType[4:])
					cli.ErrCheck(err, quiet, fmt.Sprintf("Unknown return type %s", returnType))
					bytes /= 8
				}
				item := big.NewInt(0)
				item.SetBytes(result[position+32-bytes : position+32])
				position += 32
				results = append(results, item.String())
			} else if returnType == "string" {
				offset := binary.BigEndian.Uint64(result[position+24 : position+32])
				length := binary.BigEndian.Uint64(result[offset+24 : offset+32])
				item := string(result[offset+32 : offset+32+length])
				position += 32
				results = append(results, fmt.Sprintf("\"%s\"", item))
			} else if returnType == "address" {
				item := common.BytesToAddress(result[position+12 : position+32])
				position += 32
				results = append(results, item.Hex())
			} else if returnType == "hash" {
				item := result[position : position+32]
				position += 32
				results = append(results, fmt.Sprintf("0x%s", hex.EncodeToString(item)))
			} else if returnType == "bool" {
				item := result[position+31]
				position += 32
				if item == 1 {
					results = append(results, "true")
				} else {
					results = append(results, "false")
				}
			} else {
				cli.Err(quiet, fmt.Sprintf("Unknown return type %s\nData is %s", returnType, hex.EncodeToString(result)))
				results = append(results, "?")
			}
			// TODO fixed, ufixed, bytes, function, <type>[M], bytes, <type>[], ()
		}
		fmt.Printf("%s\n", strings.Join(results, ","))
	},
}

func init() {
	contractCmd.AddCommand(contractCallCmd)
	contractFlags(contractCallCmd)
	contractCallCmd.Flags().StringVar(&contractCallFromAddress, "from", "", "Address from which to call the contract method")
	contractCallCmd.Flags().StringVar(&contractCallSignature, "signature", "", "Signature of the contract method to call")
	contractCallCmd.Flags().StringVar(&contractCallReturns, "returns", "", "Comma-separated return types")
}
