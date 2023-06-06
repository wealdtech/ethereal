// Copyright © 2022 Weald Technology Trading
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

package conn

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
)

// SignTransaction signs the given transaction, returning a signed transaction.
func (c *Conn) SignTransaction(_ context.Context,
	signer common.Address,
	tx *types.Transaction,
) (
	*types.Transaction,
	error,
) {
	var signedTx *types.Transaction
	switch {
	case viper.GetString("passphrase") != "":
		wallet, account, err := cli.ObtainWalletAndAccount(c.ChainID(), signer)
		if err != nil {
			return nil, err
		}
		signedTx, err = wallet.SignTxWithPassphrase(*account, viper.GetString("passphrase"), tx, c.ChainID())
		if err != nil {
			return nil, err
		}
	case viper.GetString("privatekey") != "":
		key, err := crypto.HexToECDSA(strings.TrimPrefix(viper.GetString("privatekey"), "0x"))
		if err != nil {
			return nil, errors.Wrap(err, "invalid private key")
		}
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if signer != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		signedTx, err = types.SignTx(tx, types.NewLondonSigner(c.ChainID()), key)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("no passphrase or private key; cannot sign")
	}
	return signedTx, nil
}
