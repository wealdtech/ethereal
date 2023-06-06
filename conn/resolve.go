// Copyright © 2022 Weald Technology Trading
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
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

// Resolve resolves a name to an address.
func (c *Conn) Resolve(name string) (common.Address, error) {
	return ens.Resolve(c.client, name)
}

// ReverseResolve resolves an address to a name.
func (c *Conn) ReverseResolve(address common.Address) (string, error) {
	return ens.ReverseResolve(c.client, address)
}
