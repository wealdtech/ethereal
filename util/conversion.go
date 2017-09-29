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
package util

import (
	"math/big"
	"regexp"
	"strings"
)

// Used in TokenValueToString
var zero = big.NewInt(0)
var ten = big.NewInt(10)
var thousand = big.NewInt(1000)
var million = big.NewInt(1000000)

func TokenValueToString(input *big.Int, decimals uint8, usePrefix bool) (output string) {
	// Take a string version of the input
	value := input.String()

	// Input sanity checks
	if input.Cmp(zero) == 0 {
		return "0"
	}

	nonDecimalLength := len(value) - int(decimals)
	if nonDecimalLength <= 0 {
		// We need to add leading decimal 0s
		output = "0." + strings.Repeat("0", int(decimals)-len(value)) + value
		output = strings.TrimRight(output, "0")
	} else if nonDecimalLength == len(value) {
		output = value
	} else {
		// We might need to add decimal point
		postDecimalsAllZero, _ := regexp.MatchString("^0+$", value[nonDecimalLength:])
		if postDecimalsAllZero {
			// Nope; just take the leading figures
			output = value[:nonDecimalLength]
		} else {
			// Yep; add the zeros
			output = value[:nonDecimalLength] + "." + value[nonDecimalLength:]
			output = strings.TrimRight(output, "0")
		}
	}
	return
}
