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
	"errors"
	"math/big"
	"math/rand"
	"time"

	"github.com/dchest/uniuri"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/ens/registrarcontract"
	"github.com/orinocopay/go-etherutils/ens/registrycontract"
)

func RegistrarContractAddress(client *ethclient.Client) (address common.Address, err error) {
	return RegistrarContractAddressFor(client, "eth")
}

func RegistrarContractAddressFor(client *ethclient.Client, root string) (address common.Address, err error) {
	// Obtain a registry contract
	registry, err := RegistryContract(client)
	if err != nil {
		return
	}

	// Obtain the registrar address from the registry
	address, err = registry.Owner(nil, NameHash(root))
	if err != nil {
		return
	}
	if address == UnknownAddress {
		err = errors.New("no registrar for that network")
	}

	return
}

// RegistrarContract obtains the registrar contract for '.eth'
func RegistrarContract(client *ethclient.Client) (registrar *registrarcontract.RegistrarContract, err error) {
	return RegistrarContractFor(client, "eth")
}

// RegistrarContract obtains the registrar contract for a named root
func RegistrarContractFor(client *ethclient.Client, root string) (registrar *registrarcontract.RegistrarContract, err error) {
	var address common.Address
	address, err = RegistrarContractAddressFor(client, root)
	if err != nil {
		return
	}

	registrar, err = registrarcontract.NewRegistrarContract(address, client)
	return
}

// CreateRegistrarSession creates a session suitable for multiple calls
func CreateRegistrarSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *registrarcontract.RegistrarContract, gasPrice *big.Int) *registrarcontract.RegistrarContractSession {
	// Create a signer
	signer := etherutils.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &registrarcontract.RegistrarContractSession{
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

// SealBid seals the elements of a bid in to a single hash
func SealBid(name string, owner *common.Address, amount big.Int, salt string) (hash common.Hash, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}
	domainHash := LabelHash(domain)

	sha := sha3.NewKeccak256()
	sha.Write(domainHash[:])
	sha.Write(owner.Bytes())
	// Amount needs to be exactly 32 bytes
	var amountBytes [32]byte
	copy(amountBytes[len(amountBytes)-len(amount.Bytes()):], amount.Bytes()[:])
	sha.Write(amountBytes[:])
	saltHash := saltHash(salt)
	sha.Write(saltHash[:])
	sha.Sum(hash[:0])
	return
}

// StartAuction starts an auction without bidding
func StartAuction(session *registrarcontract.RegistrarContractSession, name string) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	tx, err = session.StartAuction(LabelHash(domain))
	return
}

// StartAuctionAndBid starts an auction and bids in the same transaction.
func StartAuctionAndBid(session *registrarcontract.RegistrarContractSession, name string, owner *common.Address, amount big.Int, salt string, dummies int) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	sealedBid, err := SealBid(name, owner, amount, salt)
	if err != nil {
		return
	}

	var domainHashes [][32]byte
	domainHashes = make([][32]byte, 0, dummies+1)
	rand.Seed(time.Now().UnixNano())
	namePlace := rand.Intn(dummies + 1)
	for i := 0; i < dummies+1; i++ {
		var thisDomain string
		if i == namePlace {
			thisDomain = domain
		} else {
			thisDomain = uniuri.New()
		}
		domainHashes = append(domainHashes, LabelHash(thisDomain))
	}

	tx, err = session.StartAuctionsAndBid(domainHashes, sealedBid)
	return
}

// InvalidateName invalidates a non-conformant ENS registration.
func InvalidateName(session *registrarcontract.RegistrarContractSession, name string) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	tx, err = session.InvalidateName(domain)
	return
}

// NewBid bids on an existing auction
func NewBid(session *registrarcontract.RegistrarContractSession, name string, owner *common.Address, amount big.Int, salt string) (tx *types.Transaction, err error) {
	sealedBid, err := SealBid(name, owner, amount, salt)
	if err != nil {
		return
	}

	tx, err = session.NewBid(sealedBid)
	return
}

// RevealBid reveals an existing bid on an existing auction
func RevealBid(session *registrarcontract.RegistrarContractSession, name string, owner *common.Address, amount big.Int, salt string) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}
	domainHash := LabelHash(domain)
	saltHash := saltHash(salt)

	session.TransactOpts.GasLimit = 200000
	tx, err = session.UnsealBid(domainHash, &amount, saltHash)
	return
}

// FinishAuction reveals an existing bid on an existing auction
func FinishAuction(session *registrarcontract.RegistrarContractSession, name string) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	tx, err = session.FinalizeAuction(LabelHash(domain))
	return
}

func Transfer(session *registrarcontract.RegistrarContractSession, name string, to common.Address) (tx *types.Transaction, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	tx, err = session.Transfer(LabelHash(domain), to)
	return
}

// Entry obtains a registrar entry for a name
func Entry(contract *registrarcontract.RegistrarContract, client *ethclient.Client, name string) (state string, deedAddress common.Address, registrationDate time.Time, value *big.Int, highestBid *big.Int, err error) {
	domain, err := Domain(name)
	if err != nil {
		err = errors.New("invalid name")
		return
	}

	status, deedAddress, registration, value, highestBid, err := contract.Entries(nil, LabelHash(domain))
	if err != nil {
		return
	}
	registrationDate = time.Unix(registration.Int64(), 0)
	switch status {
	case 0:
		state = "Available"
	case 1:
		state = "Bidding"
	case 2:
		// Might be won or owned
		var registryContract *registrycontract.RegistryContract
		registryContract, err = RegistryContractFromRegistrar(client, contract)
		if err != nil {
			return
		}

		var owner common.Address
		owner, err = registryContract.Owner(nil, NameHash(name))
		if err != nil {
			return
		}

		if owner == UnknownAddress {
			state = "Won"
		} else {
			state = "Owned"
		}
	case 3:
		state = "Forbidden"
	case 4:
		state = "Revealing"
	case 5:
		state = "Unavailable"
	default:
		state = "Unknown"
	}
	return
}

// State obains the current state of a name
func State(contract *registrarcontract.RegistrarContract, client *ethclient.Client, name string) (state string, err error) {
	state, _, _, _, _, err = Entry(contract, client, name)

	return
}

// NameInState checks if a name is in a given state, and errors if not.
func NameInState(contract *registrarcontract.RegistrarContract, client *ethclient.Client, name string, desiredState string) (inState bool, err error) {
	state, err := State(contract, client, name)
	if err == nil {
		if state == desiredState {
			inState = true
		} else {
			switch state {
			case "Available":
				err = errors.New("this name has not been auctioned")
			case "Bidding":
				err = errors.New("this name is being auctioned")
			case "Won":
				err = errors.New("this name's auction has finished")
			case "Owned":
				err = errors.New("this name is owned")
			case "Forbidden":
				err = errors.New("this name is unavailable")
			case "Revealing":
				err = errors.New("this name is being revealed")
			case "Unavailable":
				err = errors.New("this name is not yet available")
			default:
				err = errors.New("this name is in an unknown state")
			}
		}
	}
	return
}

// Generate a simple hash for a salt
func saltHash(salt string) (hash [32]byte) {
	sha := sha3.NewKeccak256()
	sha.Write([]byte(salt))
	sha.Sum(hash[:0])
	return
}
