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

package topup

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cmd/validator"
	"github.com/wealdtech/ethereal/v2/conn"
	standardchaintime "github.com/wealdtech/ethereal/v2/services/chaintime/standard"
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
		return errors.New("no connection to a consensus client provided.  Please provide one with --consensus-connection to allow safety check operations to proceed.  If you really want to proceed without safety checks, provide the --no-safety-checks flag however doing this could lose your Ether")
	}

	consensusValidator, found, err := util.ConsensusValidatorInfo(ctx, c.consensusClient, c.validator)
	if err != nil {
		return errors.Join(errors.New("failed to obtain validator"), err)
	}
	if !found {
		return errors.New("validator is not known")
	}
	if consensusValidator.Validator == nil {
		return errors.New("validator is not yet active")
	}
	if consensusValidator.Validator.PublicKey.IsZero() {
		return errors.New("validator public key must be provided or obtained from the consensus chain to create a topup request")
	}

	amount, err := string2eth.StringToWei(c.topupAmount)
	if err != nil {
		return errors.Join(errors.New("invalid topup amount"), err)
	}
	if amount.Sign() <= 0 {
		return errors.New("topup amount must be more than zero")
	}
	if new(big.Int).Mod(amount, big.NewInt(1e18)).Sign() != 0 {
		return errors.New("topup amount must be in whole ether")
	}

	fromAddress, err := c.executionConn.Address(viper.GetString("from"), viper.GetString("privatekey"))
	if err != nil {
		return errors.Join(errors.New("failed to obtain from address"), err)
	}

	if c.consensusClient != nil && !c.noSafetyChecks {
		if err := c.runSafetyChecks(ctx, fromAddress, amount); err != nil {
			return err
		}
	}

	signedTx, err := validator.GenerateTopupRequest(ctx, c.executionConn, c.consensusClient, fromAddress, consensusValidator.Validator.PublicKey, amount, c.debug)
	if err != nil {
		return err
	}

	err = c.executionConn.SendTransaction(ctx, signedTx)
	if err != nil {
		return errors.Join(errors.New("failed to initiate validator withdrawal"), err)
	}

	c.executionConn.HandleSubmittedTransaction(signedTx, log.Fields{
		"group":   "validator",
		"command": "topup",
	})

	return nil
}

func (c *command) setup(ctx context.Context) error {
	var err error

	// Attempt to connect to the consensus node.
	c.consensusClient, err = util.ConnectToConsensusNode(ctx, &util.ConnectOpts{
		Address:       c.consensusURL,
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
		var isProvider bool
		c.validatorsProvider, isProvider = c.consensusClient.(consensusclient.ValidatorsProvider)
		if !isProvider {
			return errors.New("consensus node does not provide validator information")
		}

		specProvider, isProvider := c.consensusClient.(consensusclient.SpecProvider)
		if !isProvider {
			return errors.New("consensus node does not provide spec information")
		}
		genesisProvider, isProvider := c.consensusClient.(consensusclient.GenesisProvider)
		if !isProvider {
			return errors.New("consensus node does not provide genesis information")
		}
		c.chainTime, err = standardchaintime.New(ctx,
			standardchaintime.WithLogLevel(zerolog.Disabled),
			standardchaintime.WithSpecProvider(specProvider),
			standardchaintime.WithGenesisProvider(genesisProvider),
		)
		if err != nil {
			return errors.Join(errors.New("failed to access chaintim service"), err)
		}

		specResponse, err := specProvider.Spec(ctx, &api.SpecOpts{})
		if err != nil {
			return errors.Join(errors.New("failed to obtain spec"), err)
		}
		c.shardCommitteePeriod = specResponse.Data["SHARD_COMMITTEE_PERIOD"].(uint64)
		c.minValidatorWithdrawabilityDelay = specResponse.Data["MIN_VALIDATOR_WITHDRAWABILITY_DELAY"].(uint64)
	}

	c.executionConn, err = conn.New(ctx, c.executionURL, c.debug, c.quiet)
	if err != nil {
		return errors.Join(errors.New("failed to connect to execution node"), err)
	}

	return nil
}

func (c *command) runSafetyChecks(ctx context.Context,
	from common.Address,
	amount *big.Int,
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
		return fmt.Errorf("validator is in state %s; must be in state %s to topup",
			validatorInfo.Status,
			apiv1.ValidatorStateActiveOngoing,
		)
	}

	if validatorInfo.Validator.WithdrawalCredentials[0] == byte(0) {
		return errors.New("topup of a validator with type 0 credentials is usually non-sensical.  If you really want to do this then rerun the command with --no-safety-checks")
	}

	gweiTopupAmount := phase0.Gwei(new(big.Int).Div(amount, big.NewInt(1e9)).Uint64())
	newBalance := validatorInfo.Balance + gweiTopupAmount
	if validatorInfo.Validator.WithdrawalCredentials[0] == byte(1) {
		if newBalance > 32250000000 {
			return fmt.Errorf("topup will result in excess Ether in the validator.  If you really want to do this then rerun the command with --no-safety-checks")
		}
	}
	if validatorInfo.Validator.WithdrawalCredentials[0] == byte(2) {
		if newBalance > 2048250000000 {
			return fmt.Errorf("topup will result in excess Ether in the validator.  If you really want to do this then rerun the command with --no-safety-checks")
		}
	}

	if !bytes.Equal(validatorInfo.Validator.WithdrawalCredentials[12:], from.Bytes()) {
		return fmt.Errorf("transaction sender address %s does not match that of the validator's withdrawal credentials %s.  If you really want to send from this address then rerun the command with --no-safety-checks, however if you are sending to a validator to which you do not control the withdrawal credentials you may lose the Ether",
			common.BytesToAddress(from.Bytes()).String(),
			common.BytesToAddress(validatorInfo.Validator.WithdrawalCredentials[12:]).String(),
		)
	}

	return nil
}
