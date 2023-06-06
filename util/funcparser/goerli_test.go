// Copyright © 2019, 2022 Weald Technology Trading
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
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/ethereal/v2/util"
)

func TestGoerli(t *testing.T) {
	tests := []struct {
		json   string
		input  string
		output interface{}
	}{
		{ // 0 - no parameters
			input: `test()`,
		},
		{ // 1 - single int parameter
			input:  `testInt256(65536)`,
			output: []interface{}{big.NewInt(65536)},
		},
		{ // 2 - array of int parameters
			input: `testInt256Array([1,2,3])`,
			output: []interface{}{[]interface{}{
				big.NewInt(1), big.NewInt(2), big.NewInt(3),
			}},
		},
		{ // 3 - array of array of int parameters
			input: `testInt2562DArray([[1,2,3],[4,5,6]])`,
			output: []interface{}{[]interface{}{
				[]interface{}{big.NewInt(1), big.NewInt(2), big.NewInt(3)},
				[]interface{}{big.NewInt(4), big.NewInt(5), big.NewInt(6)},
			}},
		},
		{ // 4 - string parameter
			input:  `testString("foo")`,
			output: []interface{}{`foo`},
		},
		{ // 5 - array of string parameters
			input: `testStringArray(["foo","bar","baz"])`,
			output: []interface{}{[]string{
				`foo`, `bar`, `baz`,
			}},
		},
		{ // 6 - array of array of string parameters
			input: `testString2DArray([["foo","bar","baz"],["qux","quux","quuz"]])`,
			output: []interface{}{[][]string{
				{`foo`, `bar`, `baz`},
				{`qux`, `quux`, `quuz`},
			}},
		},
		{ // 7 - bool parameter
			input:  `testBool(true)`,
			output: []interface{}{true},
		},
		{ // 8 - array of bool parameters
			input: `testBoolArray([true, false])`,
			output: []interface{}{
				[]bool{true, false},
			},
		},
		{ // 9 - array of array of bool parameters
			input: `testBool2DArray([[true, false],[false, true]])`,
			output: []interface{}{[][]bool{
				{true, false},
				{false, true},
			}},
		},
		{ // 10 - address parameter
			input:  `testAddress(0x108b7768c04a0c750C3D6b58d44Ff5041DD90480)`,
			output: []interface{}{common.HexToAddress("0x108b7768c04a0c750C3D6b58d44Ff5041DD90480")},
		},
		{ // 11 - array of address parameters
			input: `testAddressArray([0x108b7768c04a0c750C3D6b58d44Ff5041DD90480,0x108B7768C04a0C750C3d6B58D44fF5041dd90481,0x108B7768c04A0c750C3D6b58d44fF5041dD90482])`,
			output: []interface{}{[]common.Address{
				common.HexToAddress("0x108b7768c04a0c750C3D6b58d44Ff5041DD90480"),
				common.HexToAddress("0x108B7768C04a0C750C3d6B58D44fF5041dd90481"),
				common.HexToAddress("0x108B7768c04A0c750C3D6b58d44fF5041dD90482"),
			}},
		},
		{ // 12 - array of array of address parameters
			input: `testAddress2DArray([[0x108b7768c04a0c750C3D6b58d44Ff5041DD90480,0x108B7768C04a0C750C3d6B58D44fF5041dd90481,0x108B7768c04A0c750C3D6b58d44fF5041dD90482],[0x108B7768c04a0c750C3d6b58D44FF5041DD90483,0x108b7768c04a0C750c3D6b58d44FF5041dD90484,0x108b7768c04A0c750C3D6b58d44Ff5041dd90485]])`,
			output: []interface{}{[][]common.Address{
				{
					common.HexToAddress("0x108b7768c04a0c750C3D6b58d44Ff5041DD90480"),
					common.HexToAddress("0x108B7768C04a0C750C3d6B58D44fF5041dd90481"),
					common.HexToAddress("0x108B7768c04A0c750C3D6b58d44fF5041dD90482"),
				},
				{
					common.HexToAddress("0x108B7768c04a0c750C3d6b58D44FF5041DD90483"),
					common.HexToAddress("0x108b7768c04a0C750c3D6b58d44FF5041dD90484"),
					common.HexToAddress("0x108b7768c04A0c750C3D6b58d44Ff5041dd90485"),
				},
			}},
		},
		{ // 13 - bytes parameter
			input:  `testBytes(0x0102030405)`,
			output: []interface{}{_bytes("0102030405")},
		},
		{ // 14 - array of bytes parameters
			input: `testBytesArray([0x0102030405,0x060708090a])`,
			output: []interface{}{[][]byte{
				_bytes("0102030405"),
				_bytes("060708090a"),
			}},
		},
		{ // 15 - array of array of bytes parameters
			input: `testBytes2DArray([[0x0102030405,0x060708090a],[0x0b0c0d0e0f,0x1011121314]])`,
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
	}

	// Connect to goerli
	conn, err := ethclient.Dial("https://goerli.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.Nil(t, err, "failed to connect to goerli")

	json, err := os.ReadFile("Tester.json")
	require.Nil(t, err, "failed to read Tester ABI")
	contract, err := util.ParseCombinedJSON(string(json), "Tester")
	require.Nil(t, err, "failed to parse contract JSON")

	fromAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	contractAddress := common.HexToAddress("0xf942De0Be16b8B30C07C4D0d3Fca725277306892")
	for i, test := range tests {
		method, args, err := ParseCall(conn, contract, test.input)
		require.Nil(t, err, fmt.Sprintf("failed to parse call at test %d", i))
		data, err := contract.Abi.Pack(method.Name, args...)
		assert.Nil(t, err, fmt.Sprintf("failed to pack data at test %d", i))

		// Make the call
		msg := ethereum.CallMsg{
			From: fromAddress,
			To:   &contractAddress,
			Data: data,
		}
		ctx := context.Background()
		result, err := conn.CallContract(ctx, msg, nil)
		assert.Nil(t, err, fmt.Sprintf("failed to call contract at test %d", i))
		assert.NotNil(t, result)
	}
}
