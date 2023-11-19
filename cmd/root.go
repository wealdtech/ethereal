// Copyright © 2017-2023 Weald Technology Trading
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
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	"github.com/wealdtech/ethereal/v2/util"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	cfgFile string
	quiet   bool
	verbose bool
	debug   bool
	offline bool
)

// c is the connection to the execution node.
var c *conn.Conn

// signer is the signer for the execution node.
var signer types.Signer

var err error

// Commands that can be run offline.
var offlineCmds = make(map[string]bool)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:              "ethereal",
	Short:            "Ethereum CLI",
	Long:             `Manage common Ethereum tasks from the command line.`,
	PersistentPreRun: persistentPreRun,
}

func persistentPreRun(cmd *cobra.Command, _ []string) {
	if cmd.Name() == "help" {
		// User just wants help.
		return
	}

	if cmd.Name() == "version" {
		// User just wants the version.
		return
	}

	// We bind viper here so that we bind to the correct command.
	quiet = viper.GetBool("quiet")
	verbose = viper.GetBool("verbose")
	debug = viper.GetBool("debug")
	offline = viper.GetBool("offline")

	// If the command does not require access to the chain then override offline accordingly.
	if offlineCmds[cmdPath(cmd)] {
		offline = true
	}

	if quiet && verbose {
		cli.Err(false, "Cannot supply both quiet and verbose flags")
	}
	if quiet && debug {
		cli.Err(false, "Cannot supply both quiet and debug flags")
	}

	// Lots of commands have transaction-related flags (e.g.) 'passphrase'
	// as options but we want to bind them to this particular command and
	// this is the first chance we get.
	if cmd.Flags().Lookup("passphrase") != nil {
		cli.ErrCheck(viper.BindPFlag("passphrase", cmd.Flags().Lookup("passphrase")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("address") != nil {
		cli.ErrCheck(viper.BindPFlag("address", cmd.Flags().Lookup("address")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("privatekey") != nil {
		cli.ErrCheck(viper.BindPFlag("privatekey", cmd.Flags().Lookup("privatekey")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("value") != nil {
		cli.ErrCheck(viper.BindPFlag("value", cmd.Flags().Lookup("value")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("wait") != nil {
		cli.ErrCheck(viper.BindPFlag("wait", cmd.Flags().Lookup("wait")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("limit") != nil {
		cli.ErrCheck(viper.BindPFlag("limit", cmd.Flags().Lookup("limit")), quiet, "failed to bind flag")
	}

	// Items that must be manually supplied if we are attempting to create transactions offline.
	if cmd.Flags().Lookup("chainid") != nil {
		cli.ErrCheck(viper.BindPFlag("chainid", cmd.Flags().Lookup("chainid")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("base-fee-per-gas") != nil {
		cli.ErrCheck(viper.BindPFlag("base-fee-per-gas", cmd.Flags().Lookup("base-fee-per-gas")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("nonce") != nil {
		cli.ErrCheck(viper.BindPFlag("nonce", cmd.Flags().Lookup("nonce")), quiet, "failed to bind flag")
	}

	// Set up gas prices.
	setUpGasPrices(cmd)

	if cmd.Flags().Lookup("gaslimit") != nil {
		cli.ErrCheck(viper.BindPFlag("gaslimit", cmd.Flags().Lookup("gaslimit")), quiet, "failed to bind flag")
	}

	// Create a connection to an Ethereum node (or mock).
	err = connect(context.Background())
	cli.ErrCheck(err, quiet, "Failed to connect to Ethereum node")
}

func setUpGasPrices(cmd *cobra.Command) {
	if cmd.Flags().Lookup("priority-fee-per-gas") == nil {
		// No gas price required.
		return
	}

	// Check for priority fee per gas, and confirm it can be used.
	cli.ErrCheck(viper.BindPFlag("priority-fee-per-gas", cmd.Flags().Lookup("priority-fee-per-gas")), quiet, "failed to bind flag")
	if viper.GetString("priority-fee-per-gas") == "" {
		// Set a default priority fee.
		viper.Set("priority-fee-per-gas", "1.5gwei")
	}
	_, err := string2eth.StringToWei(viper.GetString("priority-fee-per-gas"))
	cli.ErrCheck(err, quiet, "Invalid priority fee")

	// Check for max fee per gas, and confirm it can be used.
	cli.ErrCheck(viper.BindPFlag("max-fee-per-gas", cmd.Flags().Lookup("max-fee-per-gas")), quiet, "failed to bind flag")
	if viper.GetString("max-fee-per-gas") == "" {
		// Set a default fee.
		viper.Set("max-fee-per-gas", "200gwei")
	}
	_, err = string2eth.StringToWei(viper.GetString("max-fee-per-gas"))
	cli.ErrCheck(err, quiet, "Invalid fee")
}

// connect connects to an Ethereum node.
func connect(ctx context.Context) error {
	var err error

	if offline {
		// Handle offline connection.
		c, err = conn.New(ctx, "offline")
	} else {
		var address string
		address, err = connectionAddress(ctx)
		if err == nil {
			c, err = conn.New(ctx, address)
		}
	}
	if err != nil {
		return err
	}

	signer = types.NewLondonSigner(c.ChainID())

	return nil
}

// connectionAddress provides the address of an execution client.
func connectionAddress(_ context.Context) (string, error) {
	if viper.GetString("connection") != "" {
		return viper.GetString("connection"), nil
	}

	switch strings.ToLower(viper.GetString("network")) {
	case "mainnet":
		return "https://mainnet.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6", nil
	case "goerli", "gorli", "görli":
		return "https://goerli.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6", nil
	case "sepolia":
		return "https://sepolia.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6", nil
	case "holesky":
		return "https://holesky.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6", nil
	default:
		return "", fmt.Errorf("unknown network %s", viper.GetString("network"))
	}
}

// cmdPath recurses up the command information to create a path for this command through commands and subcommands.
func cmdPath(cmd *cobra.Command) string {
	if cmd.Parent() == nil || cmd.Parent().Name() == "ethereal" {
		return cmd.Name()
	}
	return fmt.Sprintf("%s:%s", cmdPath(cmd.Parent()), cmd.Name())
}

// setupLogging sets up the logging for commands that wish to write output.
func setupLogging() {
	logFile := viper.GetString("log")
	if logFile == "" {
		home, err := homedir.Dir()
		cli.ErrCheck(err, quiet, "Failed to access home directory")
		logFile = filepath.FromSlash(home + "/ethereal.log")
	}
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o755)
	cli.ErrCheck(err, quiet, "Failed to open log file")
	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
}

// handleSubmittedTransaction handles logging and waiting for a submitted transaction to be mined.
// It will not log the transaction if logFields is nil.
// If exit is true this function will exit with a suitable status.
// If exit is false this function will return false if asked to wait and the transaction is not
// mined, otherwise true.
func handleSubmittedTransaction(tx *types.Transaction, logFields log.Fields, exit bool) bool {
	if logFields != nil {
		logTransaction(tx, logFields)
	}

	if offline {
		return true
	}

	if !viper.GetBool("wait") {
		outputIf(!quiet, tx.Hash().Hex())
		if exit {
			os.Exit(exitSuccess)
		}
		return true
	}
	mined := util.WaitForTransaction(c.Client(), tx.Hash(), viper.GetDuration("limit"))
	if mined {
		outputIf(!quiet, fmt.Sprintf("%s mined", tx.Hash().Hex()))
		if exit {
			os.Exit(exitSuccess)
		}
		return true
	}
	outputIf(!quiet, fmt.Sprintf("%s submitted but not mined", tx.Hash().Hex()))
	if exit {
		os.Exit(exitNotMined)
	}
	return false
}

// logTransaction logs a transaction.
func logTransaction(tx *types.Transaction, fields log.Fields) {
	setupLogging()

	txFields := log.Fields{
		"networkid":            c.ChainID(),
		"transactionid":        tx.Hash().Hex(),
		"gas":                  tx.Gas(),
		"fee-per-gas":          tx.GasFeeCap().String(),
		"priority-fee-per-gas": tx.GasTipCap().String(),
		"value":                tx.Value().String(),
		"data":                 hex.EncodeToString(tx.Data()),
	}
	fromAddress, err := types.Sender(signer, tx)
	if err == nil {
		txFields["from"] = fromAddress.Hex()
	}
	if tx.To() != nil {
		txFields["to"] = tx.To().Hex()
	}

	log.WithFields(fields).WithFields(txFields).Info("transaction submitted")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ethereal.yaml)")
	RootCmd.PersistentFlags().String("log", "", "log activity to the named file (default $HOME/ethereal.log).  Logs are written for every action that generates a transaction")
	if err := viper.BindPFlag("log", RootCmd.PersistentFlags().Lookup("log")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Bool("quiet", false, "do not generate any output")
	if err := viper.BindPFlag("quiet", RootCmd.PersistentFlags().Lookup("quiet")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Bool("verbose", false, "generate additional output where appropriate")
	if err := viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Bool("debug", false, "generate debug output")
	if err := viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().String("connection", "", "the custom IPC or RPC path to an Ethereum node (overrides network option).  If you are running your own local instance of Ethereum this might be /home/user/.ethereum/geth.ipc (IPC) or http://localhost:8545/ (RPC)")
	if err := viper.BindPFlag("connection", RootCmd.PersistentFlags().Lookup("connection")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().String("network", "mainnet", "network to access (mainnet/goerli/sepolia/holesky) (overridden by connection option)")
	if err := viper.BindPFlag("network", RootCmd.PersistentFlags().Lookup("network")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Duration("timeout", 30*time.Second, "the time after which a network request will be deemed to have failed.  Increase this if you are running on a error-prone, high-latency or low-bandwidth connection")
	if err := viper.BindPFlag("timeout", RootCmd.PersistentFlags().Lookup("timeout")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Bool("offline", false, "work without a connection to an execution node")
	if err := viper.BindPFlag("offline", RootCmd.PersistentFlags().Lookup("offline")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Int("usbwallets", 1, "number of USB wallets to show")
	if err := viper.BindPFlag("usbwallets", RootCmd.PersistentFlags().Lookup("usbwallets")); err != nil {
		panic(err)
	}
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
			os.Exit(exitFailure)
		}

		// Search config in home directory with name ".ethereal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ethereal")
	}

	viper.AutomaticEnv() // read in environment variables that match.

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	switch err.(type) {
	case nil:
		outputIf(viper.GetBool("debug"), fmt.Sprintf("Loaded configuration file %s", viper.ConfigFileUsed()))
	case viper.ConfigFileNotFoundError:
		if cfgFile != "" {
			cli.Err(viper.GetBool("quiet"), fmt.Sprintf("Failed to access configuration file %s", viper.ConfigFileUsed()))
		} else {
			outputIf(viper.GetBool("debug"), "Using default configuration")
		}
	default:
		cli.Err(viper.GetBool("quiet"), fmt.Sprintf("Failed to load configuration file %s: %s", viper.ConfigFileUsed(), err))
	}
}

//
// Helpers
//

// Add flags for commands that carry out transactions.
func addTransactionFlags(cmd *cobra.Command, explanation string) {
	cmd.Flags().String("passphrase", "", fmt.Sprintf("passphrase for %s", explanation))
	cmd.Flags().String("privatekey", "", fmt.Sprintf("private key for %s", explanation))
	cmd.Flags().String("max-fee-per-gas", "200Gwei", "Maximum fee per gas for transaction e.g. 15Gwei, 0.000000015ether")
	cmd.Flags().String("priority-fee-per-gas", "1.5Gwei", "Priority fee per gas for transaction e.g. 1gwei")
	cmd.Flags().String("value", "", "Ether to send with the transaction")
	cmd.Flags().Int64("gaslimit", 0, "Gas limit for the transaction; 0 is auto-select")
	cmd.Flags().String("chainid", "", "chain ID; only needed when offline")
	cmd.Flags().String("base-fee-per-gas", "", "base fee per gas; only needed when offline e.g. 30Gwei")
	cmd.Flags().String("nonce", "", "nonce for account; only needed when offline")
	cmd.Flags().Bool("wait", false, "wait for the transaction to be mined before returning")
	cmd.Flags().Duration("limit", 0, "maximum time to wait for transaction to complete before failing (default forever)")
}

func generateTxOpts(sender common.Address) (*bind.TransactOpts, error) {
	// Signer depends on what information is available to us.
	var signer bind.SignerFn
	if viper.GetString("passphrase") != "" {
		wallet, account, err := cli.ObtainWalletAndAccount(c.ChainID(), sender)
		if err != nil {
			return nil, err
		}
		signer = util.AccountSigner(c.ChainID(), &wallet, account, viper.GetString("passphrase"))
	} else if viper.GetString("privatekey") != "" {
		key, err := crypto.HexToECDSA(strings.TrimPrefix(viper.GetString("privatekey"), "0x"))
		cli.ErrCheck(err, quiet, "Invalid private key")
		signer = util.KeySigner(c.ChainID(), key)
	}
	if signer == nil {
		return nil, fmt.Errorf("no signer; please supply either passphrase or private key")
	}

	var value *big.Int
	if viper.GetString("value") != "" {
		value, err = string2eth.StringToWei(viper.GetString("value"))
		cli.ErrCheck(err, quiet, "Failed to understand value")
	}

	curNonce, err := c.CurrentNonce(context.Background(), sender)
	if err != nil {
		return nil, err
	}

	// Calculate the fees.
	feePerGas, priorityFeePerGas, err := calculateFees()
	if err != nil {
		return nil, err
	}

	opts := &bind.TransactOpts{
		From:      sender,
		Signer:    signer,
		GasFeeCap: feePerGas,
		GasTipCap: priorityFeePerGas,
		Value:     value,
		NoSend:    offline,
		Nonce:     big.NewInt(0).SetUint64(curNonce),
	}

	limit := uint64(viper.GetInt64("gaslimit"))
	if limit > 0 {
		opts.GasLimit = limit
	}

	return opts, nil
}

func outputIf(condition bool, msg string) {
	if condition {
		fmt.Println(msg)
	}
}

func localContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
}

func calculateFees() (*big.Int, *big.Int, error) {
	baseFeePerGas, err := c.CurrentBaseFee(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to obtain current base fee")
	}
	outputIf(debug, fmt.Sprintf("Current base fee per gas: %v", string2eth.WeiToString(baseFeePerGas, true)))

	if viper.GetString("max-fee-per-gas") == "" {
		viper.Set("max-fee-per-gas", "200gwei")
	}
	maxFeePerGas, err := string2eth.StringToWei(viper.GetString("max-fee-per-gas"))
	if err != nil {
		return nil, nil, err
	}
	outputIf(debug, fmt.Sprintf("Max fee per gas: %v", string2eth.WeiToString(maxFeePerGas, true)))

	// Obtain priority fee per gas.
	priorityFeePerGas, err := string2eth.StringToWei(viper.GetString("priority-fee-per-gas"))
	cli.ErrCheck(err, quiet, "Failed to obtain priority fee per gas")

	// Ensure that the total fee per gas does not exceed the max allowed.
	totalFeePerGas := new(big.Int).Add(baseFeePerGas, priorityFeePerGas)
	outputIf(debug, fmt.Sprintf("Total fee per gas is %s", string2eth.WeiToString(totalFeePerGas, true)))
	if totalFeePerGas.Cmp(maxFeePerGas) > 0 {
		return nil, nil, fmt.Errorf("calculated total fee per gas of %s too high; increase with --max-fee-per-gas if you are sure you want to do this", string2eth.WeiToString(totalFeePerGas, true))
	}

	// Try to double the base fee to allow for changes in future block base fee, but not exceed the max allowed.
	doubleBaseFeePerGas := new(big.Int).Add(baseFeePerGas, baseFeePerGas)

	feePerGas := new(big.Int).Add(doubleBaseFeePerGas, priorityFeePerGas)
	if feePerGas.Cmp(maxFeePerGas) >= 0 {
		feePerGas = feePerGas.Sub(maxFeePerGas, priorityFeePerGas)
		outputIf(debug, fmt.Sprintf("Total fee per gas is higher than max fee per gas, reduced to %ss", string2eth.WeiToString(feePerGas, true)))
	}
	outputIf(debug, fmt.Sprintf("Final fee per gas is %s, priority fee per gas is %s", string2eth.WeiToString(feePerGas, true), string2eth.WeiToString(priorityFeePerGas, true)))

	return feePerGas, priorityFeePerGas, nil
}
