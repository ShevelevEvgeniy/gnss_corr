package gnss_corr

import (
	"gnss_corr/internal/config/modules"

	"github.com/pkg/errors"
)

type (
	Config struct {
		service modules.Service
		storage storage
	}

	storage struct {
		postgres modules.Postgres
	}
)

func Load() (Config, error) {
	service, err := modules.Load[modules.Service](modules.GlobalEnvPrefix)
	if err != nil {
		return Config{}, errors.Wrapf(err, "failed to load search data extractor configuration")
	}

	storageConfig, err := loadStorage(modules.GlobalEnvPrefix)
	if err != nil {
		return Config{}, errors.Wrap(err, "error loading storage configuration")
	}

	return Config{
		service: service,
		storage: storageConfig,
	}, nil
}

func loadStorage(prefix string) (storage, error) {
	postgres, err := modules.Load[modules.Postgres](prefix)
	if err != nil {
		return storage{}, errors.Wrap(err, "error loading postgres configuration")
	}

	return storage{
		postgres: postgres,
	}, nil
}
