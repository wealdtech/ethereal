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

package cmd

import (
	"context"
	"fmt"
	"math/big"
	"os"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	string2eth "github.com/wealdtech/go-string2eth"
)

// validatorCmd represents the validator command.
var validatorCmd = &cobra.Command{
	Use:     "validator",
	Aliases: []string{"beacon"},
	Short:   "Manage validators",
	Long:    `Manage consensus validators: consolidations, withdrawals etc.`,
}

func init() {
	RootCmd.AddCommand(validatorCmd)
}

func validatorFlags(_ *cobra.Command) {
}

func validatorBindings(_ *cobra.Command) {
}

// TODO move these.
const (
	// systemConsolidationContractAddress is defined in https://eips.ethereum.org/EIPS/eip-7251
	systemConsolidationContractAddress = "0x0000BBdDc7CE488642fb579F8B00f3a590007251"
	// systemWithdrawalContractAddress is defined in https://eips.ethereum.org/EIPS/eip-7002
	systemWithdrawalContractAddress = "0x00000961Ef480Eb55e80D19ad83579A64c007002"
)

func generateConsolidationRequest(ctx context.Context,
	fromAddress common.Address,
	sourcePubkey [48]byte,
	targetPubkey [48]byte,
	maxFee *big.Int,
) (
	*types.Transaction,
	error,
) {
	toAddress := common.HexToAddress(systemConsolidationContractAddress)

	data := make([]byte, 0)
	data = append(data, sourcePubkey[:]...)
	data = append(data, targetPubkey[:]...)

	// Obtain the current fee.
	withdrawalFee, err := getValidatorSystemContractFee(ctx, toAddress)
	cli.ErrCheck(err, quiet, "Failed to obtain consolidation fee")
	cli.Assert(withdrawalFee.Cmp(maxFee) <= 0, quiet, fmt.Sprintf("Fee for this operation is %s, which is higher than the maximum allowed of %s; increase maximum allowed with max-fee if you want to proceed", string2eth.WeiToString(withdrawalFee, true), string2eth.WeiToString(maxFee, true)))

	if debug {
		fmt.Fprintf(os.Stderr, "Data is %x\n", data)
	}

	txData := &conn.TransactionData{
		From:  fromAddress,
		To:    &toAddress,
		Value: withdrawalFee,
		Data:  data,
	}

	// Create and sign the transaction.
	signedTx, err := c.CreateSignedTransaction(ctx, txData)
	cli.ErrCheck(err, quiet, "Failed to create transaction")

	return signedTx, nil
}

func getValidatorSystemContractFee(ctx context.Context, address common.Address) (*big.Int, error) {
	msg := ethereum.CallMsg{
		To: &address,
	}
	result, err := c.Client().CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	fee := new(big.Int).SetBytes(result)

	if debug {
		fmt.Fprintf(os.Stderr, "Fee is %s\n", string2eth.WeiToString(fee, true))
	}

	return fee, nil
}

func generateWithdrawalRequest(ctx context.Context,
	fromAddress common.Address,
	pubkey [48]byte,
	amount *big.Int,
	maxFee *big.Int,
) (
	*types.Transaction,
	error,
) {
	toAddress := common.HexToAddress(systemWithdrawalContractAddress)

	data := make([]byte, 0)

	data = append(data, pubkey[:]...)

	// The transaction requires the withdrawal amount in gwei, not wei.
	amount = new(big.Int).Div(amount, big.NewInt(1e9))
	if amount.Sign() != 0 {
		cli.Assert(amount.Sign() > 0, quiet, "amount must be at least 1gwei")
	}

	// Amount must be 8 bytes.
	amountBytes := make([]byte, 8)
	encodedAmount := amount.Bytes()
	copy(amountBytes[8-len(encodedAmount):], encodedAmount)
	data = append(data, amountBytes...)

	// Obtain the current fee.
	withdrawalFee, err := getValidatorSystemContractFee(ctx, toAddress)
	cli.ErrCheck(err, quiet, "Failed to obtain withdrawal fee")
	cli.Assert(withdrawalFee.Cmp(maxFee) <= 0, quiet, fmt.Sprintf("Fee for this operation is %s, which is higher than maximum; increase with max-fee if you want to proceed", string2eth.WeiToString(withdrawalFee, true)))

	if debug {
		fmt.Fprintf(os.Stderr, "Data is %x\n", data)
	}

	txData := &conn.TransactionData{
		From:  fromAddress,
		To:    &toAddress,
		Value: withdrawalFee,
		Data:  data,
	}

	// Create and sign the transaction.
	signedTx, err := c.CreateSignedTransaction(ctx, txData)
	cli.ErrCheck(err, quiet, "Failed to create transaction")

	return signedTx, nil
}
