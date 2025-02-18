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
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	validatorWithdrawFromAddress string
	validatorWithdrawValidator   string
	validatorWithdrawAmount      string
	validatorWithdrawMaxFee      string
)

// validatorWithdrawCmd represents the contract call command.
var validatorWithdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "Withdraw funds from a validator",
	Long: `Withdraw funds from a consensus validator.  For example:

   ethereal validator withdraw --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --validator=0xa6372fbdec7dc4f14195e8aa2a6e6042264f1453073420ad8c5192423c4e4567af0ecef87a5cbdb8e9f574de8d312aa1 --withdrawal-amount=1eth

In quiet mode this will return 0 if the withdrawal transaction is accepted, otherwise 1.`,
	Run: func(_ *cobra.Command, _ []string) {
		ctx := context.Background()

		cli.Assert(!offline, quiet, "This command needs access to chain data, so cannot run offline")

		fromAddress, err := c.Address(validatorWithdrawFromAddress, viper.GetString("privatekey"))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain from address"))

		cli.Assert(validatorWithdrawValidator != "", quiet, "validator cannot be empty")
		pubkey, err := c.ConsensusPubkey(validatorWithdrawValidator)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain validator public key %s", validatorWithdrawValidator))

		cli.Assert(validatorWithdrawAmount != "", quiet, "withdrawal amount cannot be empty")
		amount, err := string2eth.StringToWei(validatorWithdrawAmount)
		cli.ErrCheck(err, quiet, "Invalid amount")

		cli.Assert(validatorWithdrawMaxFee != "", quiet, "max fee amount cannot be empty")
		maxFee, err := string2eth.StringToWei(validatorWithdrawMaxFee)
		cli.ErrCheck(err, quiet, "Invalid max fee")
		cli.Assert(maxFee.Sign() == 1, quiet, "Max fee must be a positive value")

		signedTx, err := generateWithdrawalRequest(ctx, fromAddress, pubkey, amount, maxFee)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			err = c.SendTransaction(ctx, signedTx)
			cli.ErrCheck(err, quiet, "Failed to initiate validator withdrawal")
			handleSubmittedTransaction(signedTx, log.Fields{
				"group":   "validator",
				"command": "withdraw",
			}, false)
		}
	},
}

func init() {
	validatorCmd.AddCommand(validatorWithdrawCmd)
	validatorFlags(validatorWithdrawCmd)
	validatorWithdrawCmd.Flags().StringVar(&validatorWithdrawFromAddress, "from", "", "Address from which to send the withdraw request")
	validatorWithdrawCmd.Flags().StringVar(&validatorWithdrawValidator, "validator", "", "Public key of the consensus validator")
	validatorWithdrawCmd.Flags().StringVar(&validatorWithdrawAmount, "withdrawal-amount", "", "Amount to withdraw from the validator")
	validatorWithdrawCmd.Flags().StringVar(&validatorWithdrawMaxFee, "max-fee", "1gwei", "Maximum fee to pay to withdraw funds from the validator (excluding gas)")
	addTransactionFlags(validatorWithdrawCmd, "the withdrawal address of the validator")
}
