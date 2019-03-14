// Copyright © 2017 Weald Technology Trading
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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var ensNameClearAddress string

// ensNameClearCmd represents the ens name set command
var ensNameClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the ENS name of an address",
	Long: `Clear the Ethereum Name Service (ENS) name for an address.  For example:

    ethereal ens name clear --address=0x1234...5678 --passphrase="my secret passphrase"

The keystore for the address must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to clear the name is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(ensNameClearDomain != "", quiet, "--domain is required")

		// Obtain the reverse registrar
		registrar, err := ens.ReverseRegistrarContract(client)
		cli.ErrCheck(err, quiet, "Failed to obtain reverse registrar")

		opts, err := generateTxOpts(address)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err := registrar.SetName(opts, "")
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		setupLogging()
		log.WithFields(log.Fields{
			"group":         "ens/name",
			"command":       "clear",
			"address":       address.Hex(),
			"networkid":     chainID,
			"gas":           signedTx.Gas(),
			"gasprice":      signedTx.GasPrice().String(),
			"transactionid": signedTx.Hash().Hex(),
		}).Info("success")

		if quiet {
			os.Exit(0)
		}

		fmt.Println(signedTx.Hash().Hex())
	},
}

func init() {
	ensNameCmd.AddCommand(ensNameClearCmd)
	ensNameFlags(ensNameClearCmd)
	ensNameClearCmd.Flags().StringVar(&ensNameClearAddress, "address", "", "Address for which to clear reverse resolution")
	addTransactionFlags(ensNameClearCmd, "passphrase for the account that owns the address")
}
