// Copyright © 2025 Weald Technology Trading.
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

package exit

import (
	"context"
	"strings"
)

func (c *command) output(_ context.Context) (string, error) {
	if c.quiet {
		return "", nil
	}

	builder := strings.Builder{}

	// TODO.
	//	switch c.validatorInfo.Validator.WithdrawalCredentials[0] {
	//	case 0:
	//		builder.WriteString("BLS credentials: ")
	//		builder.WriteString(fmt.Sprintf("%#x", c.validatorInfo.Validator.WithdrawalCredentials))
	//	case 1:
	//		builder.WriteString("Ethereum execution address: ")
	//		builder.WriteString(addressBytesToEIP55(c.validatorInfo.Validator.WithdrawalCredentials[12:]))
	//		if c.verbose {
	//			builder.WriteString("\n")
	//			builder.WriteString("Withdrawal credentials: ")
	//			builder.WriteString(fmt.Sprintf("%#x", c.validatorInfo.Validator.WithdrawalCredentials))
	//		}
	//	}

	return builder.String(), nil
}
