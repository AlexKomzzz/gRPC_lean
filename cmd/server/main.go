package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/AlexKomzzz/gRPC_lean/pkg/adder"
	"github.com/AlexKomzzz/gRPC_lean/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {

	srv := &adder.GRPCServer{}
	go func() {
		// mux
		mux := runtime.NewServeMux()
		// register
		api.RegisterAdderHandlerServer(context.Background(), mux, srv)

		// http.server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	api.RegisterAdderServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}
}
