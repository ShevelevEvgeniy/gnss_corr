package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	config "gnss_corr/internal/config/gnss_corr"
	loggingConfig "gnss_corr/internal/config/log_config"
	"gnss_corr/internal/logging"
	"gnss_corr/internal/storage/postgres"
)

func main() {
	ctx := context.Background()

	logger, err := logging.InitLogger()
	if err != nil {
		log.Fatal(errors.Wrap(err, "load search-data-extractor logger config error"))
	}

	config, err := config.Load()
	if err != nil {
		logger.Errorw("cannot load configs", "error", err.Error(), "time", time.Now().UTC().String())

		return
	}

	loggingConfig.LogConfig(config)

	taskStorage, err := postgres.NewTaskStorage(ctx, logger, config.Postgres())
	if err != nil {
		logger.Errorw("cannot create task storage", "error", err.Error(), "time", time.Now().UTC().String())

		return
	}
	defer taskStorage.Close()
}
