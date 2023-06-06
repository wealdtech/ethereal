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
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util/funcparser"
)

var (
	signatureDataStr string
	signatureTypes   string
	signatureNoHash  bool
	signaturePacked  bool
)

// signatureCmd represents the signature command.
var signatureCmd = &cobra.Command{
	Use:     "signature",
	Aliases: []string{"sig"},
	Short:   "Manage signatures",
	Long:    `Sign and verify information.`,
}

func generateDataHash() []byte {
	var data []byte
	if signatureTypes == "" {
		// No types; might be a hex string or a non-hex string.
		bytes, err := hex.DecodeString(strings.TrimPrefix(signatureDataStr, "0x"))
		if err != nil {
			// Not a hex string; keep it as a normal string.
			data = []byte(signatureDataStr)
		} else {
			// Is a hex string.
			data = bytes
		}
	} else {
		// Types are present; Ethereum types.
		args, vals := argumentsAndValues(signatureDataStr, signatureTypes)
		var err error
		if signaturePacked {
			hexStr := ""
			for i := range args {
				switch args[i].Type.T {
				case abi.IntTy, abi.UintTy:
					hexStr += fmt.Sprintf(fmt.Sprintf("%%0%dx", args[i].Type.Size/4), vals[i])
				case abi.BoolTy:
					if vals[i].(bool) {
						hexStr += "01"
					} else {
						hexStr += "00"
					}
				case abi.StringTy:
					hexStr += fmt.Sprintf("%x", vals[i].(string))
				case abi.FixedBytesTy:
					hexStr += fmt.Sprintf(fmt.Sprintf("%%0%dx", args[i].Type.Size*2), vals[i])
				case abi.BytesTy:
					hexStr += fmt.Sprintf("%x", vals[i].([]byte))
				case abi.AddressTy:
					hexStr += fmt.Sprintf("%040x", vals[i])
				default:
					cli.Err(quiet, fmt.Sprintf("Unhandled type %v\n", args[i].Type))
				}
			}
			data, err = hex.DecodeString(hexStr)
			cli.ErrCheck(err, quiet, "Failed to pack data")
		} else {
			data, err = args.Pack(vals...)
			cli.ErrCheck(err, quiet, "Failed to pack data")
		}
	}
	outputIf(verbose, fmt.Sprintf("Data is %x", data))

	// Hash if required.
	if !signatureNoHash {
		// Hash the data.
		data = crypto.Keccak256(data)
		outputIf(verbose, fmt.Sprintf("Hashed data is %x", data))
	}
	buffer := make([]byte, 0)
	buffer = append(buffer, []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data)))...)
	buffer = append(buffer, data...)
	outputIf(verbose, fmt.Sprintf("Data to sign is %x", buffer))
	return crypto.Keccak256(buffer)
}

func argumentsAndValues(items string, types string) (abi.Arguments, []interface{}) {
	parser := csv.NewReader(strings.NewReader(items))
	dataItems, err := parser.Read()
	cli.ErrCheck(err, quiet, "Failed to parse data")

	cli.Assert(types != "", quiet, "--types is required")
	parser = csv.NewReader(strings.NewReader(types))
	dataTypes, err := parser.Read()
	cli.ErrCheck(err, quiet, "Failed to parse data types")

	// Lean on ABI function parsing even though we don't have an ABI...
	vals := make([]interface{}, 0)
	arguments := abi.Arguments{}
	for i := range dataTypes {
		dataType, err := abi.NewType(dataTypes[i], "", nil)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Unknown data type %s", dataTypes[i]))
		argument := abi.Argument{
			Type: dataType,
		}
		arguments = append(arguments, argument)
		val, err := funcparser.StrTo(&dataType, dataItems[i])
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to decode argument %s", dataItems[i]))
		vals = append(vals, val)
	}

	return arguments, vals
}

func init() {
	RootCmd.AddCommand(signatureCmd)
}

func signatureFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&signatureDataStr, "data", "", "the data")
	cmd.Flags().StringVar(&signatureTypes, "types", "", "Comma-separated list of data types")
	cmd.Flags().BoolVar(&signatureNoHash, "nohash", false, "do not hash the message prior to signing")
	cmd.Flags().BoolVar(&signaturePacked, "packed", false, "use Solidity packed encoding")
}
