// Copyright © 2019 Weald Technology Trading
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

package util

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

// WaitForTransaction waits for the transaction to be mined, or for the limit to expire.
func WaitForTransaction(client *ethclient.Client, txHash common.Hash, limit time.Duration) bool {
	start := time.Now()
	first := true
	for limit == 0 || time.Since(start) < limit {
		if !first {
			time.Sleep(5 * time.Second)
		} else {
			first = false
		}
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()
		_, pending, err := client.TransactionByHash(ctx, txHash)
		if err == nil && !pending {
			return true
		}
	}
	return false
}
