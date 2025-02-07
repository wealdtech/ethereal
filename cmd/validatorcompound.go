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
	"github.com/wealdtech/ethereal/v2/cli"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	validatorCompoundFromAddress string
	validatorCompoundValidator   string
	validatorCompoundMaxFee      string
)

// validatorCompoundCmd represents the contract call command.
var validatorCompoundCmd = &cobra.Command{
	Use:   "compound",
	Short: "Compound a validator",
	Long: `Compound a consensus validators, allowing it to have a balance greater than 32 ETH.  For example:

   ethereal validator compound --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --validator=0xa6372fbdec7dc4f14195e8aa2a6e6042264f1453073420ad8c5192423c4e4567af0ecef87a5cbdb8e9f574de8d312aa1

In quiet mode this will return 0 if the compound transaction is accepted, otherwise 1.`,
	Run: func(_ *cobra.Command, _ []string) {
		ctx := context.Background()

		cli.Assert(!offline, quiet, "This command needs access to chain data, so cannot run offline")

		cli.Assert(validatorCompoundFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(validatorCompoundFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", validatorCompoundFromAddress))

		cli.Assert(validatorCompoundValidator != "", quiet, "validator cannot be empty")
		pubkey, err := c.ConsensusPubkey(validatorCompoundValidator)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain validator public key %s", validatorCompoundValidator))

		cli.Assert(validatorCompoundMaxFee != "", quiet, "max fee amount cannot be empty")
		maxFee, err := string2eth.StringToWei(validatorCompoundMaxFee)
		cli.ErrCheck(err, quiet, "Invalid max fee")
		cli.Assert(maxFee.Sign() == 1, quiet, "Max fee must be a positive value")

		// A compound is a consolidation with source and target pubkeys the same.
		signedTx, err := generateConsolidationRequest(ctx, fromAddress, pubkey, pubkey, maxFee)
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
	validatorCmd.AddCommand(validatorCompoundCmd)
	validatorFlags(validatorCompoundCmd)
	validatorCompoundCmd.Flags().StringVar(&validatorCompoundFromAddress, "from", "", "Address from which to send the compound request")
	validatorCompoundCmd.Flags().StringVar(&validatorCompoundValidator, "validator", "", "Public key of the validator to change to compounding")
	validatorCompoundCmd.Flags().StringVar(&validatorCompoundMaxFee, "max-fee", "1gwei", "Maximum fee to pay to change the validator to compounding (excluding gas)")
	addTransactionFlags(validatorCompoundCmd, "the withdrawal address of the validator")
}
