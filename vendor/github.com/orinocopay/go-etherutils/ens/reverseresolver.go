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
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/reverseregistrarcontract"
	"github.com/orinocopay/go-etherutils/ens/reverseresolvercontract"
)

// ReverseResolve resolves an address in to an ENS name
// This will return an error if the name is not found or otherwise 0
func ReverseResolve(client *ethclient.Client, input *common.Address) (name string, err error) {
	if input == nil {
		err = errors.New("No address supplied")
		return
	}

	nameHash := NameHash(input.Hex()[2:] + ".addr.reverse")

	contract, err := ReverseResolver(client)
	if err != nil {
		return
	}

	// Resolve the name
	name, err = contract.Name(nil, nameHash)

	if name == "" {
		err = errors.New("No resolution")
	}

	return
}

// ReverseResolver obtains the reverse resolver contract
func ReverseResolver(client *ethclient.Client) (resolver *reverseresolvercontract.ReverseResolver, err error) {
	registryContract, err := RegistryContract(client)
	if err != nil {
		return
	}

	// Obtain the reverse registrar address
	reverseRegistrarAddress, err := registryContract.Owner(nil, NameHash("addr.reverse"))
	if err != nil {
		return
	}
	if bytes.Compare(reverseRegistrarAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("unregistered name")
		return
	}

	// Instantiate the reverse registrar contract
	reverseRegistrarContract, err := reverseregistrarcontract.NewReverseRegistrarContract(reverseRegistrarAddress, client)
	if err != nil {
		return
	}

	// Now fetch the default resolver
	reverseResolverAddress, err := reverseRegistrarContract.DefaultResolver(nil)
	if err != nil {
		return
	}

	// Finally we can obtain the resolver itself
	resolver, err = reverseresolvercontract.NewReverseResolver(reverseResolverAddress, client)

	return
}

// CreateReverseResolverSession creates a session suitable for multiple calls
func CreateReverseResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *reverseresolvercontract.ReverseResolver, gasPrice *big.Int) *reverseresolvercontract.ReverseResolverSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &reverseresolvercontract.ReverseResolverSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     account.Address,
			Signer:   signer,
			GasPrice: gasPrice,
		},
	}

	return session
}
