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
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/registrarcontract"
	"github.com/orinocopay/go-etherutils/ens/registrycontract"
)

func RegistryContractAddress(client *ethclient.Client) (address common.Address, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return
	}

	// Instantiate the registry contract
	if chainID.Cmp(params.MainnetChainConfig.ChainId) == 0 {
		address = common.HexToAddress("314159265dd8dbb310642f98f50c066173c1259b")
	} else if chainID.Cmp(params.TestnetChainConfig.ChainId) == 0 {
		address = common.HexToAddress("112234455c3a32fd11230c42e7bccd4a84e02010")
	} else if chainID.Cmp(params.RinkebyChainConfig.ChainId) == 0 {
		address = common.HexToAddress("e7410170f87102DF0055eB195163A03B7F2Bff4A")
	} else {
		err = fmt.Errorf("No contract for network ID %v", chainID)
	}
	return
}

// RegistryContract obtains the registry contract for a chain
func RegistryContract(client *ethclient.Client) (registry *registrycontract.RegistryContract, err error) {
	var address common.Address
	address, err = RegistryContractAddress(client)
	if err != nil {
		return
	}

	// Instantiate the registry contract
	registry, err = registrycontract.NewRegistryContract(address, client)

	return
}

// RegistryContractFromRegistrar obtains the registry contract given an
// existing registrar contract
func RegistryContractFromRegistrar(client *ethclient.Client, registrar *registrarcontract.RegistrarContract) (registry *registrycontract.RegistryContract, err error) {
	registryAddress, err := registrar.Ens(nil)
	if err != nil {
		return
	}
	registry, err = registrycontract.NewRegistryContract(registryAddress, client)
	return
}

// Resolver obtains the address of the resolver for a .eth name
func Resolver(contract *registrycontract.RegistryContract, name string) (address common.Address, err error) {
	address, err = contract.Resolver(nil, NameHash(name))
	if err == nil && bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
	}
	return
}

// SetResolver sets the resolver for a name
func SetResolver(session *registrycontract.RegistryContractSession, name string, resolverAddr *common.Address) (tx *types.Transaction, err error) {
	tx, err = session.SetResolver(NameHash(name), *resolverAddr)
	return
}

// SetSubdomainOwner sets the owner for a subdomain of a name
func SetSubdomainOwner(session *registrycontract.RegistryContractSession, name string, subdomain string, ownerAddr *common.Address) (tx *types.Transaction, err error) {
	tx, err = session.SetSubnodeOwner(NameHash(name), LabelHash(subdomain), *ownerAddr)
	return
}

// CreateRegistrySession creates a session suitable for multiple calls
func CreateRegistrySession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *registrycontract.RegistryContract, gasPrice *big.Int) *registrycontract.RegistryContractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &registrycontract.RegistryContractSession{
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
