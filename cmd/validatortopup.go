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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	validatortopup "github.com/wealdtech/ethereal/v2/cmd/validator/topup"
)

// validatorTopupCmd represents the validator topup command.
var validatorTopupCmd = &cobra.Command{
	Use:   "topup",
	Short: "Topup funds from a validator",
	Long: `Topup funds from a consensus validator.  For example:

   ethereal validator topup --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --validator=0xa6372fbdec7dc4f14195e8aa2a6e6042264f1453073420ad8c5192423c4e4567af0ecef87a5cbdb8e9f574de8d312aa1 --topup-amount=1eth

In quiet mode this will return 0 if the topup transaction is accepted, otherwise 1.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		res, err := validatortopup.Run(cmd)
		if err != nil {
			return err
		}
		if viper.GetBool("quiet") {
			return nil
		}
		fmt.Fprint(os.Stdout, res)

		return nil
	},
}

func init() {
	validatorCmd.AddCommand(validatorTopupCmd)
	validatorFlags(validatorTopupCmd)
	validatorTopupCmd.Flags().String("from", "", "Address from which to send the topup request")
	validatorTopupCmd.Flags().String("validator", "", "Public key of the consensus validator")
	validatorTopupCmd.Flags().String("topup-amount", "", "Amount to topup the validator")
	validatorTopupCmd.Flags().Bool("no-safety-checks", false, "Do not carry out safety checks (warning: could lose Ether)")
	addTransactionFlags(validatorTopupCmd, "the withdrawal address of the validator")
}

func validatorTopupBindings(cmd *cobra.Command) {
	validatorBindings(cmd)
	if err := viper.BindPFlag("from", cmd.Flags().Lookup("from")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("validator", cmd.Flags().Lookup("validator")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("topup-amount", cmd.Flags().Lookup("topup-amount")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("no-safety-checks", cmd.Flags().Lookup("no-safety-checks")); err != nil {
		panic(err)
	}
}
