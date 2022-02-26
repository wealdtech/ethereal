// Copyright © 2022 Weald Technology Trading.
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

package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wealdtech/ethereal/v2/util"
)

func TestDepositInfoFromJSON(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name:  "Test",
			input: []byte(`[{"pubkey": "aad67d87ddeb2801860c135a67dc3fecdf77ed9a41da6afe7c8a5232354713bdc6d437cbe0014f3482f2a17e048e30a4", "withdrawal_credentials": "0070f9cba5c36591736e62d2a4c32bdfdecb92ea586e9cdb89d95788ce7f4975", "amount": 32000000000, "signature": "a8e83f7a0c36a4aa45906aa45039e39212b9cbd3916550adaeac488a847e216ab8cf1d9360608dd0a092b4a1ced05f2c05b5d8406c40410933ee6ccecff4e31eac088383a815b6cd8d17fa0d87586a0f9fe9f01a4d7bb9aa591851baff1dae13", "deposit_message_root": "b082661eaebf92daf5f0b08728832305cc309467642354508206cd4f09150a1a", "deposit_data_root": "2c880f13079bbae7ad9a15bad96a309730a032c497f427cb271e3435947dc646", "fork_version": "00000113", "deposit_cli_version": "1.0.0"}, {"pubkey": "94c270cb9c846a6da5619c100c674a70c986b022f0f653ca35b4c2ad849a2a6d47143a9da38e256b9578f622f8e8b851", "withdrawal_credentials": "0063f9c5c728f07a4490b3ad125d4939655fb07517602939ba181197bc525a62", "amount": 32000000000, "signature": "a981d6979692b6c2d6282de0fbf08e15485d1daa471d0a21682d3284916ee9d9defacc851a49be7e17f0d74a2ae196990536f568effff25d38fe28282b7d25e6ecf562cff5246cd023e0b42168126043a80c47a1f6a7089d406430d868c8c580", "deposit_message_root": "693e0a955b70b405f6f79d95df75f3369aa1d8455c7af93a6dff5bb4399f2967", "deposit_data_root": "5830db4b95cbcc8ff77cfb473146462a8f68e182c8ed725ad9087d61ea13f7bc", "fork_version": "00000113", "deposit_cli_version": "1.0.0"}, {"pubkey": "8b446d4ea379dfbc4d7af1e9fa79e57c6a1014aea10ad9afbdc6bc9e45792f80b859040b2db3c65e53e50d31d8f9ea1e", "withdrawal_credentials": "0047babb6b3d99ceaa8c117eef7f6be457872e65daed754db0de257b4795b507", "amount": 32000000000, "signature": "871cb2e25dce95dba3f1727f13f6f2aca1c358eac86037fdd9c5abc90986f650ad760e4f46540bea090ce8f5f4a1ba010892f9304f4e5cf2ce232c14a2f4cc083da84ce4f9c794f871e2a25574e9a85f847ab88bf8078f1b770666cea85bc4f0", "deposit_message_root": "e8dc7afcad5258bfd6383afe40de006ec32bb7add6cf70a180bc456935b61943", "deposit_data_root": "50f915bb6545e183f18df835b3472b5d35ee438b49fcdedb0c7e34ecde108ddf", "fork_version": "00000113", "deposit_cli_version": "1.0.0"}]`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := util.DepositInfoFromJSON(test.input)
			if test.err == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, test.err)
			}
		})
	}
}
