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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/orinocopay/go-etherutils/ens/deedcontract"
	"github.com/orinocopay/go-etherutils/ens/registrarcontract"
)

// DeedContract obtains the deed contract at a particular address
func DeedContract(client *ethclient.Client, address *common.Address) (deed *deedcontract.DeedContract, err error) {
	deed, err = deedcontract.NewDeedContract(*address, client)
	return
}

// DeedContract obtains the deed contract for a particular name
func DeedContractFor(client *ethclient.Client, registrar *registrarcontract.RegistrarContract, name string) (deedContract *deedcontract.DeedContract, err error) {
	_, deedAddress, _, _, _, err := Entry(registrar, client, name)
	if err != nil {
		return
	}
	deedContract, err = DeedContract(client, &deedAddress)

	return
}

// Owner obtains the owner of a deed
func Owner(contract *deedcontract.DeedContract) (address common.Address, err error) {
	address, err = contract.Owner(nil)
	return
}

// PreviousOwner obtains the previous owner of a deed
func PreviousOwner(contract *deedcontract.DeedContract) (address common.Address, err error) {
	address, err = contract.PreviousOwner(nil)
	return
}
