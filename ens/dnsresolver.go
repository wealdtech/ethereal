// Copyright 2017 Weald Technology Trading
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
	"github.com/wealdtech/ethereal/ens/dnsresolvercontract"
	"github.com/wealdtech/ethereal/util"
)

// CreateDNSResolverSession creates a session suitable for multiple calls
func CreateDNSResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *dnsresolvercontract.DNSResolverContract, gasPrice *big.Int) *dnsresolvercontract.DNSResolverContractSession {
	// Create a signer
	signer := util.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &dnsresolvercontract.DNSResolverContractSession{
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

// DNSRecord fetches a DNS record
func DNSRecord(client *ethclient.Client, domain string, name string, rrType uint16) (data []byte, err error) {
	contract, err := DNSResolverContract(client, domain)
	if err == nil {
		data, err = contract.DnsRecord(nil, NameHash(domain), LabelHash(name), rrType)
	}
	return
}

// SetDNSRecords sets DNS records
func SetDNSRecords(session *dnsresolvercontract.DNSResolverContractSession, domain string, data []byte) (tx *types.Transaction, err error) {
	tx, err = session.SetDNSRecords(NameHash(domain), data)
	return
}

// DNSResolverContractByAddress instantiates the resolver contract at aspecific address
func DNSResolverContractByAddress(client *ethclient.Client, resolverAddress common.Address) (resolver *dnsresolvercontract.DNSResolverContract, err error) {
	// Instantiate the resolver contract
	resolver, err = dnsresolvercontract.NewDNSResolverContract(resolverAddress, client)
	if err != nil {
		return
	}

	// Ensure that this is a DNS resolver
	var supported bool
	supported, err = resolver.SupportsInterface(nil, [4]byte{0xa8, 0xfa, 0x56, 0x82})
	if err != nil {
		return
	}
	if !supported {
		err = fmt.Errorf("%s is not a DNS resolver contract", resolverAddress.Hex())
	}

	return
}

// DNSResolverContract obtains the resolver contract for a domain
func DNSResolverContract(client *ethclient.Client, domain string) (resolver *dnsresolvercontract.DNSResolverContract, err error) {
	resolverAddress, err := resolverAddress(client, domain)
	if err != nil {
		return
	}
	if bytes.Compare(resolverAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
		return
	}

	resolver, err = DNSResolverContractByAddress(client, resolverAddress)
	return
}
