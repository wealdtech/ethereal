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
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	validatorExitFromAddress string
	validatorExitValidator   string
	validatorExitMaxFee      string
)

// validatorExitCmd represents the contract call command.
var validatorExitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Exit from a validator",
	Long: `Exit from a consensus validator.  For example:

   ethereal validator exit --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --validator=0xa6372fbdec7dc4f14195e8aa2a6e6042264f1453073420ad8c5192423c4e4567af0ecef87a5cbdb8e9f574de8d312aa1

In quiet mode this will return 0 if the exit transaction is accepted, otherwise 1.`,
	Run: func(_ *cobra.Command, _ []string) {
		ctx := context.Background()

		cli.Assert(!offline, quiet, "This command needs access to chain data, so cannot run offline")

		cli.Assert(validatorExitFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(validatorExitFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", validatorExitFromAddress))

		cli.Assert(validatorExitValidator != "", quiet, "validator cannot be empty")
		pubkey, err := c.ConsensusPubkey(validatorExitValidator)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain validator public key %s", validatorExitValidator))

		cli.Assert(validatorExitMaxFee != "", quiet, "max fee amount cannot be empty")
		maxFee, err := string2eth.StringToWei(validatorExitMaxFee)
		cli.ErrCheck(err, quiet, "Invalid max fee")
		cli.Assert(maxFee.Sign() == 1, quiet, "Max fee must be a positive value")

		// An exit is a withdrawal request with the amount set to 0.
		signedTx, err := generateWithdrawalRequest(ctx, fromAddress, pubkey, big.NewInt(0), maxFee)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			err = c.SendTransaction(ctx, signedTx)
			cli.ErrCheck(err, quiet, "Failed to initiate validator exit")
			handleSubmittedTransaction(signedTx, log.Fields{
				"group":   "validator",
				"command": "exit",
			}, false)
		}
	},
}

func init() {
	validatorCmd.AddCommand(validatorExitCmd)
	validatorFlags(validatorExitCmd)
	validatorExitCmd.Flags().StringVar(&validatorExitFromAddress, "from", "", "Address from which to send the exit request")
	validatorExitCmd.Flags().StringVar(&validatorExitValidator, "validator", "", "Public key of the consensus validator")
	validatorExitCmd.Flags().StringVar(&validatorExitMaxFee, "max-fee", "1gwei", "Maximum fee to pay to exit the validator (excluding gas)")
	addTransactionFlags(validatorExitCmd, "the withdrawal address of the validator")
}
