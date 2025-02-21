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

package util

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/rs/zerolog"
)

// defaultConsensusNodeAddresses are default REST endpoint addresses for consensus nodes.
var defaultConsensusNodeAddresses = []string{
	"localhost:5052", // Lighthouse, Nimbus
	"localhost:5051", // Teku
	"localhost:3500", // Prysm
}

// fallbackConsensusNode is used if no other connection is supplied.
var fallbackConsensusNode = "http://mainnet-consensus.attestant.io/"

type ConnectOpts struct {
	Address       string
	Timeout       time.Duration
	AllowInsecure bool
	LogFallback   bool
}

// ConnectToConsensusNode connects to a consensus node at the given address.
func ConnectToConsensusNode(ctx context.Context, opts *ConnectOpts) (consensusclient.Service, error) {
	if opts == nil {
		return nil, errors.New("no options specified")
	}

	if opts.Timeout == 0 {
		return nil, errors.New("no timeout specified")
	}

	if opts.Address != "" {
		// We have an explicit address; use it.
		return connectToConsensusNode(ctx, opts.Address, opts.Timeout, opts.AllowInsecure)
	}

	// Try the defaults.
	for _, address := range defaultConsensusNodeAddresses {
		client, err := connectToConsensusNode(ctx, address, opts.Timeout, opts.AllowInsecure)
		if err == nil {
			return client, nil
		}
	}

	// The user did not provide a connection, so attempt to use the fallback node.
	if opts.LogFallback {
		fmt.Fprintf(os.Stderr, "No connection supplied with --consensus-connection parameter and no local consensus node found, attempting to use mainnet fallback\n")
	}
	client, err := connectToConsensusNode(ctx, fallbackConsensusNode, opts.Timeout, true)
	if err == nil {
		return client, nil
	}

	return nil, errors.New("failed to connect to any consensus node")
}

func connectToConsensusNode(ctx context.Context, address string, timeout time.Duration, allowInsecure bool) (consensusclient.Service, error) {
	if !strings.HasPrefix(address, "http") {
		address = fmt.Sprintf("http://%s", address)
	}
	if !allowInsecure {
		// Ensure the connection is either secure or local.
		connectionURL, err := url.Parse(address)
		if err != nil {
			return nil, errors.Join(errors.New("failed to parse connection"), err)
		}
		if connectionURL.Scheme == "http" &&
			connectionURL.Host != "localhost" &&
			!strings.HasPrefix(connectionURL.Host, "localhost:") &&
			connectionURL.Host != "127.0.0.1" &&
			!strings.HasPrefix(connectionURL.Host, "127.0.0.1:") {
			fmt.Println("Connections to remote consensus nodes should be secure.  This warning can be silenced with --allow-insecure-connections")
		}
	}
	consensusClient, err := http.New(ctx,
		http.WithLogLevel(zerolog.Disabled),
		http.WithAddress(address),
		http.WithTimeout(timeout),
	)
	if err != nil {
		return nil, errors.Join(errors.New("failed to connect to consensus node"), err)
	}

	return consensusClient, nil
}

// ConsensusValidatorInfo obtains information for a consensus validator.
func ConsensusValidatorInfo(ctx context.Context,
	consensusClient consensusclient.Service,
	validatorID string,
) (
	*apiv1.Validator,
	bool,
	error,
) {
	// Work out if the input is an index or a public key.
	pubkeys := make([]phase0.BLSPubKey, 0)
	indices := make([]phase0.ValidatorIndex, 0)
	if strings.HasPrefix(validatorID, "0x") {
		pubkey := phase0.BLSPubKey{}
		jsonInput := fmt.Sprintf(`"%s"`, validatorID)
		if err := pubkey.UnmarshalJSON([]byte(jsonInput)); err != nil {
			return nil, false, errors.Join(errors.New("invalid validator public key"), err)
		}
		pubkeys = append(pubkeys, pubkey)
	} else {
		index := phase0.ValidatorIndex(0)
		jsonInput := fmt.Sprintf(`"%s"`, validatorID)
		if err := index.UnmarshalJSON([]byte(jsonInput)); err != nil {
			return nil, false, errors.Join(errors.New("invalid validator index"), err)
		}
		indices = append(indices, index)
	}

	validatorsProvider, isValidatorsProvider := consensusClient.(consensusclient.ValidatorsProvider)
	if !isValidatorsProvider {
		return nil, false, errors.New("service is not a validators provider")
	}

	validatorsResponse, err := validatorsProvider.Validators(ctx, &api.ValidatorsOpts{
		State:   "head",
		Indices: indices,
		PubKeys: pubkeys,
	})
	if err != nil {
		// TODO check for 404?
		return nil, false, errors.Join(errors.New("failed to access consensus client"), err)
	}
	validators := validatorsResponse.Data

	if len(validators) == 0 {
		return nil, false, nil
	}

	var res *apiv1.Validator
	for _, v := range validators {
		res = v
	}

	return res, true, nil
}
