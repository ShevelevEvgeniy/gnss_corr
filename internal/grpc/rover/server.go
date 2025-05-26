package rover

import (
	desc "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/rover_v1"
	"go.uber.org/zap"
)

type (
	RoverServer struct {
		desc.UnimplementedRoverServiceServer
		logger          *zap.SugaredLogger
		postgresStorage PostgresStorage
	}

	PostgresStorage interface {
		//RegisterRover(ctx context.Context, rover models.Rover) error
	}
)

func NewRoverServer(logger *zap.SugaredLogger, postgresStorage PostgresStorage) *RoverServer {
	return &RoverServer{
		logger:          logger,
		postgresStorage: postgresStorage,
	}
}
