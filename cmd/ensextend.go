// Copyright © 2019 Weald Technology Trading
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
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
	string2eth "github.com/wealdtech/go-string2eth"
)

var ensExtendDomains string

// ensExtendCmd represents the extend command.
var ensExtendCmd = &cobra.Command{
	Use:   "extend",
	Short: "Extend the registration of an ENS domain",
	Long: `Extend the registration of an Ethereum Name Service (ENS) domain.  For example:

    ethereal ens extend --domain=enstest.eth --value="0.1 Ether" --passphrase="my secret passphrase"

Multiple domains can be extended with a single command (note this still creates one transaction for each name).  For example:

    ethereal ens extend --domains="mydomain1.eth&&mydomain2.eth" --value="0.1 Ether" --passphrase="my secret passphrase"

In the latter case the value will be used for each domain, rather than spread between the domains to be extended.

The keystore for the domain(s) owner must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transactions are successfully submitted (and mined if --wait is supplied), 1 if the transactions are not successfully submitted, and 2 if the transactions are successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "" || ensExtendDomains != "", quiet, "--domain or --domains is required")

		var domains []string
		if ensExtendDomains != "" {
			domains = strings.Split(ensExtendDomains, "&&")
		} else {
			domains = make([]string, 1)
			domains[0] = ensDomain
		}

		cli.Assert(viper.GetString("value") != "", quiet, "--value is required")

		value, err := string2eth.StringToWei(viper.GetString("value"))
		cli.ErrCheck(err, quiet, "Could not understand value")
		// Extend loop.
		var lastTx *types.Transaction
		for _, domain := range domains {
			domain, err = ens.NormaliseDomain(domain)
			cli.ErrCheck(err, quiet, "Failed to normalise ENS domain")

			// Ensure the domain is owned.
			registry, err := ens.NewRegistry(c.Client())
			cli.ErrCheck(err, quiet, "Failed to obtain ENS registry")
			owner, err := registry.Owner(domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain owner for %s", domain))
			cli.Assert(owner != ens.UnknownAddress, quiet, fmt.Sprintf("%s is not registered", domain))

			controller, err := ens.NewETHController(c.Client(), ens.Domain(domain))
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain %s controller", ens.Domain(domain)))

			registrar, err := ens.NewBaseRegistrar(c.Client(), ens.Domain(domain))
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain %s registrar", ens.Domain(domain)))

			// Obtain current expiry.
			expiryTS, err := registrar.Expiry(domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain expiry for %s", domain))
			expiry := time.Unix(expiryTS.Int64(), 0)
			outputIf(verbose, fmt.Sprintf("%s expires at %s", domain, expiry.Format("2006-01-02 15:04")))

			costPerSecond, err := controller.RentCost(domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain rental cost for %s", domain))
			duration := new(big.Int).Div(value, costPerSecond)
			outputIf(verbose, fmt.Sprintf("%s will be registered until approximately %v", domain, expiry.Add(time.Duration(duration.Int64())*time.Second).Format("2006-01-02 15:04")))

			opts, err := generateTxOpts(owner)
			cli.ErrCheck(err, quiet, "failed to generate transaction options")
			lastTx, err = controller.Renew(opts, domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to submit extend transaction for %s", domain))
			logTransaction(lastTx, log.Fields{
				"group":     "ens",
				"command":   "extend",
				"ensdomain": domain,
				"expiry":    expiry.Add(time.Duration(duration.Int64()) * time.Second).Format("2006-01-02 15:04"),
			})
		}
		handleSubmittedTransaction(lastTx, nil, true)
	},
}

func init() {
	ensCmd.AddCommand(ensExtendCmd)
	ensFlags(ensExtendCmd)
	ensExtendCmd.Flags().StringVar(&ensExtendDomains, "domains", "", "multiple ENS domains to extend at the same time; separate with \"&&\" e.g. --domains='mydomain1.eth&&mydomain2.eth'")
	addTransactionFlags(ensExtendCmd, "passphrase for the account that owns the domain")
}
