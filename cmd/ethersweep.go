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
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/orinocopay/go-etherutils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var etherSweepFromAddress string
var etherSweepToAddress string

// etherSweepCmd represents the ether sweep command
var etherSweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Sweep funds to a given address",
	Long: `Sweep all Ether funds from one address to another.  For example:

    etherereal ether sweep --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --passphrase=secret

In quiet mode this will return 0 if the sweep transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(etherSweepFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, etherSweepFromAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain from address for sweep")

		cli.Assert(etherSweepToAddress != "", quiet, "--to is required")
		toAddress, err := ens.Resolve(client, etherSweepToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain to address for sweep")

		// Obtain the balance of the address
		ctx, cancel := localContext()
		defer cancel()
		balance, err := client.BalanceAt(ctx, fromAddress, nil)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(big.NewInt(0)) > 0, quiet, fmt.Sprintf("Balance of %s is 0; nothing to sweep", fromAddress.Hex()))

		// Obtain the amount of gas required to send the transaction, and calculate the amount to send
		gas, err := estimateGas(fromAddress, &toAddress, balance, nil)
		cli.ErrCheck(err, quiet, "Failed to estimate gas required to sweep funds")
		outputIf(verbose, fmt.Sprintf("Gas estimation is %v", gas))
		gasCost := big.NewInt(0).Mul(big.NewInt(int64(gas)), gasPrice)
		outputIf(verbose, fmt.Sprintf("Gas cost is %v", etherutils.WeiToString(gasCost, true)))
		amount := balance.Sub(balance, gasCost)
		outputIf(verbose, fmt.Sprintf("Sweeping %s", etherutils.WeiToString(amount, true)))

		// Create and sign the transaction
		signedTx, err := createSignedTransaction(fromAddress, &toAddress, amount, gasLimit, nil)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			ctx, cancel := localContext()
			defer cancel()
			err = client.SendTransaction(ctx, signedTx)
			cli.ErrCheck(err, quiet, "Failed to send transaction")

			setupLogging()
			log.WithFields(log.Fields{
				"group":         "ether",
				"command":       "sweep",
				"from":          fromAddress.Hex(),
				"to":            toAddress.Hex(),
				"amount":        amount.String(),
				"networkid":     chainID,
				"gas":           signedTx.Gas(),
				"gasprice":      signedTx.GasPrice().String(),
				"transactionid": signedTx.Hash().Hex(),
			}).Info("success")
			if quiet {
				os.Exit(0)
			}
			fmt.Println(signedTx.Hash().Hex())
		}
	},
}

func init() {
	initAliases(etherSweepCmd)
	etherCmd.AddCommand(etherSweepCmd)
	etherSweepCmd.Flags().StringVar(&etherSweepFromAddress, "from", "", "Address from which to sweep Ether")
	etherSweepCmd.Flags().StringVar(&etherSweepToAddress, "to", "", "Address to which to sweep Ether")
	addTransactionFlags(etherSweepCmd, "the address that holds the funds")
}
