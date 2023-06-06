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
	"encoding/hex"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensPubkeySetKey string

// ensPubkeySetCmd represents the ens pubkey set command.
var ensPubkeySetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the public key of an ENS domain",
	Long: `Set the public key of a name registered with the Ethereum Name Service (ENS) for a given name.  For example:

    ethereal ens pubkey set --domain=enstest.eth --key="(0x0102...1e1f,0x0102...1e1f)" --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		cli.Assert(ensPubkeySetKey != "", quiet, "--key is required")
		key := strings.Split(strings.ToLower(ensPubkeySetKey), ",")
		cli.Assert(len(key) == 2, quiet, "Key should be in (x,y) format")
		key[0] = strings.TrimSpace(key[0])
		key[0] = strings.TrimPrefix(key[0], "(")
		key[0] = strings.TrimPrefix(key[0], "0x")
		x := [32]byte{}
		val, err := hex.DecodeString(key[0])
		cli.ErrCheck(err, quiet, "Invalid x")
		copy(x[32-len(val):], val)
		outputIf(debug, fmt.Sprintf("x is %x", x))

		key[1] = strings.TrimSpace(key[1])
		key[1] = strings.TrimPrefix(key[1], "0x")
		key[1] = strings.TrimSuffix(key[1], ")")
		y := [32]byte{}
		val, err = hex.DecodeString(key[1])
		cli.ErrCheck(err, quiet, "Invalid y")
		copy(y[32-len(val):], val)
		outputIf(debug, fmt.Sprintf("y is %x", y))

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Fetch the owner of the name.
		owner, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "Cannot obtain owner")
		cli.Assert(!bytes.Equal(owner.Bytes(), ens.UnknownAddress.Bytes()), quiet, fmt.Sprintf("owner of %s is not set", ensDomain))

		// Obtain the resolver for this name.
		resolver, err := ens.NewResolver(c.Client(), ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")

		signedTx, err := resolver.SetPubKey(opts, x, y)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":     "ens/pubkey",
			"command":   "set",
			"ensdomain": ensDomain,
			"x":         x,
			"y":         y,
		}, true)
	},
}

func init() {
	ensPubkeyCmd.AddCommand(ensPubkeySetCmd)
	ensPubkeyFlags(ensPubkeySetCmd)
	ensPubkeySetCmd.Flags().StringVar(&ensPubkeySetKey, "key", "", "The key to set in (x,y) format")
	addTransactionFlags(ensPubkeySetCmd, "passphrase for the account that owns the domain")
}
