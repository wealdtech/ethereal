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
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	homedir "github.com/mitchellh/go-homedir"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/go-domainsale/domainsalecontract"
)

var cfgFile string
var quiet bool

var client *ethclient.Client
var chainID *big.Int
var referrer common.Address
var nonce int64

// Common variables
var gasPriceStr string
var gasPrice *big.Int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:              "ethereal",
	Short:            "Ethereum CLI",
	Long:             `Manage common Ethereum tasks from the command line.`,
	PersistentPreRun: persistentPreRun,
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	if cmd.Name() == "help" {
		// User just wants help
		return
	}

	if cmd.Name() == "version" {
		// User just wants the version
		return
	}

	quiet = viper.GetBool("quiet")

	fmt.Println("Timeout is", viper.GetDuration("timeout"))
	os.Exit(0)
	// Set default log file if no alternative is provided
	logFile := viper.GetString("log")
	if logFile == "" {
		home, err := homedir.Dir()
		cli.ErrCheck(err, quiet, "Failed to access home directory")
		logFile = filepath.FromSlash(home + "/.ethereal.log")
	}
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	cli.ErrCheck(err, quiet, "Failed to open log file")
	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})

	// Create a connection to an Ethereum node
	client, err = ethclient.Dial(viper.GetString("connection"))
	cli.ErrCheck(err, quiet, "Failed to connect to Ethereum")
	// Fetch the chain ID
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
	defer cancel()
	chainID, err = client.NetworkID(ctx)
	cli.ErrCheck(err, quiet, "Failed to obtain chain ID")

	// Set up gas price
	if gasPriceStr == "" {
		gasPrice, err = etherutils.StringToWei("4 GWei")
		cli.ErrCheck(err, quiet, "Invalid gas price")
	} else {
		gasPrice, err = etherutils.StringToWei(gasPriceStr)
		cli.ErrCheck(err, quiet, "Invalid gas price")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ethereal.yaml)")
	RootCmd.PersistentFlags().String("log", "", "log activity to the named file (default $HOME/domainsale.log)")
	viper.BindPFlag("log", RootCmd.PersistentFlags().Lookup("log"))
	RootCmd.PersistentFlags().Bool("quiet", false, "no output")
	viper.BindPFlag("quiet", RootCmd.PersistentFlags().Lookup("quiet"))
	RootCmd.PersistentFlags().String("connection", "https://api.orinocopay.com:8546/", "path to the Ethereum connection")
	viper.BindPFlag("connection", RootCmd.PersistentFlags().Lookup("connection"))
	RootCmd.PersistentFlags().Duration("timeout", 30*time.Second, "Time to wait for network calls before failing")
	viper.BindPFlag("timeout", RootCmd.PersistentFlags().Lookup("timeout"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ethereal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ethereal")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to read configuration file %s", viper.ConfigFileUsed()))
}

//
// Helpers
//

// Add flags for commands that carry out transactions
func addTransactionFlags(cmd *cobra.Command, passphraseExplanation string) {
	cmd.Flags().String("passphrase", "", passphraseExplanation)
	viper.BindPFlag("passphrase", cmd.Flags().Lookup("passphrase"))
	cmd.Flags().StringVar(&gasPriceStr, "gasprice", "4 GWei", "Gas price for the transaction")
	cmd.Flags().Int64Var(&nonce, "nonce", -1, "Nonce for the transaction; -1 is auto-select")
}

// Augment session information
func augmentSession(session *domainsalecontract.DomainSaleContractSession) {
	session.TransactOpts.GasPrice = gasPrice
	if nonce != -1 {
		session.TransactOpts.Nonce = big.NewInt(nonce)
	}
}

func obtainWalletAndAccount(address common.Address) (wallet accounts.Wallet, account *accounts.Account, err error) {
	wallet, err = cli.ObtainWallet(chainID, address)
	if err == nil {
		account, err = cli.ObtainAccount(&wallet, &address, viper.GetString("passphrase"))
	}
	return wallet, account, err
}

var wallet accounts.Wallet
var account *accounts.Account

func signTransaction(signer common.Address, tx *types.Transaction) (signedTx *types.Transaction, err error) {
	if wallet == nil {
		// Fetch the wallet and account for the sender
		wallet, account, err = obtainWalletAndAccount(signer)
		if err != nil {
			return
		}
	}
	signedTx, err = wallet.SignTxWithPassphrase(*account, viper.GetString("passphrase"), tx, chainID)
	return
}
