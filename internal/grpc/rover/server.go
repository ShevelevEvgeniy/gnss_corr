package rover

import (
	"context"

	desc "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/rover_v1"
	"gnss_corr/internal/models"
	"go.uber.org/zap"
)

type (
	RoverServer struct {
		desc.UnimplementedGNSSCorrectionServiceServer
		logger          *zap.SugaredLogger
		postgresStorage PostgresStorage
	}

	PostgresStorage interface {
		RegisterRover(ctx context.Context, rover models.Rover) error
	}
)

func NewRoverServer(logger *zap.SugaredLogger, postgresStorage PostgresStorage) *RoverServer {
	return &RoverServer{
		logger:          logger,
		postgresStorage: postgresStorage,
	}
}
