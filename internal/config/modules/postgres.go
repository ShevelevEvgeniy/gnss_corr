package modules

import (
	"fmt"
)

type Postgres struct {
	Host               string `envconfig:"POSTGRES_HOST"                 required:"true"   logKey:"postgres_host"`
	Port               uint16 `envconfig:"POSTGRES_PORT"                 required:"true"   logKey:"postgres_port"`
	User               string `envconfig:"POSTGRES_USER"                 required:"true"   logKey:"postgres_user,secret"`
	Password           string `envconfig:"POSTGRES_PASSWORD"             required:"true"   logKey:"postgres_password,secret"`
	DBName             string `envconfig:"POSTGRES_DB"                   required:"true"   logKey:"postgres_db"`
	SSLMode            string `envconfig:"POSTGRES_SSL_MODE"             default:"disable" logKey:"postgres_ssl_mode"`
	MaxTaskAttempts    uint64 `envconfig:"POSTGRES_MAX_TASK_ATTEMPTS"    default:"5"       logKey:"postgres_max_task_attempts"`
	MaxRetriesAttempts uint64 `envconfig:"POSTGRES_MAX_RETRIES_ATTEMPTS" default:"5"       logKey:"postgres_max_retries_attempts"`
}

func (c Postgres) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}
