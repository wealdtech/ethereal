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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/ethereal/v2/util"
	"github.com/wealdtech/ethereal/v2/util/funcparser/parser"
	ens "github.com/wealdtech/go-ens/v3"
)

type methodListener struct {
	*parser.BaseFuncListener
	client   *ethclient.Client
	contract *util.Contract
	curArg   int
	// Arrays are all of the same type but can be nested.
	curArray      []any
	maxArrayLevel int
	// Result of parsing the argument.
	method *abi.Method
	args   []any
	err    error
}

// newMethodListener creates a new method listener.
func newMethodListener(client *ethclient.Client, contract *util.Contract) *methodListener {
	return &methodListener{
		client:   client,
		contract: contract,
		curArg:   0,
		args:     make([]any, 0),
	}
}

func (l *methodListener) EnterFuncName(c *parser.FuncNameContext) {
	// Ensure we have the function in the contract.
	if c.GetText() == "constructor" {
		l.method = &l.contract.Abi.Constructor
	} else {
		method, exists := l.contract.Abi.Methods[c.GetText()]
		if exists {
			l.method = &method
		} else {
			l.err = fmt.Errorf("unknown method name %s", c.GetText())
		}
	}
}

func (l *methodListener) EnterIntArg(c *parser.IntArgContext) {
	if l.err == nil {
		input := l.method.Inputs[l.curArg]
		var err error
		var arg any
		baseType := baseType(&input.Type)
		switch baseType.T {
		case abi.IntTy:
			arg, err = StrToInt(baseType, c.GetText())
		case abi.UintTy:
			arg, err = StrToUint(baseType, c.GetText())
		case abi.AddressTy:
			err = fmt.Errorf("address \"%s\" looks like number; prefix it with \"0x\"", c.GetText())
		default:
			err = fmt.Errorf("unexpected type %v", baseType)
		}
		if err != nil {
			l.err = err
		} else {
			l.pushArg(arg)
		}
	}
}

func (l *methodListener) EnterBoolArg(c *parser.BoolArgContext) {
	if l.err == nil {
		input := l.method.Inputs[l.curArg]
		baseType := baseType(&input.Type)
		arg, err := StrToBool(baseType, c.GetText())
		if err != nil {
			l.err = err
		} else {
			l.pushArg(arg)
		}
	}
}

func (l *methodListener) EnterStringArg(c *parser.StringArgContext) {
	if l.err == nil {
		input := l.method.Inputs[l.curArg]
		baseType := baseType(&input.Type)
		arg, err := StrToStr(baseType, c.GetText())
		if err != nil {
			l.err = err
		} else {
			l.pushArg(arg)
		}
	}
}

func (l *methodListener) EnterArrayArg(c *parser.ArrayArgContext) {
	if l.err == nil {
		if len(l.method.Inputs) <= l.curArg {
			l.err = fmt.Errorf("too many arguments for method at %s", c.GetText())
			return
		}
		input := l.method.Inputs[l.curArg]
		baseType := baseType(&input.Type)
		level := arrayLevel(&input.Type)
		if len(l.curArray) == 0 {
			// New array.
			l.curArray = make([]any, 0)
			l.maxArrayLevel = level
		} else {
			// Extend existing array.
			level -= len(l.curArray)
		}
		array, err := makeArray(baseType, level)
		if err != nil {
			l.err = err
			return
		}
		l.curArray = append(l.curArray, array)
	}
}

// nolint:gocyclo
func (l *methodListener) ExitArrayArg(_ *parser.ArrayArgContext) {
	if l.err == nil {
		level := len(l.curArray)
		if level == 1 {
			// Only array; push to args.
			l.args = append(l.args, l.curArray[0])
		} else {
			// Nested arrays; push to one above.
			input := l.method.Inputs[l.curArg]
			baseType := baseType(&input.Type)
			level = l.maxArrayLevel + 1 - level
			parent := len(l.curArray) - 2
			child := len(l.curArray) - 1
			if level == 1 {
				switch baseType.T {
				case abi.IntTy:
					switch baseType.Size {
					case 8:
						l.curArray[parent] = append(l.curArray[parent].([][]int8), l.curArray[child].([]int8))
					case 16:
						l.curArray[parent] = append(l.curArray[parent].([][]int16), l.curArray[child].([]int16))
					case 32:
						l.curArray[parent] = append(l.curArray[parent].([][]int32), l.curArray[child].([]int32))
					case 64:
						l.curArray[parent] = append(l.curArray[parent].([][]int64), l.curArray[child].([]int64))
					default:
						l.curArray[parent] = append(l.curArray[parent].([][]*big.Int), l.curArray[child].([]*big.Int))
					}
				case abi.UintTy:
					switch baseType.Size {
					case 8:
						l.curArray[parent] = append(l.curArray[parent].([][]uint8), l.curArray[child].([]uint8))
					case 16:
						l.curArray[parent] = append(l.curArray[parent].([][]uint16), l.curArray[child].([]uint16))
					case 32:
						l.curArray[parent] = append(l.curArray[parent].([][]uint32), l.curArray[child].([]uint32))
					case 64:
						l.curArray[parent] = append(l.curArray[parent].([][]uint64), l.curArray[child].([]uint64))
					default:
						l.curArray[parent] = append(l.curArray[parent].([][]*big.Int), l.curArray[child].([]*big.Int))
					}
				case abi.BoolTy:
					l.curArray[parent] = append(l.curArray[parent].([][]bool), l.curArray[child].([]bool))
				case abi.StringTy:
					l.curArray[parent] = append(l.curArray[parent].([][]string), l.curArray[child].([]string))
				case abi.AddressTy:
					l.curArray[parent] = append(l.curArray[parent].([][]common.Address), l.curArray[child].([]common.Address))
				case abi.HashTy:
					l.curArray[parent] = append(l.curArray[parent].([][]common.Hash), l.curArray[child].([]common.Hash))
				case abi.BytesTy, abi.FixedBytesTy:
					switch baseType.Size {
					case 0:
						l.curArray[parent] = append(l.curArray[parent].([][][]byte), l.curArray[child].([][]byte))
					case 32:
						l.curArray[parent] = append(l.curArray[parent].([][][32]byte), l.curArray[child].([][32]byte))
					default:
						panic("unhandled size")
					}
				default:
					l.curArray[parent] = append(l.curArray[parent].([]any), l.curArray[child])
				}
			} else if level == 2 {
				switch baseType.T {
				case abi.IntTy:
					switch baseType.Size {
					case 8:
						l.curArray[parent] = append(l.curArray[parent].([][][]int8), l.curArray[child].([][]int8))
					case 16:
						l.curArray[parent] = append(l.curArray[parent].([][][]int16), l.curArray[child].([][]int16))
					case 32:
						l.curArray[parent] = append(l.curArray[parent].([][][]int32), l.curArray[child].([][]int32))
					case 64:
						l.curArray[parent] = append(l.curArray[parent].([][][]int64), l.curArray[child].([][]int64))
					default:
						l.curArray[parent] = append(l.curArray[parent].([][][]*big.Int), l.curArray[child].([][]*big.Int))
					}
				case abi.UintTy:
					switch baseType.Size {
					case 8:
						l.curArray[parent] = append(l.curArray[parent].([][][]uint8), l.curArray[child].([][]uint8))
					case 16:
						l.curArray[parent] = append(l.curArray[parent].([][][]uint16), l.curArray[child].([][]uint16))
					case 32:
						l.curArray[parent] = append(l.curArray[parent].([][][]uint32), l.curArray[child].([][]uint32))
					case 64:
						l.curArray[parent] = append(l.curArray[parent].([][][]uint64), l.curArray[child].([][]uint64))
					default:
						l.curArray[parent] = append(l.curArray[parent].([][][]*big.Int), l.curArray[child].([][]*big.Int))
					}
				case abi.BoolTy:
					l.curArray[parent] = append(l.curArray[parent].([][][]bool), l.curArray[child].([][]bool))
				case abi.StringTy:
					l.curArray[parent] = append(l.curArray[parent].([][][]string), l.curArray[child].([][]string))
				case abi.AddressTy:
					l.curArray[parent] = append(l.curArray[parent].([][][]common.Address), l.curArray[child].([][]common.Address))
				case abi.HashTy:
					l.curArray[parent] = append(l.curArray[parent].([][][]common.Hash), l.curArray[child].([][]common.Hash))
				case abi.BytesTy, abi.FixedBytesTy:
					l.curArray[parent] = append(l.curArray[parent].([][][][]byte), l.curArray[child].([][][]byte))
				default:
					l.curArray[parent] = append(l.curArray[parent].([][]any), l.curArray[child].([]any))
				}
			}
		}
		l.curArray = l.curArray[:len(l.curArray)-1]
	}
}

func (l *methodListener) EnterDomainArg(c *parser.DomainArgContext) {
	if l.err == nil {
		input := l.method.Inputs[l.curArg]
		var err error
		var arg any
		baseType := baseType(&input.Type)
		switch baseType.T {
		case abi.AddressTy:
			arg, err = ens.Resolve(l.client, c.GetText()[1:])
		default:
			err = fmt.Errorf("unexpected type %v", baseType)
		}
		if err != nil {
			l.err = err
		} else {
			l.pushArg(arg)
		}
	}
}

func (l *methodListener) EnterHexArg(c *parser.HexArgContext) {
	if l.err == nil {
		input := l.method.Inputs[l.curArg]
		var err error
		var arg any
		baseType := baseType(&input.Type)
		switch baseType.T {
		case abi.AddressTy:
			arg, err = StrToAddress(baseType, c.GetText())
		case abi.HashTy:
			arg, err = StrToHash(baseType, c.GetText())
		case abi.BytesTy, abi.FixedBytesTy:
			arg, err = StrToBytes(baseType, c.GetText())
		default:
			err = fmt.Errorf("unexpected type %v", baseType)
		}
		if err != nil {
			l.err = err
		} else {
			l.pushArg(arg)
		}
	}
}

func (l *methodListener) EnterArg(_ *parser.ArgContext) {
	if l.err == nil {
		if l.curArg >= len(l.method.Inputs) {
			l.err = fmt.Errorf("too many arguments (expected %d)", len(l.method.Inputs))
		}
	}
}

func (l *methodListener) ExitArg(_ *parser.ArgContext) {
	if l.err == nil {
		// We only increment the argument if we aren't in an array.
		if len(l.curArray) == 0 {
			l.curArg++
		}
	}
}

func baseType(inputType *abi.Type) *abi.Type {
	switch inputType.T {
	case abi.SliceTy:
		return baseType(inputType.Elem)
	case abi.ArrayTy:
		return baseType(inputType.Elem)
	default:
		return inputType
	}
}

// arrayLevel returns the number of levels of array in the type.
func arrayLevel(inputType *abi.Type) int {
	switch inputType.T {
	case abi.SliceTy, abi.ArrayTy:
		return 1 + arrayLevel(inputType.Elem)
	default:
		return 0
	}
}

func (l *methodListener) pushArg(arg any) {
	if len(l.curArray) == 0 {
		l.args = append(l.args, arg)
	} else {
		input := l.method.Inputs[l.curArg]
		baseType := baseType(&input.Type)
		switch baseType.T {
		case abi.IntTy:
			switch baseType.Size {
			case 8:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]int8), arg.(int8))
			case 16:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]int16), arg.(int16))
			case 32:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]int32), arg.(int32))
			case 64:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]int64), arg.(int64))
			default:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]*big.Int), arg.(*big.Int))
			}
		case abi.UintTy:
			switch baseType.Size {
			case 8:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]uint8), arg.(uint8))
			case 16:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]uint16), arg.(uint16))
			case 32:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]uint32), arg.(uint32))
			case 64:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]uint64), arg.(uint64))
			default:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]*big.Int), arg.(*big.Int))
			}
		case abi.BoolTy:
			l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]bool), arg.(bool))
		case abi.StringTy:
			l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]string), arg.(string))
		case abi.AddressTy:
			l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]common.Address), arg.(common.Address))
		case abi.HashTy:
			l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]common.Hash), arg.(common.Hash))
		case abi.BytesTy, abi.FixedBytesTy:
			switch baseType.Size {
			case 0:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([][]byte), arg.([]byte))
			case 32:
				l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([][32]byte), arg.([32]byte))
			default:
				panic("not handling size")
			}
		case abi.SliceTy:
			panic("not handling slice")
		default:
			l.curArray[len(l.curArray)-1] = append(l.curArray[len(l.curArray)-1].([]any), arg)
		}
	}
}

func makeArray(baseType *abi.Type, level int) (any, error) {
	switch level {
	case 1:
		return make1DArray(baseType)
	case 2:
		return make2DArray(baseType)
	default:
		return nil, fmt.Errorf("unhandled nesting level %d", level)
	}
}

func make2DArray(baseType *abi.Type) (any, error) {
	switch baseType.T {
	case abi.IntTy:
		switch baseType.Size {
		case 8:
			return make([][]int8, 0), nil
		case 16:
			return make([][]int16, 0), nil
		case 32:
			return make([][]int32, 0), nil
		case 64:
			return make([][]int64, 0), nil
		default:
			return make([][]*big.Int, 0), nil
		}
	case abi.UintTy:
		switch baseType.Size {
		case 8:
			return make([][]uint8, 0), nil
		case 16:
			return make([][]uint16, 0), nil
		case 32:
			return make([][]uint32, 0), nil
		case 64:
			return make([][]uint64, 0), nil
		default:
			return make([][]*big.Int, 0), nil
		}
	case abi.BoolTy:
		return make([][]bool, 0), nil
	case abi.StringTy:
		return make([][]string, 0), nil
	case abi.AddressTy:
		return make([][]common.Address, 0), nil
	case abi.HashTy:
		return make([][]common.Hash, 0), nil
	case abi.BytesTy, abi.FixedBytesTy:
		switch baseType.Size {
		case 0:
			return make([][][]byte, 0), nil
		case 32:
			return make([][][32]byte, 0), nil
		default:
			panic("unhandled size")
		}
	default:
		return nil, fmt.Errorf("unhandled array type %v", baseType.T)
	}
}

func make1DArray(baseType *abi.Type) (any, error) {
	switch baseType.T {
	case abi.IntTy:
		switch baseType.Size {
		case 8:
			return make([]int8, 0), nil
		case 16:
			return make([]int16, 0), nil
		case 32:
			return make([]int32, 0), nil
		case 64:
			return make([]int64, 0), nil
		default:
			return make([]*big.Int, 0), nil
		}
	case abi.UintTy:
		switch baseType.Size {
		case 8:
			return make([]uint8, 0), nil
		case 16:
			return make([]uint16, 0), nil
		case 32:
			return make([]uint32, 0), nil
		case 64:
			return make([]uint64, 0), nil
		default:
			return make([]*big.Int, 0), nil
		}
	case abi.BoolTy:
		return make([]bool, 0), nil
	case abi.StringTy:
		return make([]string, 0), nil
	case abi.AddressTy:
		return make([]common.Address, 0), nil
	case abi.HashTy:
		return make([]common.Hash, 0), nil
	case abi.BytesTy, abi.FixedBytesTy:
		switch baseType.Size {
		case 0:
			return make([][]byte, 0), nil
		case 32:
			return make([][32]byte, 0), nil
		default:
			panic("unhandled size")
		}
	default:
		return nil, fmt.Errorf("unhandled array type %v", baseType.T)
	}
}
