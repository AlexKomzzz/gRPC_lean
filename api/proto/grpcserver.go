package adder

import (
	"context"

	api "github.com/AlexKomzzz/gRPC_lean/pkg/api"
)

// GRPCServer
type GRPCServer struct {
}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponce, error) {
	return &api.AddResponce{Result: req.GetX() + req.GetY()}, nil
}
