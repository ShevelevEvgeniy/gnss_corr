package modules

import (
	"time"

	"go.uber.org/zap"
)

type Service struct {
	BackendPort       int           `envconfig:"BACKEND_PORT"        default:"80"`
	MetricsHealthPort int           `envconfig:"METRICS_HEALTH_PORT" default:"3000"`
	ReadHeaderTimeout time.Duration `envconfig:"READ_HEADER_TIMEOUT" default:"2s"`
	WriteTimeout      time.Duration `envconfig:"WRITE_TIMEOUT"       default:"5m"`
}

func (c Service) Log(logger *zap.SugaredLogger) {
	logger.Infow("search data extractor backend port", "value", c.BackendPort)
	logger.Infow("search data extractor metrics health port", "value", c.MetricsHealthPort)
	logger.Infow("search data extractor read header timeout", "value", c.ReadHeaderTimeout)
	logger.Infow("search data extractor write timeout", "value", c.WriteTimeout)
}
