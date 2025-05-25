package main

import (
	descBaseStation "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/base_station_v1"
	descCorrCordRover "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/coord_correction_v1"
	descRover "github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/api/rover_v1"
	baseStation "gnss_corr/internal/grpc/base_station"
	"gnss_corr/internal/grpc/coord_correction"
	"gnss_corr/internal/grpc/rover"
	"gnss_corr/internal/storage/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func registerServer(logger *zap.SugaredLogger, postgresStorage postgres.PostgresStorage) *grpc.Server {
	server := grpc.NewServer()

	reflection.Register(server)
	descRover.RegisterGNSSCorrectionServiceServer(server, rover.NewRoverServer(logger, postgresStorage))
	descBaseStation.RegisterGNSSCorrectionServiceServer(server, baseStation.NewBaseStationServer(logger, postgresStorage))
	descCorrCordRover.RegisterGNSSCorrectionServiceServer(server, coord_correction.NewCoordinatesCorrectionServer(logger, postgresStorage))

	return server
}
