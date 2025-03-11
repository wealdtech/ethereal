// Copyright © 2025 Weald Technology Trading
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

package conn

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Address obtains the address given a range of inputs.
func (c *Conn) Address(address string, privateKey string) (common.Address, error) {
	switch {
	case address != "":
		return c.Resolve(address)
	case privateKey != "":
		return c.addressFromPrivateKey(privateKey)
	default:
		return common.Address{}, errors.New("neither address nor private key provided")
	}
}

func (c *Conn) addressFromPrivateKey(privateKey string) (common.Address, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return common.Address{}, errors.Join(errors.New("failed to obtain public key from private key"), err)
	}

	return crypto.PubkeyToAddress(key.PublicKey), nil
}
