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

func contractStringToValue(argType abi.Type, val string) (interface{}, error) {
	switch argType.T {
	case abi.IntTy:
		res := big.NewInt(0)
		res, success := res.SetString(val, 10)
		if !success {
			return nil, fmt.Errorf("Bad integer %s", val)
		}
		switch argType.Size {
		case 8:
			return int8(res.Uint64()), nil
		case 16:
			return int16(res.Uint64()), nil
		case 32:
			return int32(res.Uint64()), nil
		case 64:
			return int64(res.Uint64()), nil
		default:
			return res, nil
		}
	case abi.UintTy:
		res := big.NewInt(0)
		res, success := res.SetString(val, 10)
		if !success {
			return nil, fmt.Errorf("Bad integer %s", val)
		}
		switch argType.Size {
		case 8:
			return uint8(res.Uint64()), nil
		case 16:
			return uint16(res.Uint64()), nil
		case 32:
			return uint32(res.Uint64()), nil
		case 64:
			return uint64(res.Uint64()), nil
		default:
			return res, nil
		}
	case abi.BoolTy:
		if val == "true" || val == "True" || val == "1" {
			return true, nil
		} else {
			return false, nil
		}
	case abi.StringTy:
		return val, nil
	case abi.SliceTy:
		return nil, fmt.Errorf("Unhandled type slice (%s)", argType)
	case abi.ArrayTy:
		return nil, fmt.Errorf("Unhandled type array (%s)", argType)
	case abi.AddressTy:
		return common.HexToAddress(val), nil
	case abi.FixedBytesTy:
		slice := make([]byte, argType.Size)
		var decoded []byte
		if strings.HasPrefix(val, "0x") {
			decoded, err = hex.DecodeString(val[2:])
		} else {
			decoded, err = hex.DecodeString(val)
		}
		if err == nil {
			copy(slice[argType.Size-len(decoded):argType.Size], decoded)
		}
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
		switch argType.Size {
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
		return nil, fmt.Errorf("Invalid byte size %d", argType.Size)
	case abi.BytesTy:
		if strings.HasPrefix(val, "0x") {
			return hex.DecodeString(val[2:])
		} else {
			return hex.DecodeString(val)
		}
	case abi.HashTy:
		return common.HexToHash(val), nil
	case abi.FixedPointTy:
		return nil, fmt.Errorf("Unhandled type %v", argType)
	case abi.FunctionTy:
		return nil, fmt.Errorf("Unhandled type %v", argType)
	default:
		return nil, fmt.Errorf("Unknown type %v", argType)
	}
}

func contractValueToString(argType abi.Type, val interface{}) (string, error) {
	switch argType.T {
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
		return "", fmt.Errorf("Unhandled type %v", argType)
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
		return val.(common.Address).Hex(), nil
	case abi.FixedBytesTy:
		switch argType.Size {
		case 1:
			castVal := val.([1]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 2:
			castVal := val.([2]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 3:
			castVal := val.([3]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 4:
			castVal := val.([4]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 5:
			castVal := val.([5]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 6:
			castVal := val.([6]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 7:
			castVal := val.([7]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 8:
			castVal := val.([8]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 9:
			castVal := val.([9]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 10:
			castVal := val.([10]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 11:
			castVal := val.([11]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 12:
			castVal := val.([12]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 13:
			castVal := val.([13]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 14:
			castVal := val.([14]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 15:
			castVal := val.([15]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 16:
			castVal := val.([16]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 17:
			castVal := val.([17]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 18:
			castVal := val.([18]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 19:
			castVal := val.([19]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 20:
			castVal := val.([20]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 21:
			castVal := val.([21]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 22:
			castVal := val.([22]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 23:
			castVal := val.([23]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 24:
			castVal := val.([24]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 25:
			castVal := val.([25]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 26:
			castVal := val.([26]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 27:
			castVal := val.([27]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 28:
			castVal := val.([28]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 29:
			castVal := val.([29]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 30:
			castVal := val.([30]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 31:
			castVal := val.([32]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		case 32:
			castVal := val.([32]byte)
			return fmt.Sprintf("0x%s", hex.EncodeToString(castVal[:])), nil
		}
		return "", fmt.Errorf("Invalid byte size %d", argType.Size)
	case abi.BytesTy:
		return fmt.Sprintf("0x%s", hex.EncodeToString(val.([]byte))), nil
	case abi.HashTy:
		return val.(common.Hash).Hex(), nil
	case abi.FixedPointTy:
		return "", fmt.Errorf("Unhandled type %v", argType)
	case abi.FunctionTy:
		return "", fmt.Errorf("Unhandled type %v", argType)
	default:
		return "", fmt.Errorf("Unknown type %v", argType)
	}
}
