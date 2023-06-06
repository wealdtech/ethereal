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
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	"github.com/wealdtech/ethereal/v2/util"
	"github.com/wealdtech/ethereal/v2/util/funcparser"
)

var (
	tokenDeployName     string
	tokenDeploySymbol   string
	tokenDeployDecimals uint16
	tokenDeploySupply   uint64
	tokenDeployOwner    string
)

var erc20ContractData = `{"contracts":{"ERC20Token.sol:ERC20Token":{"abi":"[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]","bin":"608060405234801561001057600080fd5b506040516108b43803806108b48339818101604052608081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815164010000000081118282018710171561007857600080fd5b5050929190602001805164010000000081111561009457600080fd5b820160208101848111156100a757600080fd5b81516401000000008111828201871017156100c157600080fd5b505060208201516040909201519093509091508061014057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4d7573742068617665206120737570706c79206f6620746f6b656e7300000000604482015290519081900360640190fd5b83516101539060009060208701906101cc565b5082516101679060019060208601906101cc565b506002805460ff191660ff84161790556003819055336000818152600460209081526040808320859055805185815290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350505050610267565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061020d57805160ff191683800117855561023a565b8280016001018555821561023a579182015b8281111561023a57825182559160200191906001019061021f565b5061024692915061024a565b5090565b61026491905b808211156102465760008155600101610250565b90565b61063e806102766000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063313ce56711610066578063313ce567146101a557806370a08231146101c357806395d89b41146101e9578063a9059cbb146101f1578063dd62ed3e1461021d57610093565b806306fdde0314610098578063095ea7b31461011557806318160ddd1461015557806323b872dd1461016f575b600080fd5b6100a061024b565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100da5781810151838201526020016100c2565b50505050905090810190601f1680156101075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101416004803603604081101561012b57600080fd5b506001600160a01b0381351690602001356102d9565b604080519115158252519081900360200190f35b61015d61033e565b60408051918252519081900360200190f35b6101416004803603606081101561018557600080fd5b506001600160a01b03813581169160208101359091169060400135610344565b6101ad610499565b6040805160ff9092168252519081900360200190f35b61015d600480360360208110156101d957600080fd5b50356001600160a01b03166104a2565b6100a06104bd565b6101416004803603604081101561020757600080fd5b506001600160a01b038135169060200135610517565b61015d6004803603604081101561023357600080fd5b506001600160a01b03813581169160200135166105dd565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102d15780601f106102a6576101008083540402835291602001916102d1565b820191906000526020600020905b8154815290600101906020018083116102b457829003601f168201915b505050505081565b6001600160a01b0382166000818152600560209081526040808320338085529083528184208690558151868152915193949390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035481565b6001600160a01b0383166000908152600460205260408120548211156103a5576040805162461bcd60e51b81526020600482015260116024820152704e6f7420656e6f75676820746f6b656e7360781b604482015290519081900360640190fd5b3360009081526005602090815260408083206001600160a01b0388168452909152902054821115610414576040805162461bcd60e51b81526020600482015260146024820152734e6f7420656e6f75676820616c6c6f77616e636560601b604482015290519081900360640190fd5b3360009081526005602090815260408083206001600160a01b038881168086529184528285208054889003905560048452828520805488900390558716808552938290208054870190558151868152915190927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef928290030190a35060019392505050565b60025460ff1681565b6001600160a01b031660009081526004602052604090205490565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102d15780601f106102a6576101008083540402835291602001916102d1565b3360009081526004602052604081205482111561056f576040805162461bcd60e51b81526020600482015260116024820152704e6f7420656e6f75676820746f6b656e7360781b604482015290519081900360640190fd5b336000818152600460209081526040808320805487900390556001600160a01b03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600192915050565b6001600160a01b038082166000908152600560209081526040808320938616835292905220549291505056fea265627a7a723058201d2ef62b40755b34ac4fd0b4e9817a06349244480de5a62bea6e1dddd9ee773e64736f6c634300050a0032"}},"version":"0.5.10+commit.5a6ea5b1.Linux.g++"}`

// tokenDeployCmd represents the token deploy command.
var tokenDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a standard token contract",
	Long: `Deploy a token contract adhering to one or more standards.  For example:

    ethereal token deploy --name="My token" --symbol="MY" --decimals=18 --totalsupply=1000000 --owner=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(tokenDeployName != "", quiet, "--name is required")
		cli.Assert(tokenDeploySymbol != "", quiet, "--symbol is required")
		cli.Assert(tokenDeploySupply != 0, quiet, "--supply is required")

		supply := new(big.Int).SetUint64(tokenDeploySupply)
		supply.Mul(supply, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDeployDecimals)), nil))

		cli.Assert(tokenDeployOwner != "", quiet, "--owner is required")
		owner, err := c.Resolve(tokenDeployOwner)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve owner address %s", tokenDeployOwner))

		contract, err := util.ParseCombinedJSON(erc20ContractData, "ERC20Token")
		cli.ErrCheck(err, quiet, "failed to parse ERC-20 contract data")

		// Set up the constructor.
		constructor := fmt.Sprintf("constructor(%q,%q,%v,%v)", tokenDeployName, tokenDeploySymbol, tokenDeployDecimals, supply)
		_, constructorArgs, err := funcparser.ParseCall(c.Client(), contract, constructor)
		cli.ErrCheck(err, quiet, "Failed to parse constructor")
		argData, err := contract.Abi.Pack("", constructorArgs...)
		cli.ErrCheck(err, quiet, "Failed to convert arguments")
		contract.Binary = append(contract.Binary, argData...)

		var gasLimit *uint64
		limit := uint64(viper.GetInt64("gaslimit"))
		if limit > 0 {
			gasLimit = &limit
		}

		// Deploy the token contract.
		signedTx, err := c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
			From:     owner,
			GasLimit: gasLimit,
			Data:     contract.Binary,
		})
		cli.ErrCheck(err, quiet, "Failed to create token contract deployment transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			err = c.SendTransaction(context.Background(), signedTx)
			cli.ErrCheck(err, quiet, "Failed to send transaction")
			handleSubmittedTransaction(signedTx, log.Fields{
				"group":    "token",
				"command":  "deploy",
				"name":     tokenDeployName,
				"symbol":   tokenDeploySymbol,
				"decimals": tokenDeployDecimals,
				"supply":   tokenDeploySupply,
			}, true)
		}
	},
}

func init() {
	tokenCmd.AddCommand(tokenDeployCmd)
	tokenFlags(tokenDeployCmd)
	tokenDeployCmd.Flags().StringVar(&tokenDeployName, "name", "", "Name for the token")
	tokenDeployCmd.Flags().StringVar(&tokenDeploySymbol, "symbol", "", "Symbol for the token")
	tokenDeployCmd.Flags().Uint16Var(&tokenDeployDecimals, "decimals", 18, "Decimals for the token")
	tokenDeployCmd.Flags().Uint64Var(&tokenDeploySupply, "supply", 0, "Total supply for the token (in whole tokens)")
	tokenDeployCmd.Flags().StringVar(&tokenDeployOwner, "owner", "", "Address that owns the initial tokens")
	addTransactionFlags(tokenDeployCmd, "the address from which to deploy the token contract")
}
