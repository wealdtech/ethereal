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
	"context"
	"fmt"
	"math/big"
	"os"

	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var etherSweepToAddress string

// etherSweepCmd represents the ether sweep command
var etherSweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Sweep funds to a given address",
	Long: `Sweep all Ether funds from one address to another.  For example:

    etherereal ether sweep --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --passphrase=secret 0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the sweep transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(len(args) == 1, quiet, "Requires a single address from which to sweep funds")
		cli.Assert(args[0] != "", quiet, "Sender address is required")

		fromAddress, err := ens.Resolve(client, args[0])
		cli.ErrCheck(err, quiet, "Failed to obtain sender for sweep")

		toAddress, err := ens.Resolve(client, etherSweepToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain recipient for sweep")

		// Obtain the balance of the address
		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(big.NewInt(0)) > 0, quiet, fmt.Sprintf("Balance of %s is 0; nothing to sweep", args[0]))

		// Obtain the amount of gas required to send the transaction, and calculate the amount to send
		gas, err := estimateGas(fromAddress, &toAddress, balance, nil)
		cli.ErrCheck(err, quiet, "Failed to estimate gas required to sweep funds")
		amount := balance.Sub(balance, gas.Mul(gas, gasPrice))

		fmt.Printf("%s - %s = %s\n", etherutils.WeiToString(balance, true), etherutils.WeiToString(gas, true), etherutils.WeiToString(amount, true))

		// Create and sign the transaction
		signedTx, err := createSignedTransaction(fromAddress, &toAddress, amount, nil)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		err = client.SendTransaction(context.Background(), signedTx)
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		log.WithFields(log.Fields{
			"group":         "ether",
			"command":       "sweep",
			"from":          fromAddress.Hex(),
			"to":            toAddress.Hex(),
			"amount":        amount.String(),
			"networkid":     chainID,
			"transactionid": signedTx.Hash().Hex(),
		}).Info("success")

		if quiet {
			os.Exit(0)
		}
		fmt.Println(signedTx.Hash().Hex())
	},
}

func init() {
	etherCmd.AddCommand(etherSweepCmd)
	etherSweepCmd.Flags().StringVar(&etherSweepToAddress, "to", "", "Address to which to sweep Ether")
	addTransactionFlags(etherSweepCmd, "Passphrase for the address that holds the funds")
}
