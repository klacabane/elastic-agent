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

package beats

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"

	"github.com/elastic/elastic-agent/internal/pkg/agent/application/paths"
	"github.com/elastic/elastic-agent/internal/pkg/agent/program"
	monitoringConfig "github.com/elastic/elastic-agent/internal/pkg/core/monitoring/config"
)

const (
	// args: data path, pipeline name, application name
	logFileFormat = "%s/logs/%s/%s"
	// args: data path, install path, pipeline name, application name
	logFileFormatWin = "%s\\logs\\%s\\%s"

	// args: pipeline name, application name
	mbEndpointFileFormatWin = `npipe:///%s-%s`

	// args: pipeline name, application name
	agentMbEndpointFileFormatWin = `npipe:///elastic-agent`
	// agentMbEndpointHTTP is used with cloud and exposes metrics on http endpoint
	agentMbEndpointHTTP = "http://%s:%d"
)

// MonitoringEndpoint is an endpoint where process is exposing its metrics.
func MonitoringEndpoint(spec program.Spec, operatingSystem, pipelineID string) string {
	if endpoint, ok := spec.MetricEndpoints[operatingSystem]; ok {
		return endpoint
	}
	if operatingSystem == "windows" {
		return fmt.Sprintf(mbEndpointFileFormatWin, pipelineID, spec.Cmd)
	}
	// unix socket path must be less than 104 characters
	path := fmt.Sprintf("unix://%s.sock", filepath.Join(paths.TempDir(), pipelineID, spec.Cmd, spec.Cmd))
	if len(path) < 104 {
		return path
	}
	// place in global /tmp (or /var/tmp on Darwin) to ensure that its small enough to fit; current path is way to long
	// for it to be used, but needs to be unique per Agent (in the case that multiple are running)
	return fmt.Sprintf(`unix:///tmp/elastic-agent/%x.sock`, sha256.Sum256([]byte(path)))
}

func getLoggingFile(spec program.Spec, operatingSystem, installPath, pipelineID string) string {
	if path, ok := spec.LogPaths[operatingSystem]; ok {
		return path
	}
	if operatingSystem == "windows" {
		return fmt.Sprintf(logFileFormatWin, paths.Home(), pipelineID, spec.Cmd)
	}
	return fmt.Sprintf(logFileFormat, paths.Home(), pipelineID, spec.Cmd)
}

// AgentMonitoringEndpoint returns endpoint with exposed metrics for agent.
func AgentMonitoringEndpoint(operatingSystem string, cfg *monitoringConfig.MonitoringHTTPConfig) string {
	if cfg != nil && cfg.Enabled {
		return fmt.Sprintf(agentMbEndpointHTTP, cfg.Host, cfg.Port)
	}

	if operatingSystem == "windows" {
		return agentMbEndpointFileFormatWin
	}
	// unix socket path must be less than 104 characters
	path := fmt.Sprintf("unix://%s.sock", filepath.Join(paths.TempDir(), "elastic-agent"))
	if len(path) < 104 {
		return path
	}
	// place in global /tmp to ensure that its small enough to fit; current path is way to long
	// for it to be used, but needs to be unique per Agent (in the case that multiple are running)
	return fmt.Sprintf(`unix:///tmp/elastic-agent/%x.sock`, sha256.Sum256([]byte(path)))
}

// AgentPrefixedMonitoringEndpoint returns endpoint with exposed metrics for agent.
func AgentPrefixedMonitoringEndpoint(operatingSystem string, cfg *monitoringConfig.MonitoringHTTPConfig) string {
	return httpPlusPrefix + AgentMonitoringEndpoint(operatingSystem, cfg)
}
