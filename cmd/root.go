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
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
	string2eth "github.com/wealdtech/go-string2eth"
)

var cfgFile string
var quiet bool
var verbose bool
var debug bool
var offline bool

var client *ethclient.Client
var chainID *big.Int

var nonce int64
var wallet accounts.Wallet
var account *accounts.Account

// Common variables
var gasPrice *big.Int
var maxSaneGasPrice = big.NewInt(1000000000000)
var gasLimit uint64

var err error

// Commands that can be run offline
var offlineCmds = make(map[string]bool)

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

	// We bind viper here so that we bind to the correct command
	quiet = viper.GetBool("quiet")
	verbose = viper.GetBool("verbose")
	debug = viper.GetBool("debug")
	offline = viper.GetBool("offline")

	// If the command does not require access to the chain then override offline accordingly
	if offlineCmds[cmdPath(cmd)] {
		offline = true
	}

	switch strings.ToLower(viper.GetString("network")) {
	case "mainnet":
		chainID = big.NewInt(1)
	case "ropsten":
		chainID = big.NewInt(3)
	case "rinkeby":
		chainID = big.NewInt(4)
	case "goerli", "gorli", "görli":
		chainID = big.NewInt(5)
	case "sepolia":
		chainID = big.NewInt(11155111)
	case "kovan":
		chainID = big.NewInt(42)
	default:
		cli.Err(quiet, fmt.Sprintf("Unknown network name %q", viper.GetString("network")))
	}

	if quiet && verbose {
		cli.Err(false, "Cannot supply both quiet and verbose flags")
	}
	if quiet && debug {
		cli.Err(false, "Cannot supply both quiet and debug flags")
	}

	// ...lots of commands have transaction-related flags (e.g.) 'passphrase'
	// as options but we want to bind them to this particular command and
	// this is the first chance we get
	if cmd.Flags().Lookup("passphrase") != nil {
		cli.ErrCheck(viper.BindPFlag("passphrase", cmd.Flags().Lookup("passphrase")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("privatekey") != nil {
		cli.ErrCheck(viper.BindPFlag("privatekey", cmd.Flags().Lookup("privatekey")), quiet, "failed to bind flag")
	}
	if cmd.Flags().Lookup("nonce") != nil {
		cli.ErrCheck(viper.BindPFlag("nonce", cmd.Flags().Lookup("nonce")), quiet, "failed to bind flag")
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
	// Set up gas price if we have it
	if cmd.Flags().Lookup("gasprice") != nil {
		cli.ErrCheck(viper.BindPFlag("gasprice", cmd.Flags().Lookup("gasprice")), quiet, "failed to bind flag")
		cli.ErrCheck(viper.BindPFlag("allowhighgasprice", cmd.Flags().Lookup("allowhighgasprice")), quiet, "failed to bind flag")
		if viper.GetString("gasprice") == "" {
			gasPrice, err = string2eth.StringToWei("4 GWei")
			cli.ErrCheck(err, quiet, "Invalid gas price")
		} else {
			switch {
			case strings.Contains(viper.GetString("gasprice"), "block"):
				// Block-based gas price
				outputIf(verbose, "xx")
				// fmt.Printf("Gas price is %v\n", string2eth.WeiToString(gasPrice, true))
				os.Exit(exitSuccess)
			case strings.Contains(viper.GetString("gasprice"), "minute"):
				// Time-based gas price
				outputIf(verbose, "yy")
				// fmt.Printf("Gas price is %v\n", string2eth.WeiToString(gasPrice, true))
				os.Exit(exitSuccess)
			default:
				gasPrice, err = string2eth.StringToWei(viper.GetString("gasprice"))
				cli.ErrCheck(err, quiet, "Invalid gas price")
			}
		}
		cli.Assert(gasPrice.Cmp(maxSaneGasPrice) <= 0 || viper.GetBool("allowhighgasprice"), quiet, "Gas price set very high.  If you are sure this is what you want you may add the --allowhighgasprice flag to continue.")
	}

	// Set up nonce if we have it
	nonce = viper.GetInt64("nonce")

	if cmd.Flags().Lookup("gaslimit") != nil {
		cli.ErrCheck(viper.BindPFlag("gaslimit", cmd.Flags().Lookup("gaslimit")), quiet, "failed to bind flag")
		if viper.GetInt("gaslimit") > 0 {
			gasLimit = uint64(viper.GetInt("gaslimit"))
		}
	}

	// Create a connection to an Ethereum node
	if !offline {
		err = connect()
		cli.ErrCheck(err, quiet, "Failed to connect to Ethereum node")
	}
}

// connect connects to an Ethereum node
func connect() error {
	var err error
	if viper.GetString("connection") != "" {
		outputIf(debug, fmt.Sprintf("Connecting to %s", viper.GetString("connection")))
		client, err = ethclient.Dial(viper.GetString("connection"))
	} else {
		switch strings.ToLower(viper.GetString("network")) {
		case "mainnet":
			outputIf(debug, "Connecting to mainnet")
			client, err = ethclient.Dial("https://mainnet.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
		case "ropsten":
			outputIf(debug, "Connecting to ropsten")
			client, err = ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
		case "rinkeby":
			outputIf(debug, "Connecting to rinkeby")
			client, err = ethclient.Dial("https://rinkeby.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
		case "goerli", "gorli", "görli":
			outputIf(debug, "Connecting to goerli")
			client, err = ethclient.Dial("https://goerli.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
		case "sepolia":
			outputIf(debug, "Connecting to sepolia")
			client, err = ethclient.Dial("https://sepolia.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
		case "kovan":
			outputIf(debug, "Connecting to kovan")
			client, err = ethclient.Dial("https://kovan.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
		default:
			cli.Err(quiet, fmt.Sprintf("Unknown network %s", viper.GetString("network")))
		}
	}
	cli.ErrCheck(err, quiet, "Failed to connect to network")
	// Fetch the chain ID
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
	defer cancel()
	chainID, err = client.NetworkID(ctx)
	return err
}

// cmdPath recurses up the command information to create a path for this command through commands and subcommands
func cmdPath(cmd *cobra.Command) string {
	if cmd.Parent() == nil || cmd.Parent().Name() == "ethereal" {
		return cmd.Name()
	}
	return fmt.Sprintf("%s:%s", cmdPath(cmd.Parent()), cmd.Name())
}

// setupLogging sets up the logging for commands that wish to write output
func setupLogging() {
	logFile := viper.GetString("log")
	if logFile == "" {
		home, err := homedir.Dir()
		cli.ErrCheck(err, quiet, "Failed to access home directory")
		logFile = filepath.FromSlash(home + "/ethereal.log")
	}
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
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

	if !viper.GetBool("wait") {
		outputIf(!quiet, tx.Hash().Hex())
		if exit {
			os.Exit(exitSuccess)
		} else {
			return true
		}
	}
	mined := util.WaitForTransaction(client, tx.Hash(), viper.GetDuration("limit"))
	if mined {
		outputIf(!quiet, fmt.Sprintf("%s mined", tx.Hash().Hex()))
		if exit {
			os.Exit(exitSuccess)
		} else {
			return true
		}
	}
	outputIf(!quiet, fmt.Sprintf("%s submitted but not mined", tx.Hash().Hex()))
	if exit {
		os.Exit(exitNotMined)
	}
	return false
}

// logTransaction logs a transaction
func logTransaction(tx *types.Transaction, fields log.Fields) {
	setupLogging()

	txFields := log.Fields{
		"networkid":     chainID,
		"transactionid": tx.Hash().Hex(),
		"gas":           tx.Gas(),
		"gasprice":      tx.GasPrice().String(),
		"value":         tx.Value().String(),
		"data":          hex.EncodeToString(tx.Data()),
	}
	fromAddress, err := txFrom(tx)
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
	RootCmd.PersistentFlags().String("network", "mainnet", "network to access (mainnet/ropsten/kovan/rinkeby/goerli/sepolia) (overridden by connection option)")
	if err := viper.BindPFlag("network", RootCmd.PersistentFlags().Lookup("network")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Duration("timeout", 30*time.Second, "the time after which a network request will be deemed to have failed.  Increase this if you are running on a error-prone, high-latency or low-bandwidth connection")
	if err := viper.BindPFlag("timeout", RootCmd.PersistentFlags().Lookup("timeout")); err != nil {
		panic(err)
	}
	RootCmd.PersistentFlags().Bool("offline", false, "print the transaction a hex string and do not send it")
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

	viper.AutomaticEnv() // read in environment variables that match

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

// Add flags for commands that carry out transactions
func addTransactionFlags(cmd *cobra.Command, explanation string) {
	cmd.Flags().String("passphrase", "", fmt.Sprintf("passphrase for %s", explanation))
	cmd.Flags().String("privatekey", "", fmt.Sprintf("private key for %s", explanation))
	cmd.Flags().String("gasprice", "", "Gas price for the transaction")
	cmd.Flags().Bool("allowhighgasprice", false, "Allow gas prices higher than 1000GWei")
	cmd.Flags().String("value", "", "Ether to send with the transaction")
	cmd.Flags().Int64("gaslimit", 0, "Gas limit for the transaction; 0 is auto-select")
	cmd.Flags().Int64("nonce", -1, "Nonce for the transaction; -1 is auto-select")
	cmd.Flags().Bool("wait", false, "wait for the transaction to be mined before returning")
	cmd.Flags().Duration("limit", 0, "maximum time to wait for transaction to complete before failing (default forever)")
}

// Obtain the current nonce for the given address
func currentNonce(address common.Address) (uint64, error) {
	var currentNonce uint64
	if nonce == -1 {
		if client == nil {
			err := connect()
			if err != nil {
				return 0, err
			}
		}

		var tmpNonce uint64
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()
		tmpNonce, err = client.PendingNonceAt(ctx, address)
		if err != nil {
			return 0, fmt.Errorf("failed to obtain nonce for %s: %v", address.Hex(), err)
		}
		currentNonce = tmpNonce
		nonce = int64(tmpNonce)
	} else {
		currentNonce = uint64(nonce)
	}
	return currentNonce, nil
}

// Obtain the next nonce for the given address
func nextNonce(address common.Address) (nextNonce uint64, err error) {
	if nonce == -1 {
		_, err = currentNonce(address)
		if err != nil {
			return
		}
	}
	nonce++
	nextNonce = uint64(nonce)
	return
}

// Estimate the gas required for a transaction
func estimateGas(fromAddress common.Address, toAddress *common.Address, amount *big.Int, data []byte) (gas uint64, err error) {
	msg := ethereum.CallMsg{From: fromAddress, To: toAddress, Value: amount, Data: data}
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
	defer cancel()
	gas, err = client.EstimateGas(ctx, msg)
	if err != nil {
		return
	}
	return
}

// Create a transaction
func createTransaction(fromAddress common.Address, toAddress *common.Address, amount *big.Int, gasLimit uint64, data []byte) (tx *types.Transaction, err error) {
	// Obtain the nonce for the transaction
	var txNonce uint64
	txNonce, err = currentNonce(fromAddress)
	if err != nil {
		return
	}

	// Gas limit for the transaction
	if gasLimit == 0 {
		gasLimit, err = estimateGas(fromAddress, toAddress, amount, data)
		if err != nil {
			return
		}
	}

	// Could fetch gas price now that we know the gas limit?

	// Create the transaction
	if toAddress == nil {
		tx = types.NewContractCreation(txNonce, amount, gasLimit, gasPrice, data)
	} else {
		tx = types.NewTransaction(txNonce, *toAddress, amount, gasLimit, gasPrice, data)
	}

	return
}

// Create a signed transaction
func createSignedTransaction(fromAddress common.Address, toAddress *common.Address, amount *big.Int, gasLimit uint64, data []byte) (signedTx *types.Transaction, err error) {
	// Create the transaction
	tx, err := createTransaction(fromAddress, toAddress, amount, gasLimit, data)
	if err != nil {
		return
	}

	// Sign the transaction
	signedTx, err = signTransaction(fromAddress, tx)
	if err != nil {
		err = fmt.Errorf("failed to sign transaction: %v", err)
		return
	}

	// Increment the nonce for the next transaction
	_, err = nextNonce(fromAddress)
	if err != nil {
		return
	}

	return
}

func generateTxOpts(sender common.Address) (*bind.TransactOpts, error) {
	// Signer depends on what information is available to us
	var signer bind.SignerFn
	if viper.GetString("passphrase") != "" {
		wallet, account, err := cli.ObtainWalletAndAccount(chainID, sender)
		if err != nil {
			return nil, err
		}
		signer = util.AccountSigner(chainID, &wallet, account, viper.GetString("passphrase"))
	} else if viper.GetString("privatekey") != "" {
		key, err := crypto.HexToECDSA(strings.TrimPrefix(viper.GetString("privatekey"), "0x"))
		cli.ErrCheck(err, quiet, "Invalid private key")
		signer = util.KeySigner(chainID, key)
	}
	if signer == nil {
		return nil, fmt.Errorf("no signer; please supply either passphrase or private key")
	}

	var value *big.Int
	if viper.GetString("value") != "" {
		value, err = string2eth.StringToWei(viper.GetString("value"))
		cli.ErrCheck(err, quiet, "Failed to understand value")
	}

	curNonce, err := currentNonce(sender)
	if err != nil {
		return nil, err
	}
	nonce = int64(curNonce)

	opts := &bind.TransactOpts{
		From:     sender,
		Signer:   signer,
		GasPrice: gasPrice,
		Value:    value,
		//DoNotSend: offline,
		Nonce: big.NewInt(0).SetInt64(nonce),
	}
	if gasLimit != 0 {
		opts.GasLimit = gasLimit
	}

	return opts, nil
}

func signTransaction(signer common.Address, tx *types.Transaction) (signedTx *types.Transaction, err error) {
	switch {
	case viper.GetString("passphrase") != "":
		if wallet == nil {
			// Fetch the wallet and account for the sender
			wallet, account, err = cli.ObtainWalletAndAccount(chainID, signer)
			if err != nil {
				return
			}
		}
		signedTx, err = wallet.SignTxWithPassphrase(*account, viper.GetString("passphrase"), tx, chainID)
	case viper.GetString("privatekey") != "":
		var key *ecdsa.PrivateKey
		key, err = crypto.HexToECDSA(strings.TrimPrefix(viper.GetString("privatekey"), "0x"))
		cli.ErrCheck(err, quiet, "Invalid private key")
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if signer != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	default:
		err = errors.New("no passphrase or private key; cannot sign")
	}
	return
}

func outputIf(condition bool, msg string) {
	if condition {
		fmt.Println(msg)
	}
}

func localContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
}

func txFrom(tx *types.Transaction) (address common.Address, err error) {
	v, _, _ := tx.RawSignatureValues()
	signer := deriveSigner(v)
	address, err = types.Sender(signer, tx)
	return
}

// Stolen from geth code as this is not exposed
func deriveChainID(v *big.Int) *big.Int {
	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

// Stolen from geth code as this is not exposed
func isProtectedV(v *big.Int) bool {
	if v.BitLen() <= 8 {
		v2 := v.Uint64()
		return v2 != 27 && v2 != 28
	}
	// anything not 27 or 28 are considered unprotected
	return true
}

// Stolen from geth code as this is not exposed
func deriveSigner(v *big.Int) types.Signer {
	if v.Sign() != 0 && isProtectedV(v) {
		return types.NewEIP155Signer(deriveChainID(v))
	}
	return types.HomesteadSigner{}
}
