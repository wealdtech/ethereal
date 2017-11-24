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
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var contractStr string
var contractAbi string

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
	cmd.Flags().StringVar(&contractAbi, "abi", "", "ABI, or path to ABI, for the contract ")
}

func contractParseAbi(input string) (output abi.ABI, err error) {
	var reader io.Reader
	if strings.Contains(input, string(filepath.Separator)) {
		// ABI value is a path
		reader, err = os.Open(input)
		if err != nil {
			return
		}
	} else {
		reader = strings.NewReader(input)
	}
	return abi.JSON(reader)
}

func contractUnpack(abi abi.ABI, name string, data []byte) (result *[]*interface{}, err error) {
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

func contractStringToValue(arg abi.Argument, val string) (interface{}, error) {
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

func contractValueToString(arg abi.Argument, val interface{}) (string, error) {
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
