package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func WithDevelopmentMode() func(config *LoggerConfig) {
	return func(config *LoggerConfig) {
		config.config = zap.NewDevelopmentConfig()
	}
}

func WithLevel(level int8) func(config *LoggerConfig) {
	return func(config *LoggerConfig) {
		config.level = zapcore.Level(level)
	}
}
