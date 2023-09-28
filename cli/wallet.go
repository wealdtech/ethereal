// Copyright 2017 - 2023 Weald Technology Trading Limited
//
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

package cli

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// ObtainWallets obtains all known wallets for a given chain.
func ObtainWallets(chainID *big.Int, debug bool) ([]accounts.Wallet, error) {
	var wallets []accounts.Wallet

	gethWallets, err := obtainGethWallets(chainID, debug)
	if err != nil {
		return nil, err
	}
	wallets = append(wallets, gethWallets...)

	parityWallets, err := obtainParityWallets(debug)
	if err != nil {
		return nil, err
	}
	wallets = append(wallets, parityWallets...)

	ledgerWallets, err := obtainLedgerWallets(debug)
	if err != nil {
		return nil, err
	}
	wallets = append(wallets, ledgerWallets...)

	return wallets, nil
}

// ObtainWalletAndAccount obtains the wallet and account for an address.
func ObtainWalletAndAccount(chainID *big.Int, address common.Address) (accounts.Wallet, *accounts.Account, error) {
	var account *accounts.Account
	wallet, err := ObtainWallet(chainID, address)
	if err == nil {
		account, err = ObtainAccount(&wallet, &address, viper.GetString("passphrase"))
	}
	return wallet, account, err
}

// ObtainWallet fetches the wallet for a given address.
func ObtainWallet(chainID *big.Int, address common.Address) (accounts.Wallet, error) {
	wallet, err := obtainGethWallet(chainID, address)
	if err == nil {
		return wallet, nil
	}

	wallet, err = obtainParityWallet(address)
	if err == nil {
		return wallet, nil
	}

	return wallet, fmt.Errorf("failed to obtain wallet for %s", address.Hex())
}

func obtainGethWallet(chainID *big.Int, address common.Address) (accounts.Wallet, error) {
	keydir := DefaultDataDir()
	switch {
	case chainID.Cmp(params.MainnetChainConfig.ChainID) == 0:
		// Nothing to add for mainnet.
	case chainID.Cmp(params.GoerliChainConfig.ChainID) == 0:
		keydir = filepath.Join(keydir, "goerli")
	case chainID.Cmp(params.SepoliaChainConfig.ChainID) == 0:
		keydir = filepath.Join(keydir, "sepolia")
	case chainID.Cmp(params.HoleskyChainConfig.ChainID) == 0:
		keydir = filepath.Join(keydir, "holesky")
	}
	keydir = filepath.Join(keydir, "keystore")
	backends := []accounts.Backend{keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)}
	accountManager := accounts.NewManager(nil, backends...)
	defer accountManager.Close()
	account := accounts.Account{Address: address}
	wallet, err := accountManager.Find(account)
	return wallet, err
}

func obtainGethWallets(chainID *big.Int, debug bool) ([]accounts.Wallet, error) {
	keydir := DefaultDataDir()
	switch {
	case chainID.Cmp(params.MainnetChainConfig.ChainID) == 0:
		// Nothing to add for mainnet.
	case chainID.Cmp(params.GoerliChainConfig.ChainID) == 0:
		keydir = filepath.Join(keydir, "goerli")
	case chainID.Cmp(params.SepoliaChainConfig.ChainID) == 0:
		keydir = filepath.Join(keydir, "sepolia")
	case chainID.Cmp(params.HoleskyChainConfig.ChainID) == 0:
		keydir = filepath.Join(keydir, "holesky")
	}
	keydir = filepath.Join(keydir, "keystore")
	if debug {
		fmt.Printf("Geth key directory is %s\n", keydir)
	}
	backends := []accounts.Backend{keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)}
	accountManager := accounts.NewManager(nil, backends...)
	defer accountManager.Close()
	return accountManager.Wallets(), nil
}

func obtainParityWallet(address common.Address) (accounts.Wallet, error) {
	keydir, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("failed to find home directory")
	}
	switch runtime.GOOS {
	case "windows":
		keydir = filepath.Join(keydir, "AppData\\Roaming\\Parity\\Ethereum\\keys\\ethereum")
	case "darwin":
		keydir = filepath.Join(keydir, "Library/Application Support/io.parity.ethereum/keys/ethereum")
	case "linux":
		keydir = filepath.Join(keydir, ".local/share/io.parity.ethereum/keys/ethereum")
	default:
		return nil, fmt.Errorf("unsupported operating system")
	}

	backends := []accounts.Backend{keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)}
	accountManager := accounts.NewManager(nil, backends...)
	defer accountManager.Close()
	account := accounts.Account{Address: address}
	wallet, err := accountManager.Find(account)
	return wallet, err
}

func obtainParityWallets(debug bool) ([]accounts.Wallet, error) {
	keydir, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("failed to find home directory")
	}
	switch runtime.GOOS {
	case "windows":
		keydir = filepath.Join(keydir, "AppData\\Roaming\\Parity\\Ethereum\\keys\\ethereum")
	case "darwin":
		keydir = filepath.Join(keydir, "Library/Application Support/io.parity.ethereum/keys/ethereum")
	case "linux":
		keydir = filepath.Join(keydir, ".local/share/io.parity.ethereum/keys/ethereum")
	default:
		return nil, fmt.Errorf("unsupported operating system")
	}

	if debug {
		fmt.Printf("Parity key directory is %s\n", keydir)
	}

	backends := []accounts.Backend{keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)}
	accountManager := accounts.NewManager(nil, backends...)
	defer accountManager.Close()
	return accountManager.Wallets(), nil
}

func obtainLedgerWallets(_ bool) ([]accounts.Wallet, error) {
	ledgerhub, err := usbwallet.NewLedgerHub()
	if err != nil {
		return nil, err
	}

	backends := []accounts.Backend{ledgerhub}
	accountManager := accounts.NewManager(nil, backends...)
	defer accountManager.Close()

	usbWallets := viper.GetInt("usbwallets")
	for _, wallet := range accountManager.Wallets() {
		err = wallet.Open("")
		if err != nil {
			continue
		}
		path := accounts.LegacyLedgerBaseDerivationPath
		for i := 0; i < usbWallets; i++ {
			path[3] = uint32(i)
			// nolint:errcheck
			wallet.Derive(path, true)
		}
	}

	return accountManager.Wallets(), nil
}

// ObtainAccount fetches the account for a given address.
func ObtainAccount(wallet *accounts.Wallet, address *common.Address, passphrase string) (*accounts.Account, error) {
	for _, account := range (*wallet).Accounts() {
		if *address == account.Address {
			if passphrase != "" && !VerifyPassphrase(*wallet, account, passphrase) {
				return nil, errors.New("invalid passphrase")
			}
			return &account, nil
		}
	}
	return nil, errors.New("account not found")
}

// VerifyPassphrase confirms that a passphrase is correct for an account.
func VerifyPassphrase(wallet accounts.Wallet, account accounts.Account, passphrase string) bool {
	_, err := wallet.SignDataWithPassphrase(account, passphrase, "application/octet-stream", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return err == nil
}

// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir.
	home := homeDir()
	if home != "" {
		switch runtime.GOOS {
		case "darwin":
			return filepath.Join(home, "Library", "Ethereum")
		case "windows":
			// We used to put everything in %HOME%\AppData\Roaming, but this caused
			// problems with non-typical setups. If this fallback location exists and
			// is non-empty, use it, otherwise DTRT and check %LOCALAPPDATA%.
			fallback := filepath.Join(home, "AppData", "Roaming", "Ethereum")
			appdata := windowsAppData()
			if appdata == "" || isNonEmptyDir(fallback) {
				return fallback
			}
			return filepath.Join(appdata, "Ethereum")
		default:
			return filepath.Join(home, ".ethereum")
		}
	}
	// As we cannot guess a stable location, return empty and handle later.
	return ""
}

func windowsAppData() string {
	v := os.Getenv("LOCALAPPDATA")
	if v == "" {
		// Windows XP and below don't have LocalAppData. Crash here because
		// we don't support Windows XP and undefining the variable will cause
		// other issues.
		panic("environment variable LocalAppData is undefined")
	}
	return v
}

func isNonEmptyDir(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return false
	}
	names, _ := f.Readdir(1)
	f.Close()
	return len(names) > 0
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
