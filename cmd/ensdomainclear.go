// Copyright © 2017-2019 Weald Technology Trading
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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensDomainClearAddress string

// ensDomainClearCmd represents the ens domain clear command.
var ensDomainClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the ENS domain of an address",
	Long: `Clear the Ethereum Name Service (ENS) domain for an address.  For example:

    ethereal ens domain clear --address=0x1234...5678 --passphrase="my secret passphrase"

The keystore for the address must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(ensDomainClearAddress != "", quiet, "--address is required")
		address, err := c.Resolve(ensDomainClearAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain address to clear domain")

		// Obtain the reverse registrar.
		registrar, err := ens.NewReverseRegistrar(c.Client())
		cli.ErrCheck(err, quiet, "Failed to obtain reverse registrar")

		opts, err := generateTxOpts(address)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err := registrar.SetName(opts, "")
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":      "ens/domain",
			"command":    "clear",
			"ensaddress": address.Hex(),
		}, true)
	},
}

func init() {
	ensDomainCmd.AddCommand(ensDomainClearCmd)
	ensDomainFlags(ensDomainClearCmd)
	ensDomainClearCmd.Flags().StringVar(&ensDomainClearAddress, "address", "", "Address for which to clear reverse resolution")
	addTransactionFlags(ensDomainClearCmd, "passphrase for the account that owns the address")
}
