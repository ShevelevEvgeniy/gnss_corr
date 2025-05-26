package modules

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
)

const GlobalEnvPrefix string = "GC"

func Load[T any](prefix string) (T, error) {
	var config T
	if err := envconfig.Process(prefix, &config); err != nil {
		return config, err
	}

	return config, nil
}

func GetGlobalPrefix(prefixes ...string) string {
	return strings.Join(prefixes, "_")
}
