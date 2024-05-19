package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/andreipimenov/zerohero/lesson127/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterTimeServiceServer(grpcServer, &server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}

type server struct {
	api.TimeServiceServer
}

func (s *server) GetNow(ctx context.Context, req *api.GetNowRequest) (*api.GetNowResponse, error) {
	tz := req.Tz
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	now := time.Now().In(loc).String()
	resp := &api.GetNowResponse{
		Now: now,
		Tz:  tz,
	}
	return resp, nil

}

func (s *server) GetTomorrow(context.Context, *api.GetTomorrowRequest) (*api.GetTomorrowResponse, error) {
	t := time.Now().Add(24 * time.Hour).UTC().String()
	resp := &api.GetTomorrowResponse{
		Tomorrow: t,
	}
	return resp, nil

}
