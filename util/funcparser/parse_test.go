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
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/ethereal/v2/util"
)

func TestParse(t *testing.T) {
	tests := []struct {
		json   string
		input  string
		output interface{}
	}{
		{ // 0 - no parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test()`,
		},
		{ // 1 - single int parameter
			json:   `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int256\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input:  `test(2)`,
			output: []interface{}{big.NewInt(2)},
		},
		{ // 2 - array of int parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int256[]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([1,2,3])`,
			output: []interface{}{[]*big.Int{
				big.NewInt(1), big.NewInt(2), big.NewInt(3),
			}},
		},
		{ // 3 - array of array of int parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int256[][]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([[1,2,3],[4,5,6]])`,
			output: []interface{}{[][]*big.Int{
				[]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)},
				[]*big.Int{big.NewInt(4), big.NewInt(5), big.NewInt(6)},
			}},
		},
		{ // 4 - string parameter
			json:   `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"string\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input:  `test("foo")`,
			output: []interface{}{`foo`},
		},
		{ // 5 - array of string parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"string[]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test(["foo","bar","baz"])`,
			output: []interface{}{[]string{
				`foo`, `bar`, `baz`,
			}},
		},
		{ // 6 - array of array of string parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"string[][]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([["foo","bar","baz"],["qux","quux","quuz"]])`,
			output: []interface{}{[][]string{
				[]string{`foo`, `bar`, `baz`},
				[]string{`qux`, `quux`, `quuz`},
			}},
		},
		{ // 7 - bool parameter
			json:   `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bool\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input:  `test(true)`,
			output: []interface{}{true},
		},
		{ // 8 - array of bool parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bool[]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([true, false])`,
			output: []interface{}{[]bool{
				true, false,
			}},
		},
		{ // 9 - array of array of bool parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bool[][]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([[true, false],[false, true]])`,
			output: []interface{}{[][]bool{
				[]bool{true, false},
				[]bool{false, true},
			}},
		},
		{ // 10 - address parameter
			json:   `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input:  `test(0x008b7768c04a0c750C3D6b58d44Ff5041DD90480)`,
			output: []interface{}{common.HexToAddress("0x008b7768c04a0c750C3D6b58d44Ff5041DD90480")},
		},
		{ // 11 - array of address parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"address[]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([0x008b7768c04a0c750C3D6b58d44Ff5041DD90480,0x008B7768C04a0C750C3d6B58D44fF5041dd90481,0x008B7768c04A0c750C3D6b58d44fF5041dD90482])`,
			output: []interface{}{[]common.Address{
				common.HexToAddress("0x008b7768c04a0c750C3D6b58d44Ff5041DD90480"),
				common.HexToAddress("0x008B7768C04a0C750C3d6B58D44fF5041dd90481"),
				common.HexToAddress("0x008B7768c04A0c750C3D6b58d44fF5041dD90482"),
			}},
		},
		{ // 12 - array of array of address parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"address[][]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([[0x008b7768c04a0c750C3D6b58d44Ff5041DD90480,0x008B7768C04a0C750C3d6B58D44fF5041dd90481,0x008B7768c04A0c750C3D6b58d44fF5041dD90482],[0x008B7768c04a0c750C3d6b58D44FF5041DD90483,0x008b7768c04a0C750c3D6b58d44FF5041dD90484,0x008b7768c04A0c750C3D6b58d44Ff5041dd90485]])`,
			output: []interface{}{[][]common.Address{
				[]common.Address{
					common.HexToAddress("0x008b7768c04a0c750C3D6b58d44Ff5041DD90480"),
					common.HexToAddress("0x008B7768C04a0C750C3d6B58D44fF5041dd90481"),
					common.HexToAddress("0x008B7768c04A0c750C3D6b58d44fF5041dD90482"),
				},
				[]common.Address{
					common.HexToAddress("0x008B7768c04a0c750C3d6b58D44FF5041DD90483"),
					common.HexToAddress("0x008b7768c04a0C750c3D6b58d44FF5041dD90484"),
					common.HexToAddress("0x008b7768c04A0c750C3D6b58d44Ff5041dd90485"),
				}}},
		},
		{ // 13 - bytes parameter
			json:   `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bytes\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input:  `test(0x0102030405)`,
			output: []interface{}{_bytes("0102030405")},
		},
		{ // 14 - array of bytes parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bytes[]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([0x0102030405,0x060708090a])`,
			output: []interface{}{[][]byte{
				_bytes("0102030405"),
				_bytes("060708090a"),
			}},
		},
		{ // 15 - array of array of bytes parameters
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"constant\":false,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bytes[][]\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"}}}`,
			input: `test([[0x0102030405,0x060708090a],[0x0b0c0d0e0f,0x1011121314]])`,
			output: []interface{}{[][][]byte{
				[][]byte{
					_bytes("0102030405"),
					_bytes("060708090a"),
				},
				[][]byte{
					_bytes("0b0c0d0e0f"),
					_bytes("1011121314"),
				}}},
		},
		{ // 16 - constructor
			json:  `{"contracts":{"Test.sol:Test":{"abi":"[{\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"}}}`,
			input: `constructor(12345)`,
		},
	}

	for i, test := range tests {
		contract, err := util.ParseCombinedJSON(test.json, "Test")
		require.Nil(t, err, fmt.Sprintf("failed to parse contract JSON at test %d", i))
		_, args, err := ParseCall(nil, contract, test.input)
		assert.Nil(t, err, fmt.Sprintf("failed to parse call at test %d", i))
		if test.output != nil {
			assert.Equal(t, test.output, args, fmt.Sprintf("incorrect value at test %d", i))
		}
	}
}

func _bytes(input string) []byte {
	bytes, _ := hex.DecodeString(input)
	return bytes
}
