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

	"github.com/ethereum/go-ethereum"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// EstimateGas estimates the gas required for the given transaction.
func (c *Conn) EstimateGas(ctx context.Context,
	txData *TransactionData,
) (
	uint64,
	error,
) {
	if c.client == nil {
		// We're offline; fetch from input.
		gasLimit := viper.GetInt64("gaslimit")
		if gasLimit <= 0 {
			return 0, errors.New("gas limit not specified")
		}
		return uint64(gasLimit), nil
	}

	msg := ethereum.CallMsg{From: txData.From, To: txData.To, Value: txData.Value, Data: txData.Data}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	gas, err := c.client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, errors.Wrap(err, "failed to estimate gas")
	}
	return gas, err
}
