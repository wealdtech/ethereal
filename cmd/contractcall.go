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
	"math/big"
	"os"
	"reflect"
	"strings"
	"unsafe"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

var contractCallFromAddress string
var contractCallCall string
var contractCallReturns string

// contractCallCmd represents the contract send command
var contractCallCmd = &cobra.Command{
	Use:   "call",
	Short: "Call a contract method",
	Long: `Call a contract method.  For example:

   ethereal contract call --contract=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07 --abi="./erc20.abi" --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --call="totalSupply()"

In quiet mode this will return 0 if the contract is successfully called, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {

		cli.Assert(contractCallFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, contractCallFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", contractCallFromAddress))

		// We need to have 'call', and either 'abi' or 'signature'
		cli.Assert(contractCallCall != "", quiet, "--call is required")
		// TODO handle 'signature'

		var abi abi.ABI
		if contractAbi == "" {
			// See if we can fetch the ABI from ENS, one day
			cli.Err(quiet, "--abi is required")
		} else {
			cli.Assert(contractAbi != "", quiet, "--abi is required (if not present in ENS)")
			abi, err = parseAbi(contractAbi)
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
			val, err := stringToValue(input, contractCallArgs[i])
			cli.ErrCheck(err, quiet, "Failed to decode argument")
			methodArgs = append(methodArgs, val)
		}

		data, err := abi.Pack(methodName, methodArgs...)

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

		abiOutput, err := unpack(abi, methodName, []byte(result))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid ABI for %s in ABI", "RewardEnd"))
		results := []string{}
		for i, _ := range *abiOutput {
			val, err := valueToString(method.Outputs[i], *((*abiOutput)[i]))
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to turn value %v in to suitable output", *((*abiOutput)[i])))
			results = append(results, val)
		}

		// Output the result
		fmt.Printf("%s\n", strings.Join(results, ","))
	},
}

func unpack(abi abi.ABI, name string, data []byte) (result *[]*interface{}, err error) {
	method, exists := abi.Methods[name]
	if !exists {
		return nil, fmt.Errorf("The method %s does not exist", name)
	}

	var res []*interface{}
	result = &res
	if len(method.Outputs) == 0 {
		return
	} else if len(method.Outputs) == 1 {
		output := reflect.New(method.Outputs[0].Type.Type.Elem()).Elem().Interface()
		err = abi.Unpack(&output, name, data)
		res = append(res, &output)
	} else {
		for i, _ := range method.Outputs {
			output := reflect.New(method.Outputs[i].Type.Type.Elem()).Elem().Interface()
			res = append(res, &output)
		}
		err = abi.Unpack(&res, name, data)
	}
	return
}

func stringToValue(arg abi.Argument, val string) (interface{}, error) {
	switch arg.Type.T {
	case abi.IntTy:
		res := big.NewInt(0)
		res, err := res.SetString(val, 10)
		if err {
			return nil, fmt.Errorf("Bad integer %s", val)
		}
		return res, nil
	case abi.UintTy:
		res := big.NewInt(0)
		res, err := res.SetString(val, 10)
		if err {
			return nil, fmt.Errorf("Bad integer %s", val)
		}
		return res, nil
	case abi.BoolTy:
		if val == "true" || val == "True" || val == "1" {
			return true, nil
		} else {
			return false, nil
		}
	case abi.StringTy:
		return val, nil
	case abi.SliceTy:
		return nil, fmt.Errorf("Unhandled type slice (%s)", arg.Type.T)
	case abi.ArrayTy:
		return nil, fmt.Errorf("Unhandled type array (%s)", arg.Type.T)
	case abi.AddressTy:
		return common.HexToAddress(val), nil
	case abi.FixedBytesTy:
		slice := make([]byte, arg.Type.Size)
		var decoded []byte
		if strings.HasPrefix(val, "0x") {
			decoded, err = hex.DecodeString(val[2:])
		} else {
			decoded, err = hex.DecodeString(val)
		}
		if err == nil {
			copy(slice[arg.Type.Size-len(decoded):arg.Type.Size], decoded)
		}
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
		switch arg.Type.Size {
		case 1:
			return *(*[1]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 2:
			return *(*[2]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 3:
			return *(*[3]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 4:
			return *(*[4]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 5:
			return *(*[5]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 6:
			return *(*[6]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 7:
			return *(*[7]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 8:
			return *(*[8]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 9:
			return *(*[9]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 10:
			return *(*[10]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 11:
			return *(*[11]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 12:
			return *(*[12]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 13:
			return *(*[13]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 14:
			return *(*[14]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 15:
			return *(*[15]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 16:
			return *(*[16]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 17:
			return *(*[17]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 18:
			return *(*[18]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 19:
			return *(*[19]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 20:
			return *(*[20]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 21:
			return *(*[21]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 22:
			return *(*[22]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 23:
			return *(*[23]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 24:
			return *(*[24]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 25:
			return *(*[25]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 26:
			return *(*[26]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 27:
			return *(*[27]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 28:
			return *(*[28]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 29:
			return *(*[29]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 30:
			return *(*[30]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 31:
			return *(*[31]uint8)(unsafe.Pointer(hdr.Data)), nil
		case 32:
			return *(*[32]uint8)(unsafe.Pointer(hdr.Data)), nil
		}
		return nil, fmt.Errorf("Invalid byte size %d", arg.Type.Size)
	case abi.BytesTy:
		if strings.HasPrefix(val, "0x") {
			return hex.DecodeString(val[2:])
		} else {
			return hex.DecodeString(val)
		}
	case abi.HashTy:
		return common.HexToHash(val), nil
	case abi.FixedPointTy:
		return nil, fmt.Errorf("Unhandled type %s", arg.Type.T)
	case abi.FunctionTy:
		return nil, fmt.Errorf("Unhandled type %s", arg.Type.T)
	default:
		return nil, fmt.Errorf("Unknown type %s", arg.Type.T)
	}
}

func valueToString(arg abi.Argument, val interface{}) (string, error) {
	switch arg.Type.T {
	case abi.IntTy:
		return val.(*big.Int).String(), nil
	case abi.UintTy:
		return val.(*big.Int).String(), nil
	case abi.BoolTy:
		if val.(bool) == true {
			return "true", nil
		} else {
			return "false", nil
		}
	case abi.StringTy:
		return val.(string), nil
	case abi.SliceTy:
		return "", fmt.Errorf("Unhandled type %s", arg.Type.T)
	case abi.ArrayTy:
		return "", fmt.Errorf("Unhandled type %s", arg.Type.T)
	case abi.AddressTy:
		return val.(common.Address).Hex(), nil
	case abi.FixedBytesTy:
		return fmt.Sprintf("0x%s", hex.EncodeToString(val.([]byte))), nil
	case abi.BytesTy:
		return fmt.Sprintf("0x%s", hex.EncodeToString(val.([]byte))), nil
	case abi.HashTy:
		return val.(common.Hash).Hex(), nil
	case abi.FixedPointTy:
		return "", fmt.Errorf("Unhandled type %s", arg.Type.T)
	case abi.FunctionTy:
		return "", fmt.Errorf("Unhandled type %s", arg.Type.T)
	default:
		return "", fmt.Errorf("Unknown type %s", arg.Type.T)
	}
}

func init() {
	contractCmd.AddCommand(contractCallCmd)
	contractFlags(contractCallCmd)
	contractCallCmd.Flags().StringVar(&contractCallFromAddress, "from", "", "Address from which to call the contract method")
	contractCallCmd.Flags().StringVar(&contractCallCall, "call", "", "Contract method to call")
	contractCallCmd.Flags().StringVar(&contractCallReturns, "returns", "", "Comma-separated return types")
}
