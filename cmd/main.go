package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	config "gnss_corr/internal/config/gnss_corr"
	"gnss_corr/internal/logging"
	"gnss_corr/internal/storage/postgres"
	loggingConfig "gnss_corr/pkg/log_config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	postgresStorage, err := postgres.NewPostgresStorage(ctx, logger, config.Storage.Postgres)
	if err != nil {
		logger.Errorw("cannot create task storage", "error", err.Error(), "time", time.Now().UTC().String())

		return
	}
	defer postgresStorage.Close()

	listen, err := net.Listen(config.Service.Transport, fmt.Sprintf(":%d", config.Service.Port))
	if err != nil {
		logger.Error(errors.Wrap(err, "listen port error"))

		return
	}

	server := registerServer(logger, postgresStorage)

	go gracefulShutdown(server, logger)

	err = server.Serve(listen)
	if err != nil {
		logger.Error(errors.Wrap(err, "serve error"))

		return
	}
}

func gracefulShutdown(server *grpc.Server, logger *zap.SugaredLogger) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	logger.Info("Received shutdown signal, shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.GracefulStop()

	<-shutdownCtx.Done()

	logger.Info("Shutdown complete")
}
