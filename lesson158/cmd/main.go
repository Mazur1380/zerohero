package main

import (
	"log"
	"net"

	"github.com/andreipimenov/zerohero/lesson158/internal/adder"
	"github.com/andreipimenov/zerohero/lesson158/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()
	srv := &adder.Server{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
