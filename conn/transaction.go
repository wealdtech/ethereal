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

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// CreateSignedTransaction creates a signed transaction.
func (c *Conn) CreateSignedTransaction(ctx context.Context,
	txData *TransactionData,
) (
	*types.Transaction,
	error,
) {
	tx, err := c.CreateTransaction(ctx, txData)
	if err != nil {
		return nil, err
	}

	// Sign the transaction.
	signedTx, err := c.SignTransaction(ctx, txData.From, tx)
	if err != nil {
		err = fmt.Errorf("failed to sign transaction: %v", err)
		return nil, err
	}

	// Increment the nonce for the next transaction.
	_, err = c.NextNonce(ctx, txData.From)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// CreateTransaction creates a transaction.
func (c *Conn) CreateTransaction(ctx context.Context,
	txData *TransactionData,
) (
	*types.Transaction,
	error,
) {
	if txData.Nonce == nil {
		// Obtain the nonce for the transaction.
		nonce, err := c.CurrentNonce(ctx, txData.From)
		if err != nil {
			return nil, err
		}
		txNonce := int64(nonce)
		txData.Nonce = &txNonce
	}

	if txData.GasLimit == nil {
		// Calculate gas limit for the transaction
		gasLimit, err := c.EstimateGas(ctx, txData)
		if err != nil {
			return nil, err
		}
		txData.GasLimit = &gasLimit
	}

	// Calculate fees.
	maxFeePerGas, maxPriorityFeePerGas, err := c.CalculateFees()
	if err != nil {
		return nil, err
	}
	if txData.MaxFeePerGas != nil {
		// Manual override for fee per gas.
		maxFeePerGas = txData.MaxFeePerGas
	}
	if txData.MaxPriorityFeePerGas != nil {
		// Manual override for priority fee per gas.
		maxPriorityFeePerGas = txData.MaxPriorityFeePerGas
	}

	// Create the transaction
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   c.ChainID(),
		Nonce:     uint64(*txData.Nonce),
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       *txData.GasLimit,
		To:        txData.To,
		Value:     txData.Value,
		Data:      txData.Data,
	}), nil
}

// SendTransaction send the supplied transaction to the network.
func (c *Conn) SendTransaction(ctx context.Context,
	tx *types.Transaction,
) error {
	if c.client == nil {
		return errors.New("cannot send transaction when offline")
	}
	if tx == nil {
		return errors.New("transaction is nil")
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	if err := c.Client().SendTransaction(ctx, tx); err != nil {
		return errors.Wrap(err, "failed to send transaction")
	}

	return nil
}
