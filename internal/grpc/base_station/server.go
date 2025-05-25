package base_station

import (
	desc "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/base_station_v1"
	"go.uber.org/zap"
)

type (
	BaseStationServer struct {
		desc.UnimplementedGNSSCorrectionServiceServer
		logger          *zap.SugaredLogger
		postgresStorage PostgresStorage
	}

	PostgresStorage interface {
	}
)

func NewBaseStationServer(logger *zap.SugaredLogger, postgresStorage PostgresStorage) *BaseStationServer {
	return &BaseStationServer{
		logger:          logger,
		postgresStorage: postgresStorage,
	}
}
