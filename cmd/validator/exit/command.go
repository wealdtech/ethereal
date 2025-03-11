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
	"time"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/conn"
	"github.com/wealdtech/ethereal/v2/services/chaintime"
)

type command struct {
	offline        bool
	quiet          bool
	verbose        bool
	debug          bool
	noSafetyChecks bool

	// Input.
	validator string
	maxFee    string

	// Connections.
	timeout                  time.Duration
	allowInsecureConnections bool
	consensusURL             string
	executionURL             string

	// Data access.
	consensusClient      consensusclient.Service
	validatorsProvider   consensusclient.ValidatorsProvider
	executionConn        *conn.Conn
	chainTime            chaintime.Service
	shardCommitteePeriod uint64
}

func newCommand(_ context.Context) (*command, error) {
	c := &command{
		offline:                  viper.GetBool("offline"),
		quiet:                    viper.GetBool("quiet"),
		verbose:                  viper.GetBool("verbose"),
		debug:                    viper.GetBool("debug"),
		noSafetyChecks:           viper.GetBool("no-safety-checks"),
		timeout:                  viper.GetDuration("timeout"),
		consensusURL:             viper.GetString("consensus-connection"),
		executionURL:             viper.GetString("connection"),
		allowInsecureConnections: viper.GetBool("allow-insecure-connections"),
		validator:                viper.GetString("validator"),
		maxFee:                   viper.GetString("max-fee"),
	}

	// Timeout.
	if c.timeout == 0 {
		return nil, errors.New("timeout is required")
	}

	if c.validator == "" {
		return nil, errors.New("validator is required")
	}

	return c, nil
}
