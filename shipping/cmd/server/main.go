package main

import (
	"context"
	"log"
	"net"

	shippingpb "github.com/elmas23/ecommerce-idl/golang/shipping"
	shippingHandler "github.com/elmas23/ecommerce/shipping/internal/handler/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3002"
)

func main() {
	ctx := context.Background()

	handler := shippingHandler.NewHandler(ctx)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	shippingpb.RegisterShippingServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on port %s", port)

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", port, err)
	}
}
