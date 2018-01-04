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
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/resolvercontract"
)

var zeroHash = make([]byte, 32)

// UnknownAddress is the address to which unknown entries resolve
var UnknownAddress = common.HexToAddress("00")

// PublicResolver obtains the public resolver for a chain
func PublicResolver(client *ethclient.Client) (address common.Address, err error) {
	address, err = resolverAddress(client, "resolver.eth")

	return
}

func resolverAddress(client *ethclient.Client, name string) (address common.Address, err error) {
	nameHash := NameHash(name)

	registryContract, err := RegistryContract(client)
	if err != nil {
		return
	}

	// Check that this name is owned
	ownerAddress, err := registryContract.Owner(nil, nameHash)
	if err != nil {
		return
	}
	if bytes.Compare(ownerAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("unregistered name")
		return
	}

	// Obtain the resolver address for this name
	address, err = registryContract.Resolver(nil, nameHash)
	if err != nil {
		return
	}
	if bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
		return
	}

	return
}

// Resolve resolves an ENS name in to an Etheruem address
// This will return an error if the name is not found or otherwise 0
func Resolve(client *ethclient.Client, input string) (address common.Address, err error) {
	if strings.HasSuffix(input, ".eth") {
		return resolveName(client, input)
	}
	if (strings.HasPrefix(input, "0x") && len(input) > 42) || (!strings.HasPrefix(input, "0x") && len(input) > 40) {
		err = errors.New("address too long")
	} else {
		address = common.HexToAddress(input)
		if address == UnknownAddress {
			err = errors.New("could not parse address")
		}
	}

	return
}

func resolveName(client *ethclient.Client, input string) (address common.Address, err error) {
	var nameHash [32]byte
	nameHash = NameHash(input)
	if bytes.Compare(nameHash[:], zeroHash) == 0 {
		err = errors.New("Bad name")
	} else {
		address, err = resolveHash(client, input)
	}
	return
}

func resolveHash(client *ethclient.Client, name string) (address common.Address, err error) {
	contract, err := ResolverContract(client, name)
	if err != nil {
		return UnknownAddress, err
	}

	// Resolve the name
	address, err = contract.Addr(nil, NameHash(name))
	if err != nil {
		return UnknownAddress, err
	}
	if bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		return UnknownAddress, errors.New("no address")
	}

	return
}

// CreateResolverSession creates a session suitable for multiple calls
func CreateResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *resolvercontract.ResolverContract, gasPrice *big.Int) *resolvercontract.ResolverContractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &resolvercontract.ResolverContractSession{
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

// SetResolution sets the address to which a name resolves
func SetResolution(session *resolvercontract.ResolverContractSession, name string, resolutionAddress *common.Address) (tx *types.Transaction, err error) {
	tx, err = session.SetAddr(NameHash(name), *resolutionAddress)
	return
}

// SetAbi sets the ABI associated with a name
func SetAbi(session *resolvercontract.ResolverContractSession, name string, abi string, contentType *big.Int) (tx *types.Transaction, err error) {
	var data []byte
	if contentType.Cmp(big.NewInt(1)) == 0 {
		// Uncompressed JSON
		data = []byte(abi)
	} else if contentType.Cmp(big.NewInt(2)) == 0 {
		// Zlib-compressed JSON
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write([]byte(abi))
		w.Close()
		data = b.Bytes()
	} else {
		err = errors.New("Unsupported content type")
	}
	// Gas cost is around 50000 + 64 per byte; add 10000 headroom to be safe
	//session.TransactOpts.GasLimit = big.NewInt(int64(600000 + len(data)*64))
	tx, err = session.SetABI(NameHash(name), contentType, data)

	return
}

// Abi returns the ABI associated with a name
func Abi(resolver *resolvercontract.ResolverContract, name string) (abi string, err error) {
	var result struct {
		ContentType *big.Int
		Data        []byte
	}
	contentTypes := big.NewInt(3)
	result, err = resolver.ABI(nil, NameHash(name), contentTypes)
	if err == nil {
		if result.ContentType.Cmp(big.NewInt(1)) == 0 {
			// Uncompressed JSON
			abi = string(result.Data)
		} else if result.ContentType.Cmp(big.NewInt(2)) == 0 {
			// Zlib-compressed JSON
			b := bytes.NewReader(result.Data)
			var z io.ReadCloser
			z, err = zlib.NewReader(b)
			if err != nil {
				return
			}
			defer z.Close()
			var uncompressed []byte
			uncompressed, err = ioutil.ReadAll(z)
			abi = string(uncompressed)
		}
	}
	return
}

// ResolverContractByAddress instantiates the resolver contract at aspecific address
func ResolverContractByAddress(client *ethclient.Client, resolverAddress common.Address) (resolver *resolvercontract.ResolverContract, err error) {
	// Instantiate the resolver contract
	resolver, err = resolvercontract.NewResolverContract(resolverAddress, client)

	return
}

// ResolverContract obtains the resolver contract for a name
func ResolverContract(client *ethclient.Client, name string) (resolver *resolvercontract.ResolverContract, err error) {
	resolverAddress, err := resolverAddress(client, name)
	if err != nil {
		return
	}
	if bytes.Compare(resolverAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		err = errors.New("no resolver")
		return
	}

	resolver, err = ResolverContractByAddress(client, resolverAddress)
	return
}
