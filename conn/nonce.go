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
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// CurrentNonce provides the current nonce for the given address.
func (c *Conn) CurrentNonce(ctx context.Context,
	address common.Address,
) (
	uint64,
	error,
) {
	c.noncesMu.Lock()
	defer c.noncesMu.Unlock()

	return c.currentNonce(ctx, address)
}

// currentNonce obtains the current nonce.
// It assumes that the nonces lock is held.
func (c *Conn) currentNonce(ctx context.Context,
	address common.Address,
) (
	uint64,
	error,
) {
	_, exists := c.nonces[address]
	if !exists {
		if c.client == nil {
			// Offline, fetch from supplied value.
			tmp := viper.GetString("nonce")
			if tmp == "" {
				return 0, errors.New("nonce not supplied")
			}
			nonce, err := strconv.ParseUint(tmp, 10, 64)
			if err != nil {
				return 0, errors.Wrap(err, "invalid nonce")
			}
			c.nonces[address] = nonce
		} else {
			// Fetch from chain.
			ctx, cancel := context.WithTimeout(ctx, c.timeout)
			defer cancel()
			nonce, err := c.client.PendingNonceAt(ctx, address)
			if err != nil {
				return 0, errors.Wrap(err, fmt.Sprintf("failed to obtain nonce for %s", address.Hex()))
			}
			c.nonces[address] = nonce
		}
	}
	return c.nonces[address], nil
}

// NextNonce obtains the next nonce for the given address.
func (c *Conn) NextNonce(ctx context.Context,
	address common.Address,
) (
	uint64,
	error,
) {
	c.noncesMu.Lock()
	defer c.noncesMu.Unlock()

	currentNonce, err := c.currentNonce(ctx, address)
	if err != nil {
		return 0, err
	}

	currentNonce++
	c.nonces[address]++

	return currentNonce, nil
}
