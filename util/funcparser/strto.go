// Copyright © 2019 Weald Technology Trading
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

package funcparser

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// StrTo turns a string in to any simple type as given in the ABI information.
func StrTo(inputType *abi.Type, input string) (interface{}, error) {
	switch inputType.T {
	case abi.IntTy:
		return StrToInt(inputType, input)
	case abi.UintTy:
		return StrToUint(inputType, input)
	case abi.StringTy:
		return StrToStr(inputType, input)
	case abi.BoolTy:
		return StrToBool(inputType, input)
	case abi.AddressTy:
		return StrToAddress(inputType, input)
	case abi.HashTy:
		return StrToHash(inputType, input)
	case abi.BytesTy, abi.FixedBytesTy:
		return StrToBytes(inputType, input)
		//	case abi.ArrayTy, abi.SliceTy:
		//		baseType := baseType(inputType)
		//		level := arrayLevel(inputType)
		//		res, err := makeArray(baseType, level)
		//		if err != nil {
		//			return nil, err
		//		}
		//		for i := 0; i < arrayVal.Len(); i++ {
		//			elemRes, err := StrTo(inputType.Elem, fmt.Sprintf("%v", arrayVal.Index(i).Interface()))
		//			if err != nil {
		//				return nil, err
		//			}
		//			res = append(res, elemRes)
		//		}
		//		return res, nil
	default:
		return nil, fmt.Errorf("unhandled type %v", inputType)
	}
}

// StrToInt turns a string in to an int type as given by the ABI information.
// It can return various types so return interface{}.
func StrToInt(inputType *abi.Type, input string) (interface{}, error) {
	val := big.NewInt(0)
	val, success := val.SetString(input, 10)
	if !success {
		return nil, fmt.Errorf("invalid integer %s", input)
	}
	switch inputType.Size {
	case 8:
		return int8(val.Int64()), nil
	case 16:
		return int16(val.Int64()), nil
	case 32:
		return int32(val.Int64()), nil
	case 64:
		return val.Int64(), nil
	case 256:
		return val, nil
	default:
		return nil, fmt.Errorf("unexpected int size %d", inputType.Size)
	}
}

// _zero is for checking negative values against unsigned definitions.
var _zero = big.NewInt(0)

// StrToUint turns a string in to a uint type as given by the ABI information.
// It can return various types so return interface{}.
func StrToUint(inputType *abi.Type, input string) (interface{}, error) {
	val := big.NewInt(0)
	val, success := val.SetString(input, 10)
	if !success {
		return nil, fmt.Errorf("invalid unsigned integer %s", input)
	}
	if val.Cmp(_zero) < 0 {
		return nil, fmt.Errorf("invalid unsigned integer %s", input)
	}
	switch inputType.Size {
	case 8:
		return uint8(val.Uint64()), nil
	case 16:
		return uint16(val.Uint64()), nil
	case 32:
		return uint32(val.Uint64()), nil
	case 64:
		return val.Uint64(), nil
	case 256:
		return val, nil
	default:
		return nil, fmt.Errorf("unexpected int size %d", inputType.Size)
	}
}

// StrToStr turns a string in to a string type as given by the ABI information.
func StrToStr(_ *abi.Type, input string) (string, error) {
	rep := strings.NewReplacer(`\"`, "")
	input = rep.Replace(input)
	return input[1 : len(input)-1], nil
}

// StrToBool turns a string in to a boolean type as given by the ABI information.
func StrToBool(_ *abi.Type, input string) (bool, error) {
	return input == "true", nil
}

// StrToAddress turns a string in to an address type as given by the ABI information.
func StrToAddress(_ *abi.Type, input string) (common.Address, error) {
	return common.HexToAddress(strings.TrimPrefix(input, "0x")), nil
}

// StrToHash turns a string in to a hash type as given by the ABI information.
func StrToHash(_ *abi.Type, input string) (common.Hash, error) {
	return common.HexToHash(strings.TrimPrefix(input, "0x")), nil
}

// StrToBytes turns a string in to a bytes type as given by the ABI information.
// It can return various types so return interface{}.
// nolint:gocyclo
func StrToBytes(inputType *abi.Type, input string) (interface{}, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid byte string %s", input)
	}
	if inputType.Size == 0 {
		return decoded, nil
	}
	slice := make([]byte, inputType.Size)
	copy(slice[inputType.Size-len(decoded):inputType.Size], decoded)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	switch inputType.Size {
	case 0:
		return slice, nil
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
	case 31:
		return *(*[31]uint8)(unsafe.Pointer(hdr.Data)), nil
	case 32:
		return *(*[32]uint8)(unsafe.Pointer(hdr.Data)), nil
	}
	return nil, fmt.Errorf("invalid byte size %d", inputType.Size)
}
