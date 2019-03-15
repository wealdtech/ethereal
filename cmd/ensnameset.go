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

var ensNameSetAddress string
var ensNameSetDomain string

// ensNameSetCmd represents the ens name set command
var ensNameSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the ENS name of an address",
	Long: `Set the Ethereum Name Service (ENS) name for an address.  For example:

    ethereal ens name set --domain=enstest.eth --address=0x1234...5678 --passphrase="my secret passphrase"

The keystore for the address name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to set the name is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(ensNameSetAddress != "", quiet, "--address is required")
		address, err := ens.Resolve(client, ensNameSetAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain address to set name; to clear the name use \"ens name clear\"")

		cli.Assert(ensNameSetDomain != "", quiet, "--domain is required")

		// Obtain the reverse registrar
		registrar, err := ens.ReverseRegistrarContract(client)
		cli.ErrCheck(err, quiet, "Failed to obtain reverse registrar")

		opts, err := generateTxOpts(address)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		ensNameSetDomain = ens.Normalize(ensNameSetDomain)
		signedTx, err := registrar.SetName(opts, ensNameSetDomain)
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		logTransaction(signedTx, log.Fields{
			"group":   "ens/name",
			"command": "set",
			"domain":  ensDomain,
			"address": address.Hex(),
		})

		if quiet {
			os.Exit(0)
		}

		fmt.Println(signedTx.Hash().Hex())
	},
}

func init() {
	ensNameCmd.AddCommand(ensNameSetCmd)
	ensNameFlags(ensNameSetCmd)
	ensNameSetCmd.Flags().StringVar(&ensNameSetAddress, "address", "", "Address for which to set reverse resolution")
	ensNameSetCmd.Flags().StringVar(&ensNameSetDomain, "domain", "", "The reverse resolution domain")
	addTransactionFlags(ensNameSetCmd, "passphrase for the account that owns the address")
}
