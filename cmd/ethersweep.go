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
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	ens "github.com/wealdtech/go-ens/v3"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	etherSweepFromAddress string
	etherSweepToAddress   string
)

// etherSweepCmd represents the ether sweep command.
var etherSweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Sweep funds to a given address",
	Long: `Sweep all Ether funds from one address to another.  For example:

    etherereal ether sweep --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(etherSweepFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(etherSweepFromAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain from address for sweep")

		cli.Assert(etherSweepToAddress != "", quiet, "--to is required")
		toAddress, err := c.Resolve(etherSweepToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain to address for sweep")

		// Obtain the balance of the address.
		ctx, cancel := localContext()
		defer cancel()
		balance, err := c.Client().BalanceAt(ctx, fromAddress, nil)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(big.NewInt(0)) > 0, quiet, fmt.Sprintf("Balance of %s is 0; nothing to sweep", ens.Format(c.Client(), fromAddress)))

		// Obtain the amount of gas required to send the transaction, and calculate the amount to send.
		gas, err := c.EstimateGas(context.Background(), &conn.TransactionData{
			From:  fromAddress,
			To:    &toAddress,
			Value: balance,
		})
		cli.ErrCheck(err, quiet, "Failed to estimate gas required to sweep funds")
		outputIf(verbose, fmt.Sprintf("Gas estimation is %v", gas))

		// Obtain current base fee, multiply it by 150%.
		baseFee, err := c.CurrentBaseFee(context.Background())
		cli.ErrCheck(err, quiet, "failed to obtain current base fee")
		gasFee := new(big.Int).Div(baseFee.Mul(baseFee, big.NewInt(3)), big.NewInt(2))

		gasCost := new(big.Int).Mul(big.NewInt(int64(gas)), gasFee)
		outputIf(verbose, fmt.Sprintf("Gas cost is %v", string2eth.WeiToString(gasCost, true)))
		amount := balance.Sub(balance, gasCost)
		outputIf(verbose, fmt.Sprintf("Sweeping %s", string2eth.WeiToString(amount, true)))

		var gasLimit *uint64
		limit := uint64(viper.GetInt64("gaslimit"))
		if limit > 0 {
			gasLimit = &limit
		}

		// Create and sign the transaction.
		signedTx, err := c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
			From:                 fromAddress,
			To:                   &toAddress,
			Value:                amount,
			GasLimit:             gasLimit,
			MaxFeePerGas:         gasFee,
			MaxPriorityFeePerGas: gasFee,
		})
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
			os.Exit(exitSuccess)
		}
		err = c.SendTransaction(context.Background(), signedTx)
		cli.ErrCheck(err, quiet, "Failed to send transaction")
		handleSubmittedTransaction(signedTx, log.Fields{
			"group":   "ether",
			"command": "sweep",
		}, true)
	},
}

func init() {
	etherCmd.AddCommand(etherSweepCmd)
	etherSweepCmd.Flags().StringVar(&etherSweepFromAddress, "from", "", "Address from which to sweep Ether")
	etherSweepCmd.Flags().StringVar(&etherSweepToAddress, "to", "", "Address to which to sweep Ether")
	addTransactionFlags(etherSweepCmd, "the address that holds the funds")
}
