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

package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/wealdtech/ethereal/v2/conn"
	"github.com/wealdtech/go-string2eth"
)

const (
	// systemConsolidationContractAddress is defined in https://eips.ethereum.org/EIPS/eip-7251
	systemConsolidationContractAddress = "0x0000BBdDc7CE488642fb579F8B00f3a590007251"
	// systemWithdrawalContractAddress is defined in https://eips.ethereum.org/EIPS/eip-7002
	systemWithdrawalContractAddress = "0x00000961Ef480Eb55e80D19ad83579A64c007002"
)

func getValidatorSystemContractFee(ctx context.Context, c *conn.Conn, address common.Address) (*big.Int, error) {
	msg := ethereum.CallMsg{
		To: &address,
	}
	result, err := c.Client().CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	fee := new(big.Int).SetBytes(result)

	return fee, nil
}

func GenerateWithdrawalRequest(ctx context.Context,
	executionConn *conn.Conn,
	fromAddress common.Address,
	pubkey phase0.BLSPubKey,
	amount *big.Int,
	maxFee *big.Int,
	debug bool,
) (
	*types.Transaction,
	error,
) {
	toAddress := common.HexToAddress(systemWithdrawalContractAddress)

	data := make([]byte, 0)

	data = append(data, pubkey[:]...)

	// The transaction requires the withdrawal amount in gwei, not wei.
	amount = new(big.Int).Div(amount, big.NewInt(1e9))
	if amount.Sign() <= 0 {
		return nil, errors.New("withdrawal amount must be at least 1gwei")
	}

	// Amount must be 8 bytes.
	amountBytes := make([]byte, 8)
	encodedAmount := amount.Bytes()
	copy(amountBytes[8-len(encodedAmount):], encodedAmount)
	data = append(data, amountBytes...)

	// Obtain the current fee.
	withdrawalFee, err := getValidatorSystemContractFee(ctx, executionConn, toAddress)
	if err != nil {
		return nil, errors.Join(errors.New("failed to obtain current withdrawal fee"), err)
	}
	if debug {
		fmt.Fprintf(os.Stderr, "Withdrawal fee is %s\n", string2eth.WeiToString(withdrawalFee, true))
	}
	if withdrawalFee.Cmp(maxFee) > 0 {
		return nil, fmt.Errorf("fee for this operation is %s, which is higher than maximum; increase with max-fee if you want to proceed", string2eth.WeiToString(withdrawalFee, true))
	}

	txData := &conn.TransactionData{
		From:  fromAddress,
		To:    &toAddress,
		Value: withdrawalFee,
		Data:  data,
	}

	// Create and sign the transaction.
	signedTx, err := executionConn.CreateSignedTransaction(ctx, txData)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func GenerateConsolidationRequest(ctx context.Context,
	executionConn *conn.Conn,
	fromAddress common.Address,
	sourcePubkey phase0.BLSPubKey,
	targetPubkey phase0.BLSPubKey,
	maxFee *big.Int,
	debug bool,
) (
	*types.Transaction,
	error,
) {
	toAddress := common.HexToAddress(systemConsolidationContractAddress)

	data := make([]byte, 0)
	data = append(data, sourcePubkey[:]...)
	data = append(data, targetPubkey[:]...)

	// Obtain the current fee.
	consolidationFee, err := getValidatorSystemContractFee(ctx, executionConn, toAddress)
	if err != nil {
		return nil, errors.Join(errors.New("failed to obtain current consolidation fee"), err)
	}
	if debug {
		fmt.Fprintf(os.Stderr, "Consolidation fee is %s\n", string2eth.WeiToString(consolidationFee, true))
	}
	if consolidationFee.Cmp(maxFee) > 0 {
		return nil, fmt.Errorf("fee for this operation is %s, which is higher than maximum; increase with max-fee if you want to proceed", string2eth.WeiToString(consolidationFee, true))
	}

	txData := &conn.TransactionData{
		From:  fromAddress,
		To:    &toAddress,
		Value: consolidationFee,
		Data:  data,
	}

	// Create and sign the transaction.
	signedTx, err := executionConn.CreateSignedTransaction(ctx, txData)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
