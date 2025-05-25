package rover

import (
	"context"

	"github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r RoverServer) RegisterRover(ctx context.Context, request *dto.RegisterRoverRequest) (*dto.RegisterRoverResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = r.postgresStorage
	return nil, nil
}
