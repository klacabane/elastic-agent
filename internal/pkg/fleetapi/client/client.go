// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/elastic/elastic-agent/internal/pkg/agent/errors"
	"github.com/elastic/elastic-agent/internal/pkg/core/logger"
	"github.com/elastic/elastic-agent/internal/pkg/release"
	"github.com/elastic/elastic-agent/internal/pkg/remote"
)

// Sender is an sender interface describing client behavior.
type Sender interface {
	Send(
		ctx context.Context,
		method string,
		path string,
		params url.Values,
		headers http.Header,
		body io.Reader,
	) (*http.Response, error)

	URI() string
}

var baseRoundTrippers = func(rt http.RoundTripper) (http.RoundTripper, error) {
	rt = NewFleetUserAgentRoundTripper(rt, release.Version())
	return rt, nil
}

func init() {
	val, ok := os.LookupEnv("DEBUG_AGENT")
	if ok && val == "1" {
		fn := baseRoundTrippers
		baseRoundTrippers = func(rt http.RoundTripper) (http.RoundTripper, error) {
			rt, err := fn(rt)
			if err != nil {
				return nil, err
			}

			l, err := logger.New("fleet_client", false)
			if err != nil {
				return nil, errors.New(err, "could not create the logger for debugging HTTP request")
			}

			return remote.NewDebugRoundTripper(rt, l), nil
		}
	}
}

// NewAuthWithConfig returns a fleet-server client that will:
//
// - Send the API Key on every HTTP request.
// - Ensure a minimun version of fleet-server is required.
// - Send the Fleet User Agent on every HTTP request.
func NewAuthWithConfig(log *logger.Logger, apiKey string, cfg remote.Config) (*remote.Client, error) {
	return remote.NewWithConfig(log, cfg, func(rt http.RoundTripper) (http.RoundTripper, error) {
		rt, err := baseRoundTrippers(rt)
		if err != nil {
			return nil, err
		}

		rt, err = NewFleetAuthRoundTripper(rt, apiKey)
		if err != nil {
			return nil, err
		}

		return rt, nil
	})
}

// NewWithConfig takes a fleet-server configuration and create a remote.client with the appropriate tripper.
func NewWithConfig(log *logger.Logger, cfg remote.Config) (*remote.Client, error) {
	return remote.NewWithConfig(log, cfg, baseRoundTrippers)
}

// ExtractError extracts error from a fleet-server response
func ExtractError(resp io.Reader) error {
	// Lets try to extract a high level fleet-server error.
	e := &struct {
		StatusCode int    `json:"statusCode"`
		Error      string `json:"error"`
		Message    string `json:"message"`
	}{}

	data, err := ioutil.ReadAll(resp)
	if err != nil {
		return errors.New(err, "fail to read original error")
	}

	err = json.Unmarshal(data, e)
	if err == nil {
		// System errors doesn't return a message, fleet code can return a Message key which has more
		// information.
		if len(e.Message) == 0 {
			return fmt.Errorf("status code: %d, fleet-server returned an error: %s", e.StatusCode, e.Error)
		}
		return fmt.Errorf(
			"status code: %d, fleet-server returned an error: %s, message: %s",
			e.StatusCode,
			e.Error,
			e.Message,
		)
	}

	return fmt.Errorf("could not decode the response, raw response: %s", string(data))
}
