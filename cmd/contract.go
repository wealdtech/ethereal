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
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/util"
	ens "github.com/wealdtech/go-ens"
)

var contractStr string
var contractAbi string
var contractFunction string
var contractJSON string
var contractName string

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Manage contracts",
	Long:  `Call contracts directly.`,
}

func init() {
	RootCmd.AddCommand(contractCmd)
}

func contractFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&contractStr, "contract", "", "address of the contract")
	cmd.Flags().StringVar(&contractAbi, "abi", "", "ABI, or path to ABI, for the contract")
	cmd.Flags().StringVar(&contractFunction, "function", "", "Signature of function")
	cmd.Flags().StringVar(&contractJSON, "json", "", "JSON, or path to JSON, for the contract as output by solc --combined-json=bin,abi")
	cmd.Flags().StringVar(&contractName, "name", "", "Name of the contract (required when using json)")
}

// parse contract given the information from various flags
func parseContract(binStr string) *util.Contract {
	var contract *util.Contract
	if contractJSON != "" {
		if contractName == "" {
			// Attempt to obtain the contract name from the JSON file
			contractName = strings.Split(filepath.Base(contractJSON), ".")[0]
		}
		contract, err = util.ParseCombinedJSON(contractJSON, contractName)
		cli.ErrCheck(err, quiet, "Failed to parse JSON")
	} else {
		contract = &util.Contract{}

		// Add name if present
		if contractName != "" {
			contract.Name = contractName
		}

		// Add binary if present
		var bin []byte
		bin, err = hex.DecodeString(strings.TrimPrefix(binStr, "0x"))
		cli.ErrCheck(err, quiet, "Failed to decode data")
		contract = &util.Contract{Binary: bin}

		// Add ABI if present either directly or via a function
		if contractAbi != "" {
			abi, err := contractParseAbi(contractAbi)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse ABI %s", contractAbi))
			contract.Abi = abi
		} else if contractFunction != "" {
			abi, err := contractParseFunction(contractFunction)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to parse function %s", contractFunction))
			contract.Abi = *abi
		}
	}
	return contract
}

func contractParseAbi(input string) (output abi.ABI, err error) {
	var reader io.Reader

	if strings.HasPrefix(contractAbi, "[") {
		// ABI is direct
		reader = strings.NewReader(input)
	} else {
		// ABI value is a path
		reader, err = os.Open(input)
		if err != nil {
			return
		}
	}
	return abi.JSON(reader)
}

var intFixRe = regexp.MustCompile(`^([u]?int)($|[^0-9])`)

// contractParseFunction turns a function definition in to an ABI
// function definition is  a string of form "methodName(argtype [argname],...) returns (outputtype [outputname],...)"
func contractParseFunction(input string) (*abi.ABI, error) {

	input = strings.TrimSpace(input)
	bits := strings.Split(input, "(")
	// Method name is part before first "("
	methodName := bits[0]
	// Method arguments are comma-separated values before first ")"
	argsBits := strings.Split(strings.Split(bits[1], ")")[0], ",")
	methodArgs := make([]abi.Argument, len(argsBits))
	for i, argsBit := range argsBits {
		argBits := strings.Split(argsBit, " ")
		argType := argBits[0]
		var argName string
		if len(argBits) > 1 {
			argName = argBits[len(argBits)-1]
		}
		argType = intFixRe.ReplaceAllString(argType, `${1}256${2}`)
		t, err := abi.NewType(argType, nil)
		if err != nil {
			return nil, err
		}
		methodArgs[i] = abi.Argument{
			Name: argName,
			Type: t,
		}
	}
	var methodOutputs []abi.Argument
	if len(bits) > 2 {
		// Method outputs are comma-separated values after last "("
		outputTypes := strings.Split(strings.TrimSuffix(bits[2], ")"), ",")
		methodOutputs = make([]abi.Argument, len(outputTypes))
		for i, outputType := range outputTypes {
			outputType = strings.Split(outputType, " ")[0]
			outputType = intFixRe.ReplaceAllString(outputType, `${1}256${2}`)
			t, err := abi.NewType(outputType, nil)
			if err != nil {
				return nil, err
			}
			methodOutputs[i] = abi.Argument{
				Type: t,
			}
		}
	}

	method := abi.Method{
		Name:    methodName,
		Inputs:  methodArgs,
		Outputs: methodOutputs,
	}

	res := &abi.ABI{
		Methods: make(map[string]abi.Method),
	}
	res.Methods[methodName] = method

	return res, nil
}

func contractValueToString(argType abi.Type, val interface{}) (string, error) {
	switch argType.T {
	case abi.IntTy:
		return fmt.Sprintf("%v", val), nil
	case abi.UintTy:
		return fmt.Sprintf("%v", val), nil
	case abi.BoolTy:
		if val.(bool) == true {
			return "true", nil
		}
		return "false", nil
	case abi.StringTy:
		return val.(string), nil
	case abi.SliceTy:
		res := make([]string, 0)
		arrayVal := reflect.ValueOf(val)
		for i := 0; i < arrayVal.Len(); i++ {
			elemRes, err := contractValueToString(*argType.Elem, arrayVal.Index(i).Interface())
			if err != nil {
				return "", err
			}
			res = append(res, elemRes)
		}
		return "[" + strings.Join(res, ",") + "]", nil
	case abi.ArrayTy:
		res := make([]string, 0)
		arrayVal := reflect.ValueOf(val)
		for i := 0; i < arrayVal.Len(); i++ {
			elemRes, err := contractValueToString(*argType.Elem, arrayVal.Index(i).Interface())
			if err != nil {
				return "", err
			}
			res = append(res, elemRes)
		}
		return "[" + strings.Join(res, ",") + "]", nil
	case abi.AddressTy:
		addr := val.(common.Address)
		return ens.Format(client, &addr), nil
	case abi.FixedBytesTy:
		arrayVal := reflect.ValueOf(val)
		castVal := make([]byte, arrayVal.Len())
		for i := 0; i < arrayVal.Len(); i++ {
			castVal[i] = byte(arrayVal.Index(i).Uint())
		}
		return fmt.Sprintf("0x%s", hex.EncodeToString(castVal)), nil
	case abi.BytesTy:
		return fmt.Sprintf("0x%s", hex.EncodeToString(val.([]byte))), nil
	case abi.HashTy:
		return val.(common.Hash).Hex(), nil
	case abi.FixedPointTy:
		return "", fmt.Errorf("unhandled type %v", argType)
	case abi.FunctionTy:
		return "", fmt.Errorf("unhandled type %v", argType)
	default:
		return "", fmt.Errorf("unknown type %v", argType)
	}
}
