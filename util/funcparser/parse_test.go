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
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/ethereal/v2/util"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		input  string
		output interface{}
	}{
		{
			name:  "NoParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test()`,
		},
		{
			name:   "SingleIntParameter",
			json:   `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"int256"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input:  `test(2)`,
			output: []interface{}{big.NewInt(2)},
		},
		{
			name:  "ArrayOfIntParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"int256[]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([1,2,3])`,
			output: []interface{}{[]*big.Int{
				big.NewInt(1), big.NewInt(2), big.NewInt(3),
			}},
		},
		{
			name:  "ArrayOfArrayOfIntParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"int256[][]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([[1,2,3],[4,5,6]])`,
			output: []interface{}{[][]*big.Int{
				{big.NewInt(1), big.NewInt(2), big.NewInt(3)},
				{big.NewInt(4), big.NewInt(5), big.NewInt(6)},
			}},
		},
		{
			name:   "StringParameter",
			json:   `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"string"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input:  `test("foo")`,
			output: []interface{}{`foo`},
		},
		{
			name:  "ArrayOfStringParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"string[]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test(["foo","bar","baz"])`,
			output: []interface{}{[]string{
				`foo`, `bar`, `baz`,
			}},
		},
		{
			name:  "ArrayOfArrayOfStringParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"string[][]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([["foo","bar","baz"],["qux","quux","quuz"]])`,
			output: []interface{}{[][]string{
				{`foo`, `bar`, `baz`},
				{`qux`, `quux`, `quuz`},
			}},
		},
		{
			name:   "BoolParameter",
			json:   `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bool"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input:  `test(true)`,
			output: []interface{}{true},
		},
		{
			name:  "ArrayOfBoolParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bool[]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([true, false])`,
			output: []interface{}{[]bool{
				true, false,
			}},
		},
		{
			name:  "ArrayOfArrayOfBoolParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bool[][]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([[true, false],[false, true]])`,
			output: []interface{}{[][]bool{
				{true, false},
				{false, true},
			}},
		},
		{
			name:   "AddressParameter",
			json:   `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"address"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input:  `test(0x008b7768c04a0c750C3D6b58d44Ff5041DD90480)`,
			output: []interface{}{common.HexToAddress("0x008b7768c04a0c750C3D6b58d44Ff5041DD90480")},
		},
		{
			name:  "ArrayOfAddressParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"address[]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([0x008b7768c04a0c750C3D6b58d44Ff5041DD90480,0x008B7768C04a0C750C3d6B58D44fF5041dd90481,0x008B7768c04A0c750C3D6b58d44fF5041dD90482])`,
			output: []interface{}{[]common.Address{
				common.HexToAddress("0x008b7768c04a0c750C3D6b58d44Ff5041DD90480"),
				common.HexToAddress("0x008B7768C04a0C750C3d6B58D44fF5041dd90481"),
				common.HexToAddress("0x008B7768c04A0c750C3D6b58d44fF5041dD90482"),
			}},
		},
		{
			name:  "ArrayOfArrayOfAddressParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"address[][]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([[0x008b7768c04a0c750C3D6b58d44Ff5041DD90480,0x008B7768C04a0C750C3d6B58D44fF5041dd90481,0x008B7768c04A0c750C3D6b58d44fF5041dD90482],[0x008B7768c04a0c750C3d6b58D44FF5041DD90483,0x008b7768c04a0C750c3D6b58d44FF5041dD90484,0x008b7768c04A0c750C3D6b58d44Ff5041dd90485]])`,
			output: []interface{}{[][]common.Address{
				{
					common.HexToAddress("0x008b7768c04a0c750C3D6b58d44Ff5041DD90480"),
					common.HexToAddress("0x008B7768C04a0C750C3d6B58D44fF5041dd90481"),
					common.HexToAddress("0x008B7768c04A0c750C3D6b58d44fF5041dD90482"),
				},
				{
					common.HexToAddress("0x008B7768c04a0c750C3d6b58D44FF5041DD90483"),
					common.HexToAddress("0x008b7768c04a0C750c3D6b58d44FF5041dD90484"),
					common.HexToAddress("0x008b7768c04A0c750C3D6b58d44Ff5041dd90485"),
				},
			}},
		},
		{
			name:   "BytesParameter",
			json:   `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bytes"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input:  `test(0x0102030405)`,
			output: []interface{}{_bytes("0102030405")},
		},
		{
			name:  "ArrayOfBytesParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bytes[]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([0x0102030405,0x060708090a])`,
			output: []interface{}{[][]byte{
				_bytes("0102030405"),
				_bytes("060708090a"),
			}},
		},
		{
			name:  "ArrayOfArrayOfBytesParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bytes[][]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([[0x0102030405,0x060708090a],[0x0b0c0d0e0f,0x1011121314]])`,
			output: []interface{}{[][][]byte{
				{
					_bytes("0102030405"),
					_bytes("060708090a"),
				},
				{
					_bytes("0b0c0d0e0f"),
					_bytes("1011121314"),
				},
			}},
		},
		{
			name:  "ArrayOfArrayOfFixedBytesParameters",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"constant":false,"inputs":[{"name":"arg1","type":"bytes32[][]"}],"name":"test","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]}}}`,
			input: `test([[0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f,0x202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f],[0x404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f,0x606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f]])`,
			output: []interface{}{[][][32]byte{
				{
					_bytes32("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
					_bytes32("202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f"),
				},
				{
					_bytes32("404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f"),
					_bytes32("606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f"),
				},
			}},
		},
		{
			name:  "Constructor",
			json:  `{"contracts":{"Test.sol:Test":{"abi":[{"inputs":[{"name":"arg1","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]}}}`,
			input: `constructor(12345)`,
		},
		// Need to implement tuple.
		// {
		// 	name:  "Tuple",
		// 	json:  `{"contracts":{"Test.sol:Test":{"abi":[{"inputs":[{"components":[{"internalType":"uint32","name":"field1","type":"uint32"},{"internalType":"uint64","name":"field2","type":"uint64"},{"internalType":"bool","name":"field3","type":"bool"}],"internalType":"struct Test.TestTuple","name":"arg1","type":"tuple"}],"name":"testTuple","outputs":[{"components":[{"internalType":"uint32","name":"field1","type":"uint32"},{"internalType":"uint64","name":"field2","type":"uint64"},{"internalType":"bool","name":"field3","type":"bool"}],"internalType":"struct Test.TestTuple","name":"","type":"tuple"}],"stateMutability":"pure","type":"function"}]}}}`,
		// 	input: `testTuple(12345,23456,(34567,45678,false))`,
		// },
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			contract, err := util.ParseCombinedJSON(test.json, "Test")
			require.Nil(t, err, fmt.Sprintf("failed to parse contract JSON at test %d", i))
			_, args, err := ParseCall(nil, contract, test.input)
			assert.Nil(t, err, fmt.Sprintf("failed to parse call at test %d", i))
			if test.output != nil {
				assert.Equal(t, test.output, args, fmt.Sprintf("incorrect value at test %d", i))
			}
		})
	}
}

func _bytes32(input string) [32]byte {
	var res [32]byte
	if len(strings.TrimPrefix(input, "0x")) != 64 {
		panic("incorrect length")
	}
	bytes, _ := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	copy(res[:], bytes)
	return res
}

func _bytes(input string) []byte {
	bytes, _ := hex.DecodeString(input)
	return bytes
}
