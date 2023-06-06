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

package funcparser

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/ethereal/v2/util"
	"github.com/wealdtech/ethereal/v2/util/funcparser/parser"
)

// ParseCall parses a call string and returns a suitable Method.
func ParseCall(client *ethclient.Client, contract *util.Contract, call string) (*abi.Method, []interface{}, error) {
	if contract == nil {
		return nil, nil, errors.New("no contract")
	}

	is := antlr.NewInputStream(call)
	lexer := parser.NewFuncLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	tree := parser.NewFuncParser(stream).Start_()
	methodListener := newMethodListener(client, contract)
	antlr.ParseTreeWalkerDefault.Walk(methodListener, tree)

	return methodListener.method, methodListener.args, methodListener.err
}
