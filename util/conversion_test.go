package util

import (
	"math/big"
	"testing"
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
