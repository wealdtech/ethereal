// Copyright © 2018 Weald Technology Trading
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

package txdata

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
	"golang.org/x/crypto/sha3"
)

var (
	functions map[[4]byte]function
	events    map[[32]byte]function
)

type function struct {
	name   string
	params []string
}

var blacklist = map[string]bool{
	"AcoraidaMonicaWantsToKeepALogOfTheWinner(address)":        true,
	"available_assert_time(uint16,uint64)":                     true,
	"clash550254402()":                                         true,
	"collate_propagate_storage(bytes16)":                       true,
	"coral_cable_news(uint256)":                                true,
	"gasprice_bit_ether(int128)":                               true,
	"getACLRole5999294130779334338()":                          true,
	"ideal_warn_timed(uint256,uint128)":                        true,
	"link_classic_internal(uint64,uint64)":                     true,
	"many_msg_babbage(bytes1)":                                 true,
	"message_hour(uint256,int8,uint16,bytes32)":                true,
	"passphrase_calculate_transfer(uint64,address)":            true,
	"pizza_mandate_apology(uint256)":                           true,
	"remove_good(uint256[],bytes8,bool)":                       true,
	"setTarget(address)":                                       true,
	"sign_szabo_bytecode(bytes16,uint128)":                     true,
	"transfer(bytes4[9],bytes5[6],int48[11])":                  true,
	"voting_var(address,uint256,int128,int128)":                true,
	"hello(int16[15],bytes22[5],int176[8])":                    true,
	"adminResetRank()":                                         true,
	"blockHashAmphithyronVersify(uint256)":                     true,
	"blockHashAmarilloNonspontaneously(uint256)":               true,
	"blockHashAddendsInexpansible(uint256)":                    true,
	"left_branch_block(uint32)":                                true,
	"overdiffusingness(bytes,uint256,uint256,uint256,uint256)": true,
	"get_block_hash_257335279069929(uint256)":                  true,
	"withdrawByAdmin_Unau(uint256[])":                          true,
	"bright_peace(bytes32,bytes)":                              true,
	"initialFundsReleaseNumerator()":                           true,
	"link_classic_internal(uint64,int64)":                      true,
	"branch_passphrase_public(uint256,bytes8)":                 true,
}

// String implements the stringer interface.
func (f *function) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s(", f.name))
	for i, param := range f.params {
		buffer.WriteString(param)
		if i < len(f.params)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

// DataToString takes a transaction's data bytes and converts it in to a useful representation if one exists.
func DataToString(client *ethclient.Client, input []byte) string {
	if len(input) == 0 {
		return ""
	}
	if len(input) < 4 {
		return fmt.Sprintf("0x%x", input)
	}
	var sig [4]byte
	copy(sig[:], input[:4])
	function, exists := functions[sig]
	if !exists {
		return fmt.Sprintf("0x%x", input)
	}
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s(", function.name))
	for i, param := range function.params {
		t, err := abi.NewType(param, "", nil)
		if err == nil {
			res, err := contractValueToString(client, t, uint32(i), input)
			if err != nil {
				res = err.Error()
			}
			buffer.WriteString(res)
			if i < len(function.params)-1 {
				buffer.WriteString(",")
			}
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

// EventToString takes a transaction's event information and converts it to a useful representation if one exists.
func EventToString(client *ethclient.Client, input *types.Log) string {
	function, exists := events[input.Topics[0]]
	if !exists {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s(", function.name))

	// Turn topics in to a byte array.
	topics := make([]byte, 32*len(input.Topics))
	for i := range input.Topics {
		copy(topics[i*32:i*32+32], input.Topics[i].Bytes())
	}

	curTopic := 1
	for i, param := range function.params {
		t, err := abi.NewType(param, "", nil)
		if err == nil {
			var res string
			var err error
			if len(input.Topics) > curTopic {
				res, err = valueToString(client, t, uint32(curTopic), 0, topics)
				curTopic++
			} else {
				res, err = valueToString(client, t, uint32(i+1-curTopic), 0, input.Data)
			}
			if err != nil {
				res = err.Error()
			}
			buffer.WriteString(res)
			if i < len(function.params)-1 {
				buffer.WriteString(",")
			}
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func contractValueToString(client *ethclient.Client, argType abi.Type, index uint32, data []byte) (string, error) {
	return valueToString(client, argType, index, 4, data)
}

func valueToString(client *ethclient.Client, argType abi.Type, index uint32, offset uint32, data []byte) (string, error) {
	switch argType.T {
	case abi.IntTy:
		return big.NewInt(0).SetBytes(data[offset+index*32 : offset+index*32+32]).String(), nil
	case abi.UintTy:
		return big.NewInt(0).SetBytes(data[offset+index*32 : offset+index*32+32]).String(), nil
	case abi.BoolTy:
		if data[offset+index*32+31] == 0x01 {
			return "true", nil
		}
		return "false", nil
	case abi.StringTy:
		start := binary.BigEndian.Uint32(data[offset+index*32+28 : offset+index*32+32])
		length := binary.BigEndian.Uint32(data[offset+start+28 : offset+start+32])
		return fmt.Sprintf("\"%s\"", string(data[offset+start+32:offset+start+32+length])), nil
	case abi.SliceTy, abi.ArrayTy:
		res := make([]string, 0)
		start := binary.BigEndian.Uint32(data[offset+index*32+28 : offset+index*32+32])
		entries := binary.BigEndian.Uint32(data[offset+start+28 : offset+start+32])
		for i := uint32(0); i < entries; i++ {
			elemRes, err := valueToString(client, *argType.Elem, 1+start/32+i, offset, data)
			if err != nil {
				return "", err
			}
			res = append(res, elemRes)
		}
		return "[" + strings.Join(res, ",") + "]", nil
	case abi.AddressTy:
		address := common.BytesToAddress(data[offset+index*32+12 : offset+index*32+32])
		return ens.Format(client, address), nil
	case abi.FixedBytesTy:
		return fmt.Sprintf("0x%x", data[offset+index*32+32-uint32(argType.Size):offset+index*32+32]), nil
	case abi.BytesTy:
		start := binary.BigEndian.Uint32(data[offset+index*32+28 : offset+index*32+32])
		length := binary.BigEndian.Uint32(data[offset+start+28 : offset+start+32])
		return fmt.Sprintf("0x%x", data[offset+start+32:offset+start+32+length]), nil
	case abi.HashTy:
		return fmt.Sprintf("0x%x", data[offset+index*32:offset+index*32+32]), nil
	case abi.FixedPointTy:
		return "", fmt.Errorf("unhandled type %v", argType)
	case abi.FunctionTy:
		return "", fmt.Errorf("unhandled type %v", argType)
	default:
		return "", fmt.Errorf("unknown type %v", argType)
	}
}

// AddFunctionSignature adds a function signature to the translation list.
func AddFunctionSignature(signature string) {
	// Start off removing parameter names if present.
	sigBits := strings.Split(strings.TrimSuffix(signature, ")"), "(")
	name := sigBits[0]
	params := strings.Split(sigBits[1], ",")
	if params[0] == "" {
		params = make([]string, 0)
	}
	for i := range params {
		params[i] = strings.TrimSpace(params[i])
		params[i] = strings.Split(params[i], " ")[0]
	}
	signature = fmt.Sprintf("%s(%s)", name, strings.Join(params, ","))

	// Do not add if on the blacklist.
	if _, exists := blacklist[signature]; exists {
		return
	}

	var hash [32]byte
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write([]byte(signature))
	if err != nil {
		return
	}
	sha.Sum(hash[:0])
	var sig [4]byte
	copy(sig[:], hash[:4])

	functions[sig] = function{name: name, params: params}
	// Also add to events.
	events[hash] = function{name: name, params: params}
}

// AddEventSignature adds an event signature to the translation list.
func AddEventSignature(signature string) {
	// Start off removing parameter names if present.
	sigBits := strings.Split(strings.TrimSuffix(signature, ")"), "(")
	name := sigBits[0]
	params := strings.Split(sigBits[1], ",")
	if params[0] == "" {
		params = make([]string, 0)
	}
	for i := range params {
		params[i] = strings.TrimSpace(params[i])
		params[i] = strings.Split(params[i], " ")[0]
	}
	signature = fmt.Sprintf("%s(%s)", name, strings.Join(params, ","))

	var hash [32]byte
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write([]byte(signature))
	if err != nil {
		return
	}
	sha.Sum(hash[:0])
	var sig [4]byte
	copy(sig[:], hash[:4])

	events[hash] = function{name: name, params: params}
}
