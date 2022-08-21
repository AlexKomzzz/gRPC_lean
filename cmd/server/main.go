package main

import (
	"log"
	"net"

	"github.com/AlexKomzzz/gRPC_lean/pkg/adder"
	"github.com/AlexKomzzz/gRPC_lean/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}
}
