package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenon := "127.0.0.1:8080"

	listener, err := net.Listen("tcp", listenon)

	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenon, err)
	}

	server := grpc.NewServer()

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
