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

package ens

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/stretchr/testify/assert"
)

var client, _ = ethclient.Dial("https://ropsten.orinocopay.com:8546/")

func TestSealBid1(t *testing.T) {
	contract, err := RegistrarContract(client)
	assert.Nil(t, err, "Failed to obtain contract")

	address, err := Resolve(client, "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1")
	assert.Nil(t, err, "Failed to obtain address")
	nameHash := LabelHash("foo")
	amount, err := etherutils.StringToWei("0.01 Ether")
	assert.Nil(t, err, "Failed to obtain amount")
	salt := "foo"
	saltHash := saltHash(salt)
	sealedBid1, err := contract.ShaBid(nil, nameHash, address, amount, saltHash)
	assert.Nil(t, err, "Failed to seal bid (1)")

	sealedBid2, err := SealBid("foo.eth", &address, *amount, salt)
	assert.Nil(t, err, "Failed to seal bid (2)")

	assert.Equal(t, hex.EncodeToString(sealedBid1[:]), sealedBid2.Hex()[2:], "Hashes do not match")
}

func TestSealBid2(t *testing.T) {
	expected := "9ab01ba2d5496808710e66c99a2a79e78e7f5002b19a5590a33042bc9ca03583"

	address, err := Resolve(client, "0x388ea662ef2c223ec0b047d41bf3c0f362142ad5")
	assert.Nil(t, err, "Failed to obtain address")
	amount, err := etherutils.StringToWei("0.01 Ether")
	assert.Nil(t, err, "Failed to obtain amount")
	salt := "foo"
	actual, err := SealBid("testtest.eth", &address, *amount, salt)
	assert.Nil(t, err, "Failed to hash bid")

	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}
