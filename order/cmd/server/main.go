package main

import (
	"context"
	"log"
	"net"

	"github.com/elmas23/ecommerce/order/internal/controller"
	"github.com/elmas23/ecommerce/order/internal/handler"
	"github.com/elmas23/ecommerce/order/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	orderpb "github.com/elmas23/ecommerce-idl/golang/order"
)

const (
	port = ":50051"
)

func main() {
	ctx := context.Background()

	repository := repository.NewRepository(ctx)
	controller := controller.NewController(ctx, repository)
	handler := handler.NewHandler(ctx, controller)

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