package logging

import (
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	level  zapcore.Level
	config zap.Config
}

func InitLogger() (*zap.SugaredLogger, error) {
	loggerConfig, err := LoadLoggerConfig()
	if err != nil {
		return nil, errors.Wrap(err, "load search-service logger config error")
	}

	return NewLogger(loggerConfig), nil
}

func NewLogger(config EnvLoggerConfig) *zap.SugaredLogger {
	if config.LogDevMode {
		return BuildLogger(WithDevelopmentMode(), WithLevel(config.LogLevel))
	}

	return BuildLogger(WithLevel(config.LogLevel))
}

func BuildLogger(options ...func(logger *LoggerConfig)) *zap.SugaredLogger {
	loggerConfig := &LoggerConfig{}
	productionConfig := zap.NewProductionConfig()
	productionConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.RFC3339))
	}
	productionConfig.Level.SetLevel(zapcore.InfoLevel)

	loggerConfig.config = productionConfig

	for _, o := range options {
		o(loggerConfig)
	}

	logger, err := loggerConfig.config.Build()
	if err != nil {
		panic("Failed to setup logger")
	}

	return logger.Sugar()
}
