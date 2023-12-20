// Copyright © 2022, 2023 Weald Technology Trading
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
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Conn is a connection to an Ethereum execution client.
type Conn struct {
	timeout time.Duration

	rpcClient *rpc.Client
	client    *ethclient.Client
	// config    *params.ChainConfig

	// nonces tracks per-address nonces.
	nonces   map[common.Address]uint64
	noncesMu sync.Mutex

	// Information for offline connections.
	offline bool
	chainID *big.Int
}

// New creates a new execution client.
func New(ctx context.Context, url string) (*Conn, error) {
	if url == "offline" {
		// We are offline...
		return newOffline(ctx)
	}

	rpcClient, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to RPC client")
	}

	client := ethclient.NewClient(rpcClient)
	if client == nil {
		return nil, errors.New("failed to create client")
	}

	// Fetch chain ID to confirm connection.
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, errors.New("unable to contact client")
	}

	timeout := viper.GetDuration("timeout")
	if timeout == 0 {
		return nil, errors.New("timeout not specified")
	}

	conn := &Conn{
		timeout:   timeout,
		rpcClient: rpcClient,
		client:    client,
		chainID:   chainID,
		nonces:    make(map[common.Address]uint64),
	}

	return conn, nil
}

func newOffline(_ context.Context) (*Conn, error) {
	var chainID *big.Int
	if viper.GetString("network") == "" && viper.GetString("chainid") == "" {
		return nil, errors.New("network or chainid is required when offline")
	}
	switch strings.ToLower(viper.GetString("network")) {
	case "mainnet":
		chainID = params.MainnetChainConfig.ChainID
	case "goerli", "gorli", "görli":
		chainID = params.GoerliChainConfig.ChainID
	case "sepolia":
		chainID = params.SepoliaChainConfig.ChainID
	case "holesky":
		chainID = params.HoleskyChainConfig.ChainID
	default:
		switch {
		case strings.HasPrefix(viper.GetString("chainid"), "0x"):
			// Hex.
			tmp, err := hex.DecodeString(viper.GetString("chainid")[2:])
			if err != nil {
				return nil, errors.Wrap(err, "invalid chain ID")
			}
			chainID = new(big.Int).SetBytes(tmp)
		default:
			// Assume decimal.
			tmp, err := strconv.ParseUint(viper.GetString("chainid"), 10, 64)
			if err != nil {
				return nil, errors.Wrap(err, "invalid chain ID")
			}
			chainID = new(big.Int).SetUint64(tmp)
		}
	}

	return &Conn{
		offline: true,
		chainID: chainID,
		nonces:  make(map[common.Address]uint64),
	}, nil
}

// Client returns the ethclient for the connection.
func (c *Conn) Client() *ethclient.Client {
	return c.client
}

// ChainID returns the chain ID for the connection.
func (c *Conn) ChainID() *big.Int {
	return c.chainID
}
