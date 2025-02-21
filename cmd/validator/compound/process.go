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

package compound

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	consensusclient "github.com/attestantio/go-eth2-client"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cmd/validator"
	"github.com/wealdtech/ethereal/v2/conn"
	"github.com/wealdtech/ethereal/v2/util"
	"github.com/wealdtech/go-string2eth"
)

func (c *command) process(ctx context.Context) error {
	if c.offline {
		return errors.New("this command cannot run when offline")
	}

	// Obtain information we need to process.
	if err := c.setup(ctx); err != nil {
		return err
	}

	if c.consensusClient == nil && !c.noSafetyChecks {
		return errors.New("no connection to a consensus client provided.  Please provide one with --consensus-connection to allow safety check operations to proceed.  If you really want to proceed without safety checks, provide the --no-safety-checks flag however doing this could lose your Ether.")
	}

	maxFee, err := string2eth.StringToWei(c.maxFee)
	if err != nil {
		return errors.Join(errors.New("invalid maximum fee"), err)
	}
	if maxFee.Sign() != 1 {
		return errors.New("max fee must be a positive value")
	}

	pubkey, err := util.ConsensusPubkey(c.validator)
	if err != nil {
		return errors.Join(errors.New("failed to obtain validator"), err)
	}

	fromAddress, err := c.executionConn.Address(viper.GetString("from"), viper.GetString("privatekey"))
	if err != nil {
		return errors.Join(errors.New("failed to obtain from address"), err)
	}

	if c.consensusClient != nil && !c.noSafetyChecks {
		if err := c.runSafetyChecks(ctx, fromAddress); err != nil {
			return err
		}
	}

	signedTx, err := validator.GenerateConsolidationRequest(ctx, c.executionConn, fromAddress, pubkey, pubkey, maxFee, c.debug)
	if err != nil {
		return err
	}

	err = c.executionConn.SendTransaction(ctx, signedTx)
	if err != nil {
		return errors.Join(errors.New("failed to initiate validator compound"), err)
	}

	c.executionConn.HandleSubmittedTransaction(signedTx, log.Fields{
		"group":   "validator",
		"command": "compound",
	})

	return nil
}

func (c *command) setup(ctx context.Context) error {
	var err error

	// Attempt to connect to the consensus node.
	c.consensusClient, err = util.ConnectToConsensusNode(ctx, &util.ConnectOpts{
		Address:       c.consensusUrl,
		Timeout:       c.timeout,
		AllowInsecure: c.allowInsecureConnections,
		LogFallback:   !c.quiet,
	})
	if err != nil {
		if c.debug {
			fmt.Fprintf(os.Stderr, "failed to connect to consensus node: %v", err)
		}
		c.consensusClient = nil
	} else {
		// Obtain the validators provider.
		var isProvider bool
		c.validatorsProvider, isProvider = c.consensusClient.(consensusclient.ValidatorsProvider)
		if !isProvider {
			return errors.New("consensus node does not provide validator information")
		}
	}

	c.executionConn, err = conn.New(ctx, c.executionUrl, c.debug)
	if err != nil {
		return errors.Join(errors.New("failed to connect to execution node"), err)
	}

	return nil
}

func (c *command) runSafetyChecks(ctx context.Context,
	from common.Address,
) error {
	validatorInfo, found, err := util.ConsensusValidatorInfo(ctx, c.consensusClient, c.validator)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("validator is not known")
	}

	if validatorInfo.Validator == nil {
		return errors.New("validator is not yet active")
	}

	if validatorInfo.Status != apiv1.ValidatorStateActiveOngoing {
		return fmt.Errorf("validator is in state %s; must be in state %s to compound",
			validatorInfo.Status,
			apiv1.ValidatorStateActiveOngoing,
		)
	}

	if validatorInfo.Validator.WithdrawalCredentials[0] != byte(1) {
		return fmt.Errorf("validator has type %d credentials; can only compound a validator with type 1 credentials",
			int(validatorInfo.Validator.WithdrawalCredentials[0]),
		)
	}

	if !bytes.Equal(validatorInfo.Validator.WithdrawalCredentials[12:], from.Bytes()) {
		return fmt.Errorf("transaction sender address %s does not match that of the validator's withdrawal credentials %s",
			common.BytesToAddress(from.Bytes()).String(),
			common.BytesToAddress(validatorInfo.Validator.WithdrawalCredentials[12:]).String(),
		)
	}

	return nil
}
