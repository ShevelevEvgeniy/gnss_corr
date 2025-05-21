package modules

import (
	"fmt"

	"go.uber.org/zap"
)

type Postgres struct {
	Host               string `envconfig:"POSTGRES_HOST"                 required:"true"`
	Port               uint16 `envconfig:"POSTGRES_PORT"                 required:"true"`
	User               string `envconfig:"POSTGRES_USER"                 required:"true"`
	Password           string `envconfig:"POSTGRES_PASSWORD"             required:"true"`
	DBName             string `envconfig:"POSTGRES_DB"                   required:"true"`
	SSLMode            string `envconfig:"POSTGRES_SSL_MODE"             default:"disable"`
	MaxTaskAttempts    uint64 `envconfig:"POSTGRES_MAX_TASK_ATTEMPTS"    default:"5"`
	MaxRetriesAttempts uint64 `envconfig:"POSTGRES_MAX_RETRIES_ATTEMPTS" default:"5"`
}

func (c Postgres) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

func (c Postgres) HasUser() bool {
	return c.User != ""
}

func (c Postgres) HasPassword() bool {
	return c.Password != ""
}

func (c Postgres) Log(logger *zap.SugaredLogger) {
	logger.Infow("postgres host", "value", c.Host)
	logger.Infow("postgres port", "value", c.Port)
	logger.Infow("postgres db name", "value", c.DBName)
	logger.Infow("postgres ssl mode", "value", c.SSLMode)
	logger.Infow("postgres user exists", "value", c.HasUser())
	logger.Infow("postgres password exists", "value", c.HasPassword())
	logger.Infow("postgres max task attempts", "value", c.MaxTaskAttempts)
	logger.Infow("postgres max retries attempts", "value", c.MaxTaskAttempts)
}
