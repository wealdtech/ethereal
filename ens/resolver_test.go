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

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestResolveEmpty(t *testing.T) {
	_, err := Resolve(client, "")
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveZero(t *testing.T) {
	_, err := Resolve(client, "0")
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveNotPresent(t *testing.T) {
	_, err := Resolve(client, "sirnotappearinginthisregistry.eth")
	assert.NotNil(t, err, "Resolved name that does not exist")
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
}

func TestResolveNoResolver(t *testing.T) {
	_, err := Resolve(client, "noresolver.eth")
	assert.NotNil(t, err, "Resolved name without a resolver")
	assert.Equal(t, "no resolver", err.Error(), "Unexpected error")
}

func TestResolveBadResolver(t *testing.T) {
	_, err := Resolve(client, "resolvestozero.eth")
	assert.NotNil(t, err, "Resolved name with a bad resolver")
	assert.Equal(t, "no address", err.Error(), "Unexpected error")
}

func TestResolveTestEnsTest(t *testing.T) {
	expected := "388ea662ef2c223ec0b047d41bf3c0f362142ad5"
	actual, err := Resolve(client, "test.enstest.eth")
	assert.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveResolverEth(t *testing.T) {
	expected := "5ffc014343cd971b7eb70732021e26c35b744cc4"
	actual, err := Resolve(client, "resolver.eth")
	assert.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveNickJohnson(t *testing.T) {
	expected := "70abd981e01ad3e6eb1726a5001000877ab04417"
	actual, err := Resolve(client, "nickjohnson.eth")
	assert.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveAddress(t *testing.T) {
	expected := "5ffc014343cd971b7eb70732021e26c35b744cc4"
	actual, err := Resolve(client, "0x5ffc014343cd971b7eb70732021e26c35b744cc4")
	assert.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveShortAddress(t *testing.T) {
	expected := "0000000000000000000000000000000000000001"
	actual, err := Resolve(client, "0x1")
	assert.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveHexString(t *testing.T) {
	_, err := Resolve(client, "0xe32c6d1a964749b6de2130e20daed821a45b9e7261118801ff5319d0ffc6b54a")
	assert.NotNil(t, err, "Resolved too-long hex string")
}

func TestReverseResolveTestEnsTest(t *testing.T) {
	expected := "domainsale"
	address := common.HexToAddress("0x388ea662ef2c223ec0b047d41bf3c0f362142ad5")
	actual, err := ReverseResolve(client, &address)
	assert.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}
