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
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
	ens "github.com/wealdtech/go-ens/v3"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	ensRegisterDomains  string
	ensRegisterOwnerStr string
)

// ensRegisterCmd represents the register command.
var ensRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an ENS domain",
	Long: `Register an Ethereum Name Service (ENS) domain.  For example:

    ethereal ens register --domain=enstest.eth --passphrase="my secret passphrase"

Multiple domains can be registered with a single command (note this still creates one transaction for each name).  For example:

    ethereal ens register --domains=mydomain1.eth&&mydomain2.eth --passphrase="my secret passphrase"

The keystore for the domain(s) owner must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transactions are successfully submitted (and mined if --wait is supplied), 1 if the transactions are not successfully submitted, and 2 if the transactions are successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "" || ensRegisterDomains != "", quiet, "--domain or --domains is required")

		cli.Assert(ensRegisterOwnerStr != "", quiet, "--owner is required")
		owner, err := c.Resolve(ensRegisterOwnerStr)
		cli.ErrCheck(err, quiet, "Failed to obtain new owner address")
		cli.Assert(!bytes.Equal(owner.Bytes(), ens.UnknownAddress.Bytes()), quiet, "Unknown owner")

		cli.Assert(viper.GetString("value") != "", quiet, "--value is required")
		value, err := string2eth.StringToWei(viper.GetString("value"))
		cli.ErrCheck(err, quiet, "Could not understand value")

		var domains []string
		if ensRegisterDomains != "" {
			domains = strings.Split(ensRegisterDomains, "&&")
		} else {
			domains = make([]string, 1)
			domains[0] = ensDomain
		}

		controller, err := ens.NewETHController(c.Client(), ens.Domain(domains[0]))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain %s controller", ens.Domain(domains[0])))

		tmp, err := controller.MinCommitmentInterval()
		cli.ErrCheck(err, quiet, "Failed to find out minimum commitment interval")
		// Interval is min commitment interval plus 2 minutes.
		interval := time.Duration(tmp.Int64()+120) * time.Second
		minDuration, err := controller.MinRegistrationDuration()
		cli.ErrCheck(err, quiet, "Failed to obtain minimum registration duration")

		// Check loop.
		for _, domain := range domains {
			domain, err = ens.NormaliseDomain(domain)
			cli.ErrCheck(err, quiet, "Failed to normalise ENS domain")

			valid, err := controller.IsValid(domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to find out if %s is valid", domain))
			cli.Assert(valid, quiet, fmt.Sprintf("%s is not valid", domain))

			available, err := controller.IsAvailable(domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to find out if %s is available", domain))
			cli.Assert(available, quiet, fmt.Sprintf("%s is not available", domain))

			costPerSecond, err := controller.RentCost(domain)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain rental cost for %s", domain))
			duration := new(big.Int).Div(value, costPerSecond)
			// Ensure duration is greater than minimum duration.
			cli.Assert(big.NewInt(int64(minDuration.Seconds())).Cmp(duration) <= 0, quiet, fmt.Sprintf("Not enough funds to cover minimum duration of %v for %s", minDuration, domain))
			outputIf(verbose, fmt.Sprintf("%s will be registered until approximately %v", domain, time.Now().Add(time.Duration(duration.Int64())*time.Second).Format("2006-01-02 15:04")))
		}

		// Commit loop.
		secrets := make(map[string][32]byte)
		var lastTx *types.Transaction
		for _, domain := range domains {
			domain, err = ens.NormaliseDomain(domain)
			cli.ErrCheck(err, quiet, "Failed to normalise ENS domain")

			var secret [32]byte
			_, err = rand.Read(secret[:])
			cli.ErrCheck(err, quiet, "failed to generate secret")
			secrets[domain] = secret

			opts, err := generateTxOpts(owner)
			cli.ErrCheck(err, quiet, "failed to generate transaction options")

			// Commitments have no value.
			opts.Value = nil
			cli.ErrCheck(err, quiet, "failed to generate commit transaction options")
			lastTx, err = controller.Commit(opts, domain, owner, secrets[domain])
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to submit commit transaction for %s", domain))
			logTransaction(lastTx, log.Fields{
				"group":     "ens",
				"command":   "register",
				"stage":     "commit",
				"ensdomain": domain,
				"ensowner":  owner.Hex(),
				"secret":    hex.EncodeToString(secret[:]),
			})
			outputIf(verbose, fmt.Sprintf("Commit transaction %x submitted for %s", lastTx.Hash(), domain))
			_, err = c.NextNonce(context.Background(), owner)
			cli.ErrCheck(err, quiet, "failed to increment nonce")
		}

		// Wait.
		outputIf(!quiet, "Waiting for commit transaction(s) to be mined")
		mined := util.WaitForTransaction(c.Client(), lastTx.Hash(), 0)
		cli.Assert(mined, quiet, "Failed to mine commit transaction(s)")
		outputIf(!quiet, fmt.Sprintf("Waiting for commit/reveal interval to pass (done at %s)", time.Now().Add(interval).Format("15:04:05")))
		time.Sleep(interval)

		// Reveal loop.
		for _, domain := range domains {
			domain, err = ens.NormaliseDomain(domain)
			cli.ErrCheck(err, quiet, "Failed to normalise ENS domain")

			opts, err := generateTxOpts(owner)
			cli.ErrCheck(err, quiet, "failed to generate reveal transaction options")
			lastTx, err = controller.Reveal(opts, domain, owner, secrets[domain])
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to submit reveal transaction for %s", domain))
			secret := secrets[domain]
			logTransaction(lastTx, log.Fields{
				"group":     "ens",
				"command":   "register",
				"stage":     "reveal",
				"ensdomain": domain,
				"ensowner":  owner.Hex(),
				"secret":    hex.EncodeToString(secret[:]),
			})
			outputIf(verbose, fmt.Sprintf("Reveal transaction %x submitted for %s", lastTx.Hash(), domain))
			_, err = c.NextNonce(context.Background(), owner)
			cli.ErrCheck(err, quiet, "failed to increment nonce")
		}
		handleSubmittedTransaction(lastTx, nil, true)
	},
}

func init() {
	ensCmd.AddCommand(ensRegisterCmd)
	ensFlags(ensRegisterCmd)
	ensRegisterCmd.Flags().StringVar(&ensRegisterDomains, "domains", "", "multiple ENS domains to register at the same time; separate with \"&&\" e.g. --domains='mydomain1.eth&&mydomain2.eth'")
	ensRegisterCmd.Flags().StringVar(&ensRegisterOwnerStr, "owner", "", "The owner's name or address")
	addTransactionFlags(ensRegisterCmd, "passphrase for the account that owns the domain")
}
