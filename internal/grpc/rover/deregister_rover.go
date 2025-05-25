package rover

import (
	"context"

	"github.com/ShevelevEvgeniy/geodesy_proto/gen/proto/dto"
)

func (r RoverServer) DeregisterRover(ctx context.Context, request *dto.RoverID) (*dto.DeregisterRoverResponse, error) {
	return nil, nil
}
