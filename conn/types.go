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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TransactionData contains data to build a transaction.
type TransactionData struct {
	From common.Address
	To   *common.Address

	GasLimit *uint64

	Value *big.Int
	Data  []byte

	Nonce *int64

	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}
