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
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensControllerSetControllerStr string

// ensControllerSetCmd represents the ens controller set command.
var ensControllerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the controller of an ENS domain",
	Long: `Set the controller of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens controller set --domain=enstest.eth --controller=0x1234...5678 --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Fetch the controller of the name.
		controller, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "Cannot obtain current controller")
		cli.Assert(!bytes.Equal(controller.Bytes(), ens.UnknownAddress.Bytes()), quiet, fmt.Sprintf("%s has no controller", ensDomain))

		cli.Assert(ensControllerSetControllerStr != "", quiet, "--controller is required")
		newControllerAddress, err := c.Resolve(ensControllerSetControllerStr)
		cli.Assert(!bytes.Equal(newControllerAddress.Bytes(), ens.UnknownAddress.Bytes()), quiet, "Attempt to set controller to 0x00 disallowed")
		cli.ErrCheck(err, quiet, "Failed to obtain new controller address")

		opts, err := generateTxOpts(controller)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err := registry.SetOwner(opts, ensDomain, newControllerAddress)
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":         "ens/controller",
			"command":       "set",
			"ensdomain":     ensDomain,
			"enscontroller": newControllerAddress.Hex(),
		}, false)
	},
}

func init() {
	ensControllerCmd.AddCommand(ensControllerSetCmd)
	ensControllerFlags(ensControllerSetCmd)
	ensControllerSetCmd.Flags().StringVar(&ensControllerSetControllerStr, "controller", "", "The new controller's name or address")
	addTransactionFlags(ensControllerSetCmd, "passphrase for the account that owns the domain")
}
