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
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// PrepareTx prepares a transaction, setting the gas limit and access list if applicable.
// Consider using CreateTransaction instead of this call, for a more comprehensive process.
func (c *Conn) PrepareTx(ctx context.Context,
	txData *TransactionData,
) error {
	if c.client == nil {
		// We're offline; fetch gas limit from input.
		gasLimit := viper.GetInt64("gaslimit")
		if gasLimit <= 0 {
			return errors.New("gas limit not specified")
		}
		return nil
	}

	// Obtain the gas limit from estimateGas.
	basicGas, err := c.EstimateGas(ctx, txData)
	if err != nil {
		return err
	}

	// Obtain the gas limit and access list from createTransactionList.
	accessList, accessListGas, err := c.CreateAccessList(ctx, txData)
	if err != nil {
		return err
	}

	if basicGas < accessListGas {
		if c.debug {
			fmt.Fprintf(os.Stderr, "Not using access list (%d vs %d)\n", accessListGas, basicGas)
		}
		// Add 25% overhead.
		basicGas *= 5 / 4
		txData.GasLimit = &basicGas
	} else {
		if c.debug {
			fmt.Fprintf(os.Stderr, "Using access list (%d vs %d)\n", accessListGas, basicGas)
		}
		// Add 25% overhead and access list.
		accessListGas *= 5 / 4
		txData.GasLimit = &accessListGas
		txData.AccessList = *accessList
	}

	return nil
}

// EstimateGas estimates the gas required for the given transaction.
// Consider using CreateTransaction instead of this call, for a more comprehensive process.
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

	msg := ethereum.CallMsg{
		From:       txData.From,
		To:         txData.To,
		Value:      txData.Value,
		Data:       txData.Data,
		AccessList: txData.AccessList,
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	gas, err := c.client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, errors.Wrap(err, "failed to estimate gas")
	}
	return gas, err
}

// CreateAccessList creates an access list for the given transaction.
// Consider using CreateTransaction instead of this call, for a more comprehensive process.
func (c *Conn) CreateAccessList(ctx context.Context,
	txData *TransactionData,
) (
	*types.AccessList,
	uint64,
	error,
) {
	if c.client == nil {
		// We're offline; no access list.  Return with max gas used to hint that this isn't a good result.
		return &types.AccessList{}, 0xffffffff, nil
	}

	msg := ethereum.CallMsg{
		From:  txData.From,
		To:    txData.To,
		Value: txData.Value,
		Data:  txData.Data,
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	accessList, gasUsed, errString, err := c.gethClient.CreateAccessList(ctx, msg)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "failed to create access list: %s", errString)
	}

	return accessList, gasUsed, nil
}
