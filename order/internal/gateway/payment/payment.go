package payment

import (
	"context"
	"fmt"

	paymentpb "github.com/elmas23/ecommerce-idl/golang/payment"
	"github.com/elmas23/ecommerce/order/internal/entity"
	"github.com/elmas23/ecommerce/order/internal/gateway/payment/mapper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const paymentServiceAddress = "localhost:3001"

type Gateway interface {
	Charge(ctx context.Context, paymentOrder entity.PaymentOrder) error
}

type gateway struct {
	// we need to inject the payment client dependency into the gateway
	client paymentpb.PaymentClient
}

func NewGateway(ctx context.Context) *gateway {
	// create a new gRPC client connection to the payment service
	conn, err := grpc.NewClient(
		paymentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to payment service: %v", err))
	}
	defer conn.Close()
	// create a new payment client
	client := paymentpb.NewPaymentClient(conn)
	return &gateway{client: client}
}

func (g *gateway) Charge(ctx context.Context, paymentOrder entity.PaymentOrder) error {
	// build the request for Charge endpoint
	req := mapper.ToCreatePaymentRequest(paymentOrder)
	_, err := g.client.Create(ctx, req)
	return err
}
