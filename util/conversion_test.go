// Copyright © 2017 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package util

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper to obtain a bigint from a string
func bigInt(val string) *big.Int {
	x := big.NewInt(0)
	x.SetString(val, 10)
	return x
}

func TestTokenValueToString(t *testing.T) {
	tests := []struct {
		input     *big.Int
		decimals  uint8
		usePrefix bool
		output    string
	}{
		{bigInt("0"), 0, false, "0"},
		{bigInt("1"), 0, false, "1"},
		{bigInt("1000"), 0, false, "1000"},
		{bigInt("1"), 2, false, "0.01"},
		{bigInt("10"), 2, false, "0.1"},
		{bigInt("100"), 2, false, "1"},
		{bigInt("1001"), 2, false, "10.01"},
		{bigInt("1010"), 2, false, "10.1"},
		{bigInt("9"), 18, false, "0.000000000000000009"},
		{bigInt("999999999"), 18, false, "0.000000000999999999"},
		{bigInt("99999999999999999"), 18, false, "0.099999999999999999"},
		{bigInt("999999999999999999"), 18, false, "0.999999999999999999"},
		{bigInt("9999999999999999999"), 18, false, "9.999999999999999999"},
		{bigInt("9000000000000000000"), 18, false, "9"},
		{bigInt("9100000000000000000"), 18, false, "9.1"},
		{bigInt("90000000000000000000"), 18, false, "90"},
		{bigInt("900000000000000000000"), 18, false, "900"},
		{bigInt("90000000000000000000000"), 18, false, "90000"},
		{bigInt("900000000000000000000000"), 18, false, "900000"},
		{bigInt("9000000000000000000000000"), 18, false, "9000000"},
		{bigInt("9000000000000000000000001"), 18, false, "9000000.000000000000000001"},
	}

	for _, tt := range tests {
		result := TokenValueToString(tt.input, tt.decimals, tt.usePrefix)
		if tt.output != result {
			t.Errorf("Failure: (%v, %v, %v) => %v (expected %v)\n", tt.input, tt.decimals, tt.usePrefix, result, tt.output)
		}
	}
}

func TestStringToTokenValue(t *testing.T) {
	tests := []struct {
		input    string
		decimals uint8
		output   *big.Int
	}{
		{"", 0, bigInt("0")},
		{"", 18, bigInt("0")},
		{"2", 0, bigInt("2")},
		{"3", 1, bigInt("30")},
		{"4", 18, bigInt("4000000000000000000")},
		{"5.000000000000000005", 18, bigInt("5000000000000000005")},
		{"6.000000000000000006", 18, bigInt("6000000000000000006")},
		{"777.777", 3, bigInt("777777")},
	}

	for _, tt := range tests {
		result, err := StringToTokenValue(tt.input, tt.decimals)
		assert.Nil(t, err, "Received error")
		assert.Equal(t, result, tt.output, "Did not receive expected result")
		if tt.output.Cmp(result) != 0 {
			t.Errorf("Failure: (\"%v\", %v) => %v (expected %v)\n", tt.input, tt.decimals, result, tt.output)
		}
	}
}
