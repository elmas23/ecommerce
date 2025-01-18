package main

import (
	"context"
	"log"
	"net"

	orderpb "github.com/elmas23/ecommerce-idl/golang/order"
	orderHandler "github.com/elmas23/ecommerce/order/internal/handler/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3000"
)

func main() {
	ctx := context.Background()

	handler := orderHandler.NewHandler(ctx)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on port %s", port)

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", port, err)
	}
}
