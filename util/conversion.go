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

// Package util contains utilities for Ethereal.
package util

import (
	"math/big"
	"regexp"
	"strings"
)

// Used in TokenValueToString.
var zero = big.NewInt(0)

// TokenValueToString converts a token value to a suitable string representation.
func TokenValueToString(input *big.Int, decimals uint8, _ bool) string {
	output := ""

	// Take a string version of the input.
	value := input.String()

	// Input sanity checks.
	if input.Cmp(zero) == 0 {
		return "0"
	}

	nonDecimalLength := len(value) - int(decimals)
	switch {
	case nonDecimalLength <= 0:
		// We need to add leading decimal 0s.
		output = "0." + strings.Repeat("0", int(decimals)-len(value)) + value
		output = strings.TrimRight(output, "0")
	case nonDecimalLength == len(value):
		output = value
	default:
		// We might need to add decimal point.
		postDecimalsAllZero, _ := regexp.MatchString("^0+$", value[nonDecimalLength:])
		if postDecimalsAllZero {
			// Nope; just take the leading figures.
			output = value[:nonDecimalLength]
		} else {
			// Yep; add the zeros.
			output = value[:nonDecimalLength] + "." + value[nonDecimalLength:]
			output = strings.TrimRight(output, "0")
		}
	}
	return output
}

// StringToTokenValue converts a string to a number of tokens.
func StringToTokenValue(input string, decimals uint8) (*big.Int, error) {
	output := big.NewInt(0)
	if input == "" {
		return output, nil
	}

	// Count the number of items after the decimal point.
	parts := strings.Split(input, ".")
	var additionalZeros int
	if len(parts) == 2 {
		// There is a decimal place.
		additionalZeros = int(decimals) - len(parts[1])
	} else {
		// There is not a decimal place.
		additionalZeros = int(decimals)
	}
	// Remove the decimal point.
	tmp := strings.ReplaceAll(input, ".", "")
	// Add zeros to ensure that there are an appropriate number of decimals.
	tmp += strings.Repeat("0", additionalZeros)

	// Set the output
	output.SetString(tmp, 10)

	return output, nil
}
