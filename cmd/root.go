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
var logFile string
var quiet bool
var connection string

var client *ethclient.Client
var chainID *big.Int
var referrer common.Address
var nonce int64

// Common variables
var passphrase string
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

	// Set default log file if no alternative is provided
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
	client, err = ethclient.Dial(connection)
	cli.ErrCheck(err, quiet, "Failed to connect to Ethereum")
	// Fetch the chain ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
	RootCmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "log activity to the named file (default $HOME/.ethereal.log)")
	RootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "no output")
	RootCmd.PersistentFlags().StringVarP(&connection, "connection", "c", "https://api.orinocopay.com:8546/", "path to the Ethereum connection")
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

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//
// Helpers
//

// Add flags for commands that carry out transactions
func addTransactionFlags(cmd *cobra.Command, passphraseExplanation string) {
	cmd.Flags().StringVarP(&passphrase, "passphrase", "p", "", passphraseExplanation)
	cmd.Flags().StringVarP(&gasPriceStr, "gasprice", "g", "4 GWei", "Gas price for the transaction")
	cmd.Flags().Int64VarP(&nonce, "nonce", "n", -1, "Nonce for the transaction; -1 is auto-select")
}

// Augment session information
func augmentSession(session *domainsalecontract.DomainSaleContractSession) {
	session.TransactOpts.GasPrice = gasPrice
	if nonce != -1 {
		session.TransactOpts.Nonce = big.NewInt(nonce)
	}
}

func obtainWalletAndAccount(address common.Address, passphrase string) (wallet accounts.Wallet, account *accounts.Account, err error) {
	wallet, err = cli.ObtainWallet(chainID, address)
	if err == nil {
		account, err = cli.ObtainAccount(&wallet, &address, passphrase)
	}
	return wallet, account, err
}
