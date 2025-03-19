// Copyright © 2025 Weald Technology Trading
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

package validator

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBumpFee(t *testing.T) {
	tests := []struct {
		name    string
		initial *big.Int
		res     *big.Int
	}{
		{
			name:    "Low",
			initial: big.NewInt(1),
			res:     big.NewInt(1000001),
		},
		{
			name:    "Mid",
			initial: big.NewInt(1e9),
			res:     big.NewInt(15e8),
		},
		{
			name:    "High",
			initial: big.NewInt(1e15),
			res:     big.NewInt(1001e12),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := bumpFee(test.initial)
			require.Equal(t, test.res, res)
		})
	}
}
