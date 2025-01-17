package payment

import (
	"context"
	"fmt"

	paymentpb "github.com/elmas23/ecommerce-idl/golang/payment"
	"github.com/elmas23/ecommerce/payment/internal/controller/payment"
	"github.com/elmas23/ecommerce/payment/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	paymentpb.UnimplementedPaymentServer
	paymentController payment.Controller
}

func NewHandler(ctx context.Context) *handler {
	return &handler{
		paymentController: payment.NewController(ctx),
	}
}

func (h *handler) Create(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	newPayment := entity.NewPayment(req.UserId, req.OrderId, req.TotalPrice)
	result, err := h.paymentController.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &paymentpb.CreatePaymentResponse{PaymentId: result.ID}, nil
}
