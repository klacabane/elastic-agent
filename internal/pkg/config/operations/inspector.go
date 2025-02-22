// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package operations

import (
	"fmt"

	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent/internal/pkg/agent/application/info"
	"github.com/elastic/elastic-agent/internal/pkg/agent/application/paths"
	"github.com/elastic/elastic-agent/internal/pkg/agent/configuration"
	"github.com/elastic/elastic-agent/internal/pkg/agent/errors"
	"github.com/elastic/elastic-agent/internal/pkg/agent/storage"
	"github.com/elastic/elastic-agent/internal/pkg/agent/storage/store"
	"github.com/elastic/elastic-agent/internal/pkg/config"
	"github.com/elastic/elastic-agent/internal/pkg/fleetapi"
	"github.com/elastic/elastic-agent/pkg/core/logger"
)

var (
	// ErrNoFleetConfig is returned when no configuration was retrieved from fleet just yet.
	ErrNoFleetConfig = fmt.Errorf("no fleet config retrieved yet")
)

// LoadFullAgentConfig load agent config based on provided paths and defined capabilities.
// In case fleet is used, config from policy action is returned.
func LoadFullAgentConfig(logger *logger.Logger, cfgPath string, failOnFleetMissing bool) (*config.Config, error) {
	rawConfig, err := loadConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	cfg, err := configuration.NewFromConfig(rawConfig)
	if err != nil {
		return nil, err
	}

	if configuration.IsStandalone(cfg.Fleet) {
		// When in standalone we load the configuration again with inputs that are defined in the paths.ExternalInputs.
		loader := config.NewLoader(logger, paths.ExternalInputs())
		discover := config.Discoverer(cfgPath, cfg.Settings.Path, paths.ExternalInputs())
		files, err := discover()
		if err != nil {
			return nil, fmt.Errorf("could not discover configuration files: %w", err)
		}
		if len(files) == 0 {
			return nil, config.ErrNoConfiguration
		}
		c, err := loader.Load(files)
		if err != nil {
			return nil, fmt.Errorf("failed to load or merge configuration: %w", err)
		}
		return c, nil
	}

	fleetConfig, err := loadFleetConfig()
	if err != nil {
		return nil, err
	} else if fleetConfig == nil {
		if failOnFleetMissing {
			return nil, ErrNoFleetConfig
		}

		// resolving fleet config but not fleet config retrieved yet, returning last applied config
		return rawConfig, nil
	}

	return config.NewConfigFrom(fleetConfig)
}

func loadConfig(configPath string) (*config.Config, error) {
	rawConfig, err := config.LoadFile(configPath)
	if err != nil {
		return nil, err
	}

	path := paths.AgentConfigFile()

	store := storage.NewEncryptedDiskStore(path)
	reader, err := store.Load()
	if err != nil {
		return nil, errors.New(err, "could not initialize config store",
			errors.TypeFilesystem,
			errors.M(errors.MetaKeyPath, path))
	}

	config, err := config.NewConfigFrom(reader)
	if err != nil {
		return nil, errors.New(err,
			fmt.Sprintf("fail to read configuration %s for the elastic-agent", path),
			errors.TypeFilesystem,
			errors.M(errors.MetaKeyPath, path))
	}

	// merge local configuration and configuration persisted from fleet.
	_ = rawConfig.Merge(config)

	if err := info.InjectAgentConfig(rawConfig); err != nil {
		return nil, err
	}

	return rawConfig, nil
}

func loadFleetConfig() (map[string]interface{}, error) {
	log, err := newErrorLogger()
	if err != nil {
		return nil, err
	}

	stateStore, err := store.NewStateStoreWithMigration(log, paths.AgentActionStoreFile(), paths.AgentStateStoreFile())
	if err != nil {
		return nil, err
	}

	for _, c := range stateStore.Actions() {
		cfgChange, ok := c.(*fleetapi.ActionPolicyChange)
		if !ok {
			continue
		}

		return cfgChange.Policy, nil
	}
	return nil, nil
}

func newErrorLogger() (*logger.Logger, error) {
	return logger.NewWithLogpLevel("", logp.ErrorLevel, false)
}
