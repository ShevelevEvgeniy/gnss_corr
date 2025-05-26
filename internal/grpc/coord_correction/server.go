package coord_correction

import (
	desc "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/coord_correction_v1"
	"go.uber.org/zap"
)

type (
	CoordinatesCorrectionServer struct {
		desc.UnimplementedCoordinatesCorrectionServiceServer
		logger          *zap.SugaredLogger
		postgresStorage PostgresStorage
	}

	PostgresStorage interface {
	}
)

func NewCoordinatesCorrectionServer(logger *zap.SugaredLogger, postgresStorage PostgresStorage) *CoordinatesCorrectionServer {
	return &CoordinatesCorrectionServer{
		logger:          logger,
		postgresStorage: postgresStorage,
	}
}
