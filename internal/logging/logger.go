package logging

import (
	"gnss_corr/internal/config/modules"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvLoggerConfig struct {
	LogLevel   int8 `envconfig:"LOG_LEVEL"    default:"0"`
	LogDevMode bool `envconfig:"LOG_DEV_MODE" default:"false"`
}

func LoadLoggerConfig() (EnvLoggerConfig, error) {
	var config EnvLoggerConfig
	if err := envconfig.Process(modules.GlobalEnvPrefix, &config); err != nil {
		return EnvLoggerConfig{}, errors.Wrap(err, "cannot get gnss corrections logger config")
	}

	if config.LogLevel < -1 || config.LogLevel > 2 {
		return EnvLoggerConfig{}, errors.Wrap(errors.New("invalid log level"), "cannot get gnss corrections logger level")
	}

	return config, nil
}
