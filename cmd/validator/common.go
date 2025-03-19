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

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/api"
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

var (
	feeBumpDivisor = big.NewInt(2)

	// minFeeBump is the minimum fee bump, in wei.
	minFeeBump = big.NewInt(1e6)

	// maxFeeBump is the maximum fee bump, in wei.
	maxFeeBump = big.NewInt(1e12)
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

	fee = bumpFee(fee)

	return fee, nil
}

// bumpFee bumps up the fee provided by the system contract to allow for
func bumpFee(initial *big.Int) *big.Int {
	// Add a 50% fee bump.
	feeBump := new(big.Int).Div(initial, feeBumpDivisor)
	if feeBump.Cmp(minFeeBump) <= 0 {
		// Fee bump is very small, increase to minimum.
		feeBump = minFeeBump
	}
	if feeBump.Cmp(maxFeeBump) >= 0 {
		// Fee bump is very large, decrease to maximum.
		feeBump = maxFeeBump
	}
	return new(big.Int).Add(initial, feeBump)
}

func obtainSystemDepositContractAddress(ctx context.Context,
	consensusClient consensusclient.Service,
) (
	common.Address,
	error,
) {
	specProvider, isProvider := consensusClient.(consensusclient.SpecProvider)
	if !isProvider {
		return common.Address{}, errors.New("consensus client is not a spec provider")
	}

	specResponse, err := specProvider.Spec(ctx, &api.SpecOpts{})
	if err != nil {
		return common.Address{}, errors.Join(errors.New("failed to obtain spec"), err)
	}
	spec := specResponse.Data

	tmp := spec["DEPOSIT_CONTRACT_ADDRESS"]
	addr, isSlice := tmp.([]byte)
	if !isSlice {
		return common.Address{}, errors.New("DEPOSIT_CONTRACT_ADDRESS not of expected type")
	}

	return common.BytesToAddress(addr), nil
}

func GenerateTopupRequest(ctx context.Context,
	executionConn *conn.Conn,
	consensusClient consensusclient.Service,
	fromAddress common.Address,
	pubkey phase0.BLSPubKey,
	amount *big.Int,
	debug bool,
) (
	*types.Transaction,
	error,
) {
	toAddress, err := obtainSystemDepositContractAddress(ctx, consensusClient)
	if err != nil {
		return nil, err
	}
	if debug {
		fmt.Fprintf(os.Stderr, "Deposit contract address is %#x\n", toAddress[:])
	}

	// Withdrawal credentials are ignored for topups, leave as 0.
	withdrawalCredentials := make([]byte, 32)
	// Signature is ignored for topups, set to infinity.
	signature := phase0.BLSSignature{0xc0}

	gweiAmount := new(big.Int).Div(amount, big.NewInt(1e9)).Uint64()
	depositData := &phase0.DepositData{
		PublicKey: pubkey,
		// WithdrawalCredentials: make([]byte, 32),
		WithdrawalCredentials: withdrawalCredentials,
		Amount:                phase0.Gwei(gweiAmount),
		// Signature:             phase0.BLSSignature{0xc0},
		Signature: signature,
	}

	depositDataRoot, err := depositData.HashTreeRoot()
	if err != nil {
		return nil, errors.Join(errors.New("failed to generate deposit data root"), err)
	}

	data := make([]byte, 0)
	// Add the function identifier:
	data = append(data, []byte{0x22, 0x89, 0x51, 0x18}...)
	// Add the parameter offsets.
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}...)
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0}...)
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x20}...)
	// Deposit data root.
	data = append(data, depositDataRoot[:]...)
	// Public key.
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30}...)
	data = append(data, pubkey[:]...)
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}...)
	// Withdrawal credentials.
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20}...)
	data = append(data, withdrawalCredentials...)
	// Signature.
	data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60}...)
	data = append(data, signature[:]...)
	if debug {
		fmt.Fprintf(os.Stderr, "Transaction data is %#x\n", data)
	}

	txData := &conn.TransactionData{
		From:  fromAddress,
		To:    &toAddress,
		Value: amount,
		Data:  data,
	}

	// Create and sign the transaction.
	signedTx, err := executionConn.CreateSignedTransaction(ctx, txData)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
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
