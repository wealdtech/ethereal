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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wealdtech/go-string2eth"
)

// CurrentBaseFee returns the current base fee of the chain.
func (c *Conn) CurrentBaseFee(ctx context.Context) (*big.Int, error) {
	// If we have been supplied with a base fee then use it.
	if viper.GetString("base-fee-per-gas") != "" {
		baseFee, err := string2eth.StringToWei(viper.GetString("base-fee-per-gas"))
		if err != nil {
			return nil, err
		}
		return baseFee, nil
	}

	// If we're offline we cannot go any further.
	if c.client == nil {
		return nil, errors.New("no client connection; please supply base fee with base-fee-per-gas option")
	}

	// Obtain the base fee from the current block.
	blockNum, err := c.client.BlockNumber(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain current block number")
	}
	block, err := c.client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to obtain block %d", blockNum))
	}
	baseFee := eip1559.CalcBaseFee(&params.ChainConfig{
		LondonBlock: big.NewInt(0),
	}, block.Header())

	return baseFee, nil
}
