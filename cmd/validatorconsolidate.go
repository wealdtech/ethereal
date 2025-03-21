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
	validatorconsolidate "github.com/wealdtech/ethereal/v2/cmd/validator/consolidate"
)

// validatorConsolidateCmd represents the contract call command.
var validatorConsolidateCmd = &cobra.Command{
	Use:   "consolidate",
	Short: "Consolidate two validators",
	Long: `Consolidate two consensus validators.  For example:

   ethereal validator consolidate --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --validator=0xa6372fbdec7dc4f14195e8aa2a6e6042264f1453073420ad8c5192423c4e4567af0ecef87a5cbdb8e9f574de8d312aa1 --withdrawal-amount=1eth

In quiet mode this will return 0 if the consolidate transaction is accepted, otherwise 1.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		res, err := validatorconsolidate.Run(cmd)
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
	validatorCmd.AddCommand(validatorConsolidateCmd)
	validatorFlags(validatorConsolidateCmd)
	validatorConsolidateCmd.Flags().String("from", "", "Address from which to send the consolidation request")
	validatorConsolidateCmd.Flags().String("source-validator", "", "Public key of the consensus validator")
	validatorConsolidateCmd.Flags().String("target-validator", "", "Public key of the consensus validator")
	validatorConsolidateCmd.Flags().String("max-fee", "1gwei", "Maximum fee to pay to consolidate the validatorx (excluding gas)")
	validatorConsolidateCmd.Flags().Bool("no-safety-checks", false, "Do not carry out safety checks (warning: could lose Ether)")
	addTransactionFlags(validatorConsolidateCmd, "the withdrawal address of the validator")
}

func validatorConsolidateBindings(cmd *cobra.Command) {
	validatorBindings(cmd)
	if err := viper.BindPFlag("from", cmd.Flags().Lookup("from")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("source-validator", cmd.Flags().Lookup("source-validator")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("target-validator", cmd.Flags().Lookup("target-validator")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("max-fee", cmd.Flags().Lookup("max-fee")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("no-safety-checks", cmd.Flags().Lookup("no-safety-checks")); err != nil {
		panic(err)
	}
}
