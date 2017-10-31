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
	"fmt"
	"math/big"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
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

		if quiet {
			os.Exit(0)
		}

		// Format the result
		cli.Assert(contractCallReturns != "", quiet, "--returns is required")
		returnTypes := strings.Split(contractCallReturns, ",")
		position := 0
		for _, returnType := range returnTypes {
			switch returnType {
			case "uint":
			case "uint256":
				item := big.NewInt(0)
				item.SetBytes(result[position : position+32])
				position += 32
				fmt.Println(item.String())
			case "string":
				offset := binary.BigEndian.Uint64(result[position+24 : position+32])
				length := binary.BigEndian.Uint64(result[offset+24 : offset+32])
				item := string(result[offset+32 : offset+32+length])
				position += 32
				fmt.Println(item)
			default:
				cli.Err(quiet, fmt.Sprintf("Unknown return type %s", returnType))
			}
		}
	},
}

func init() {
	contractCmd.AddCommand(contractCallCmd)
	contractFlags(contractCallCmd)
	contractCallCmd.Flags().StringVar(&contractCallFromAddress, "from", "", "Address from which to call the contract method")
	contractCallCmd.Flags().StringVar(&contractCallSignature, "signature", "", "Signature of the contract method to call")
	contractCallCmd.Flags().StringVar(&contractCallReturns, "returns", "", "Comma-separated return types")
}
