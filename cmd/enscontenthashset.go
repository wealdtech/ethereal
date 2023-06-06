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

var ensContenthashSetContentStr string

// ensContenthashSetCmd represents the ens content hash set command.
var ensContenthashSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the content hash of an ENS domain",
	Long: `Set the content hash of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens contenthash set --domain=enstest.eth --content=/swarm/d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162 --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Fetch the owner of the name.
		owner, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "Cannot obtain owner")
		cli.Assert(!bytes.Equal(owner.Bytes(), ens.UnknownAddress.Bytes()), quiet, fmt.Sprintf("owner of %s is not set", ensDomain))

		cli.Assert(ensContenthashSetContentStr != "", quiet, "--content is required")
		data, err := ens.StringToContenthash(ensContenthashSetContentStr)
		cli.ErrCheck(err, quiet, "Unknown content")
		outputIf(verbose, fmt.Sprintf("Content hash is 0x%x", data))

		// Obtain the resolver for this name.
		resolver, err := ens.NewResolver(c.Client(), ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")

		signedTx, err := resolver.SetContenthash(opts, data)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":       "ens/contenthash",
			"command":     "set",
			"ensdomain":   ensDomain,
			"contenthash": ensContenthashSetContentStr,
		}, true)
	},
}

func init() {
	ensContenthashCmd.AddCommand(ensContenthashSetCmd)
	ensContenthashFlags(ensContenthashSetCmd)
	ensContenthashSetCmd.Flags().StringVar(&ensContenthashSetContentStr, "content", "", "The address to set e.g. /ipfs/QmdTEBPdNxJFFsH1wRE3YeWHREWDiSex8xhgTnqknyxWgu")
	addTransactionFlags(ensContenthashSetCmd, "passphrase for the account that owns the domain")
}
