package adder

import (
	"context"

	"github.com/andreipimenov/zerohero/lesson158/pkg/api"
)

type Server struct {
	api.UnimplementedAdderServer
}

func (s *Server) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponce, error) {
	x := req.X
	y := req.Y
	sum := x + y

	res := &api.AddResponce{
		Result: sum,
	}
	return res, nil
}
