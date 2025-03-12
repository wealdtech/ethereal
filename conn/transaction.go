// Copyright © 2022, 2025 Weald Technology Trading
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
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/util"
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
	signedTx, err := c.SignTransaction(ctx, txData.From, tx, c.debug)
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
		// Set the gas limit (and perhaps the access list).
		if err := c.PrepareTx(ctx, txData); err != nil {
			return nil, err
		}
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
		ChainID:    c.ChainID(),
		Nonce:      uint64(*txData.Nonce),
		GasFeeCap:  maxFeePerGas,
		GasTipCap:  maxPriorityFeePerGas,
		Gas:        *txData.GasLimit,
		To:         txData.To,
		Value:      txData.Value,
		Data:       txData.Data,
		AccessList: txData.AccessList,
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

// HandleSubmittedTransaction handles logging and waiting for a submitted transaction to be mined.
// It will not log the transaction if logFields is nil.
// This function will return false if asked to wait and the transaction is not mined, otherwise true.
func (c *Conn) HandleSubmittedTransaction(tx *types.Transaction, logFields log.Fields) bool {
	if logFields != nil {
		c.logTransaction(tx, logFields)
	}

	if !viper.GetBool("wait") {
		if !c.quiet {
			fmt.Fprintf(os.Stdout, "%s\n", tx.Hash().Hex())
		}

		return true
	}

	mined := util.WaitForTransaction(c.Client(), tx.Hash(), viper.GetDuration("limit"))
	if mined {
		fmt.Fprintf(os.Stdout, "%s mined\n", tx.Hash().Hex())

		return true
	}
	fmt.Fprintf(os.Stdout, "%s submitted but not mined\n", tx.Hash().Hex())

	return false
}

// logTransaction logs a transaction.
func (c *Conn) logTransaction(tx *types.Transaction, fields log.Fields) {
	txFields := log.Fields{
		"networkid":            c.ChainID(),
		"transactionid":        tx.Hash().Hex(),
		"gas":                  tx.Gas(),
		"fee-per-gas":          tx.GasFeeCap().String(),
		"priority-fee-per-gas": tx.GasTipCap().String(),
		"value":                tx.Value().String(),
		"data":                 hex.EncodeToString(tx.Data()),
	}
	if c.signer != nil {
		fromAddress, err := types.Sender(c.signer, tx)
		if err == nil {
			txFields["from"] = fromAddress.Hex()
		}
	}
	if tx.To() != nil {
		txFields["to"] = tx.To().Hex()
	}

	log.WithFields(fields).WithFields(txFields).Info("transaction submitted")
}
