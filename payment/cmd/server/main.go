package main

import (
	"context"
	"log"
	"net"

	paymentpb "github.com/elmas23/ecommerce-idl/golang/payment"
	paymentHanlder "github.com/elmas23/ecommerce/payment/internal/handler/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

func main() {
	ctx := context.Background()

	handler := paymentHanlder.NewHandler(ctx)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	paymentpb.RegisterPaymentServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server on port %s", port)

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", port, err)
	}
}
