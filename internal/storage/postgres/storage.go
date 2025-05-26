package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"gnss_corr/internal/config/modules"
	"go.uber.org/zap"
)

type PostgresStorage struct {
	logger             *zap.SugaredLogger
	db                 *pgxpool.Pool
	maxTaskAttempts    uint64
	maxRetriesAttempts uint64
}

func NewPostgresStorage(ctx context.Context, logger *zap.SugaredLogger, cfg modules.Postgres) (PostgresStorage, error) {
	postgresClient, err := NewPostgresClient(ctx, cfg)
	if err != nil {
		logger.Errorw("cannot create postgres clients", "error", err.Error(), "time", time.Now().UTC().String())

		return PostgresStorage{}, errors.Wrap(err, "cannot create postgres clients")
	}

	return PostgresStorage{
		logger:             logger,
		db:                 postgresClient,
		maxRetriesAttempts: cfg.MaxRetriesAttempts,
	}, nil
}

func NewPostgresClient(ctx context.Context, config modules.Postgres) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(ctx, config.GetConnectionString())
	if err != nil {
		return nil, errors.Wrap(err, "error pgxpool.New")
	}

	if err = dbPool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "error dbPool.Ping")
	}

	return dbPool, nil
}

func (t PostgresStorage) Close() {
	t.db.Close()
}
