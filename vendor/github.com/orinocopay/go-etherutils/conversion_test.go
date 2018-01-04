// Copyright 2017 Orinoco Payments
//
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

package etherutils

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToWeiWithNormalValue(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000", 10)
	result, err := StringToWei("1000")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithLargeValue(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000000000000000000", 10)
	result, err := StringToWei("1000000000000000000000")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiProb1(t *testing.T) {
	expected, _ := new(big.Int).SetString("24000000000000000", 10)
	result, err := StringToWei("0.024ether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithWeiUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000000000000000000", 10)
	result, err := StringToWei("1000000000000000000000 Wei")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithMicroEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("85748574000000000000", 10)
	result, err := StringToWei("85748574 microether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithMilliEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("85748574000000000000000", 10)
	result, err := StringToWei("85748574 milliether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000000000000000", 10)
	result, err := StringToWei("1 ether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithKiloEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000000000000000000", 10)
	result, err := StringToWei("1 kiloether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithMegaEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	result, err := StringToWei("1 megaether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithGigaEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	result, err := StringToWei("1 gigaether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithTeraEtherUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("5000000000000000000000000000000000", 10)
	result, err := StringToWei("5000 Teraether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithKWeiUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("100", 10)
	result, err := StringToWei("0.1 kwei")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithKiloEtherDecimal1(t *testing.T) {
	expected, _ := new(big.Int).SetString("100000000000000000", 10)
	result, err := StringToWei("0.0001 kiloether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithEtherDecimal2(t *testing.T) {
	expected, _ := new(big.Int).SetString("100000000000000000", 10)
	result, err := StringToWei(".0000001 megaether")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithMWeiUnitMissingDecimal(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000000", 10)
	result, err := StringToWei("1. Mwei")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithGWeiUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("21000000000", 10)
	result, err := StringToWei("21 Gwei")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiNoUnit(t *testing.T) {
	expected, _ := new(big.Int).SetString("1000", 10)
	result, err := StringToWei("1000 ")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiNoSpace(t *testing.T) {
	expected, _ := new(big.Int).SetString("2000000", 10)
	result, err := StringToWei("2megawei")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWhitespace(t *testing.T) {
	expected, _ := new(big.Int).SetString("2000000", 10)
	result, err := StringToWei("   2 megawei   ")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithSplitUnits(t *testing.T) {
	expected, _ := new(big.Int).SetString("2000000", 10)
	result, err := StringToWei("2 mega wei")
	assert.Nil(t, err, "Failed to convert normal string to Wei")
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestStringToWeiWithEmptyInput(t *testing.T) {
	_, err := StringToWei("")
	assert.NotNil(t, err, "Converted empty string to Wei value")
}

func TestStringToWeiWithUnknownUnit(t *testing.T) {
	_, err := StringToWei("1000 foo")
	assert.NotNil(t, err, "Converted string with bad unit to Wei value")
}

func TestStringToWeiWithUnknownDecimalUnit(t *testing.T) {
	_, err := StringToWei("1000.5 foo")
	assert.NotNil(t, err, "Converted string with bad unit to Wei value")
}

func TestStringToWeiWithBadNumeric(t *testing.T) {
	_, err := StringToWei("onehundred ether")
	assert.NotNil(t, err, "Converted string with bad number to Wei value")
}

func TestStringToWeiWithBadDecimalNumeric(t *testing.T) {
	_, err := StringToWei("onehundred.5 ether")
	assert.NotNil(t, err, "Converted string with bad number to Wei value")
}

func TestStringToWeiWithWeiDecimal1(t *testing.T) {
	_, err := StringToWei("0.0001 kwei")
	assert.NotNil(t, err, "Converted string with decimal Wei to Wei value")
}

func TestStringToWeiWithWeiNegative(t *testing.T) {
	_, err := StringToWei("-2 wei")
	assert.NotNil(t, err, "Converted string with negative Wei to Wei value")
}

func TestWeiToStringWithZero(t *testing.T) {
	expected := "0"
	wei := big.NewInt(0)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithNegativeValue(t *testing.T) {
	expected := "-12.345 KWei"
	wei := big.NewInt(-12345)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithSmallValue(t *testing.T) {
	expected := "1 Wei"
	wei := big.NewInt(1)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithLargerValue(t *testing.T) {
	expected := "2.034 KWei"
	wei := big.NewInt(2034)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithExpectedDecimalValue(t *testing.T) {
	expected := "1.23456789 GWei"
	wei := big.NewInt(1234567890)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithEtherValue(t *testing.T) {
	expected := "1 Ether"
	wei := big.NewInt(1000000000000000000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithEtherDecimalValue(t *testing.T) {
	expected := "1.000000000000000001 Ether"
	wei := big.NewInt(1000000000000000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard1(t *testing.T) {
	expected := "1 Wei"
	wei := big.NewInt(1)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard2(t *testing.T) {
	expected := "999 Wei"
	wei := big.NewInt(999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard3(t *testing.T) {
	expected := "1 KWei"
	wei := big.NewInt(1000)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard4(t *testing.T) {
	expected := "1.001 KWei"
	wei := big.NewInt(1001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard5(t *testing.T) {
	expected := "999.999 KWei"
	wei := big.NewInt(999999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard6(t *testing.T) {
	expected := "1 MWei"
	wei := big.NewInt(1000000)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard7(t *testing.T) {
	expected := "1.000001 MWei"
	wei := big.NewInt(1000001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard8(t *testing.T) {
	expected := "999.999999 MWei"
	wei := big.NewInt(999999999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard9(t *testing.T) {
	expected := "1 GWei"
	wei := big.NewInt(1000000000)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard10(t *testing.T) {
	expected := "1.000000001 GWei"
	wei := big.NewInt(1000000001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard11(t *testing.T) {
	expected := "999.999999999 GWei"
	wei := big.NewInt(999999999999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard12(t *testing.T) {
	expected := "0.000001 Ether"
	wei := big.NewInt(1000000000000)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard13(t *testing.T) {
	expected := "0.000001000000000001 Ether"
	wei := big.NewInt(1000000000001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard14(t *testing.T) {
	expected := "0.000999999999999999 Ether"
	wei := big.NewInt(999999999999999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard15(t *testing.T) {
	expected := "0.000001 Ether"
	wei := big.NewInt(1000000000000)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard16(t *testing.T) {
	expected := "0.000001000000000001 Ether"
	wei := big.NewInt(1000000000001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard17(t *testing.T) {
	expected := "0.000999999999999999 Ether"
	wei := big.NewInt(999999999999999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard18(t *testing.T) {
	expected := "1 Ether"
	wei := big.NewInt(1000000000000000000)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard19(t *testing.T) {
	expected := "0.001000000000000001 Ether"
	wei := big.NewInt(1000000000000001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard20(t *testing.T) {
	expected := "0.999999999999999999 Ether"
	wei := big.NewInt(999999999999999999)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard21(t *testing.T) {
	expected := "1000 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000000000", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard22(t *testing.T) {
	expected := "1.000000000000000001 Ether"
	wei := big.NewInt(1000000000000000001)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard23(t *testing.T) {
	expected := "999.999999999999999999 Ether"
	wei, _ := new(big.Int).SetString("999999999999999999999", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard24(t *testing.T) {
	expected := "1000000 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard25(t *testing.T) {
	expected := "1000.000000000000000001 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000000001", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard26(t *testing.T) {
	expected := "999999.999999999999999999 Ether"
	wei, _ := new(big.Int).SetString("999999999999999999999999", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard27(t *testing.T) {
	expected := "1000000000 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard28(t *testing.T) {
	expected := "1000000.000000000000000001 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000000000001", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard29(t *testing.T) {
	expected := "999999999.999999999999999999 Ether"
	wei, _ := new(big.Int).SetString("999999999999999999999999999", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithStandard30(t *testing.T) {
	expected := "1000000000000 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000000000000000000", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard1(t *testing.T) {
	expected := "1 Wei"
	wei := big.NewInt(1)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard2(t *testing.T) {
	expected := "999 Wei"
	wei := big.NewInt(999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard3(t *testing.T) {
	expected := "1 KWei"
	wei := big.NewInt(1000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard4(t *testing.T) {
	expected := "1.001 KWei"
	wei := big.NewInt(1001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard5(t *testing.T) {
	expected := "999.999 KWei"
	wei := big.NewInt(999999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard6(t *testing.T) {
	expected := "1 MWei"
	wei := big.NewInt(1000000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard7(t *testing.T) {
	expected := "1.000001 MWei"
	wei := big.NewInt(1000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard8(t *testing.T) {
	expected := "999.999999 MWei"
	wei := big.NewInt(999999999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard9(t *testing.T) {
	expected := "1 GWei"
	wei := big.NewInt(1000000000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard10(t *testing.T) {
	expected := "1.000000001 GWei"
	wei := big.NewInt(1000000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard11(t *testing.T) {
	expected := "999.999999999 GWei"
	wei := big.NewInt(999999999999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard12(t *testing.T) {
	expected := "1 Microether"
	wei := big.NewInt(1000000000000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard13(t *testing.T) {
	expected := "1.000000000001 Microether"
	wei := big.NewInt(1000000000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard14(t *testing.T) {
	expected := "999.999999999999 Microether"
	wei := big.NewInt(999999999999999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard15(t *testing.T) {
	expected := "1 Microether"
	wei := big.NewInt(1000000000000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard16(t *testing.T) {
	expected := "1.000000000001 Microether"
	wei := big.NewInt(1000000000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard17(t *testing.T) {
	expected := "999.999999999999 Microether"
	wei := big.NewInt(999999999999999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard18(t *testing.T) {
	expected := "1 Ether"
	wei := big.NewInt(1000000000000000000)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard19(t *testing.T) {
	expected := "1.000000000000001 Milliether"
	wei := big.NewInt(1000000000000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard20(t *testing.T) {
	expected := "999.999999999999999 Milliether"
	wei := big.NewInt(999999999999999999)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard21(t *testing.T) {
	expected := "1 Kiloether"
	wei, _ := new(big.Int).SetString("1000000000000000000000", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard22(t *testing.T) {
	expected := "1.000000000000000001 Ether"
	wei := big.NewInt(1000000000000000001)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard23(t *testing.T) {
	expected := "999.999999999999999999 Ether"
	wei, _ := new(big.Int).SetString("999999999999999999999", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard24(t *testing.T) {
	expected := "1 Megaether"
	wei, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard25(t *testing.T) {
	expected := "1.000000000000000000001 Kiloether"
	wei, _ := new(big.Int).SetString("1000000000000000000001", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard26(t *testing.T) {
	expected := "999.999999999999999999999 Kiloether"
	wei, _ := new(big.Int).SetString("999999999999999999999999", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard27(t *testing.T) {
	expected := "1 Gigaether"
	wei, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard28(t *testing.T) {
	expected := "1.000000000000000000000001 Megaether"
	wei, _ := new(big.Int).SetString("1000000000000000000000001", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard29(t *testing.T) {
	expected := "999.999999999999999999999999 Megaether"
	wei, _ := new(big.Int).SetString("999999999999999999999999999", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithoutStandard30(t *testing.T) {
	expected := "1 Teraether"
	wei, _ := new(big.Int).SetString("1000000000000000000000000000000", 10)
	result := WeiToString(wei, false)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestWeiToStringWithSmallEtherDecimalValue(t *testing.T) {
	expected := "1.000000000000000001 Ether"
	wei, _ := new(big.Int).SetString("1000000000000000001", 10)
	result := WeiToString(wei, true)
	assert.Equal(t, expected, result, "Did not receive expected result")
}

func TestRoundTripWithSmallValue(t *testing.T) {
	first := "1 Wei"
	second, err := StringToWei(first)
	assert.Nil(t, err, "Failed to convert Ether to Wei")
	assert.Equal(t, second, big.NewInt(1), "Unexpected result converting Ether to Wei")
	third := WeiToString(second, false)
	assert.Equal(t, first, third, "Did not receive expected result")
	fourth := WeiToString(second, true)
	assert.Equal(t, first, fourth, "Did not receive expected result")
}

func TestRoundTripWithNormalValue(t *testing.T) {
	first := "1 Ether"
	second, err := StringToWei(first)
	assert.Nil(t, err, "Failed to convert Ether to Wei")
	assert.Equal(t, second, big.NewInt(1000000000000000000), "Unexpected result converting Ether to Wei")
	third := WeiToString(second, true)
	assert.Equal(t, first, third, "Did not receive expected result")
}

func ExampleUnitToMultiplier() {
	multiplier, err := UnitToMultiplier("ether")
	if err != nil {
		return
	}
	fmt.Println(multiplier.Text(10))
	// Output: 1000000000000000000
}
