package modules

import (
	"time"
)

type Service struct {
	Transport         string        `envconfig:"SERVER_TRANSPORT"    default:"tcp" logKey:"transport"`
	Port              int           `envconfig:"SERVER_PORT"         default:"80"  logKey:"port"`
	MetricsHealthPort int           `envconfig:"METRICS_HEALTH_PORT" default:"3000"`
	ReadHeaderTimeout time.Duration `envconfig:"READ_HEADER_TIMEOUT" default:"2s"`
	WriteTimeout      time.Duration `envconfig:"WRITE_TIMEOUT"       default:"5m"`
}
