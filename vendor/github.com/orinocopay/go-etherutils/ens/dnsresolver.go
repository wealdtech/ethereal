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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/dnsresolvercontract"
)

// CreateDnsResolverSession creates a session suitable for multiple calls
func CreateDnsResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *dnsresolvercontract.DnsResolverContract, gasPrice *big.Int) *dnsresolvercontract.DnsResolverContractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &dnsresolvercontract.DnsResolverContractSession{
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

// SetDns sets a DNS resolution
func SetDns(session *dnsresolvercontract.DnsResolverContractSession, name string, rrType uint16, key string, data []byte) (tx *types.Transaction, err error) {
	tx, err = session.SetDns(NameHash(name), rrType, key, data)
	return
}

func Dns(client *ethclient.Client, name string, rrType uint16, key string) (data []byte, err error) {
	contract, err := DnsResolverContract(client, name)
	if err == nil {
		data, err = contract.Dns(nil, NameHash(name), rrType, key)
	}
	return
}

// DnsResolverContractByAddress instantiates the resolver contract at aspecific address
func DnsResolverContractByAddress(client *ethclient.Client, resolverAddress common.Address) (resolver *dnsresolvercontract.DnsResolverContract, err error) {
	// Instantiate the resolver contract
	resolver, err = dnsresolvercontract.NewDnsResolverContract(resolverAddress, client)
	if err != nil {
		return
	}

	// Ensure that this is a DNS resolver
	var supported bool
	supported, err = resolver.SupportsInterface(nil, [4]byte{0xaf, 0x6e, 0x6e, 0x9e})
	if err != nil {
		return
	}
	if !supported {
		err = fmt.Errorf("%s is not a DNS resolver contract", resolverAddress.Hex())
	}

	return
}

// DnsResolverContract obtains the resolver contract for a name
func DnsResolverContract(client *ethclient.Client, name string) (resolver *dnsresolvercontract.DnsResolverContract, err error) {
	resolverAddress, err := resolverAddress(client, name)
	if err != nil {
		return
	}
	if bytes.Compare(resolverAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
		return
	}

	resolver, err = DnsResolverContractByAddress(client, resolverAddress)
	return
}
