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
	"math/big"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wealdtech/go-string2eth"
)

// CalculateFees calculates the base and priority fees.
func (c *Conn) CalculateFees() (*big.Int, *big.Int, error) {
	// Set max fee per gas.
	feePerGas, err := c.CurrentBaseFee(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to obtain current base fee")
	}

	if viper.GetString("max-fee-per-gas") == "" {
		viper.Set("max-fee-per-gas", "200gwei")
	}
	maxFeePerGas, err := string2eth.StringToWei(viper.GetString("max-fee-per-gas"))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to obtain max fee per gas")
	}

	// Set priority fee per gas.
	priorityFeePerGas, err := string2eth.StringToWei(viper.GetString("priority-fee-per-gas"))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to obtain max priority fee per gas")
	}

	// Ensure that the total fee per gas does not exceed the max allowed.
	totalFeePerGas := new(big.Int).Add(feePerGas, priorityFeePerGas)
	if totalFeePerGas.Cmp(maxFeePerGas) > 0 {
		return nil, nil, fmt.Errorf("base fee %s plus priority fee %s (total %s) is higher than specified maximum (%s); increase with --max-fee-per-gas if you are sure you want to do this", string2eth.WeiToGWeiString(feePerGas), string2eth.WeiToGWeiString(priorityFeePerGas), string2eth.WeiToGWeiString(totalFeePerGas), string2eth.WeiToGWeiString(maxFeePerGas))
	}

	// Try to double the base fee, but not exceed the max allowed.
	feePerGas = new(big.Int).Mul(feePerGas, big.NewInt(2))
	totalFeePerGas = totalFeePerGas.Add(feePerGas, priorityFeePerGas)
	if totalFeePerGas.Cmp(maxFeePerGas) >= 0 {
		feePerGas = feePerGas.Sub(maxFeePerGas, priorityFeePerGas)
	}

	// Priority fee per gas cannot be higher than fee per gas.
	if priorityFeePerGas.Cmp(feePerGas) > 0 {
		// In this situation add the priority fee to the base fee.
		feePerGas = feePerGas.Add(feePerGas, priorityFeePerGas)
	}

	return feePerGas, priorityFeePerGas, nil
}
