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
