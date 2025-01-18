package mapper

import (
	paymentpb "github.com/elmas23/ecommerce-idl/golang/payment"
	"github.com/elmas23/ecommerce/order/internal/entity"
)

func ToCreatePaymentRequest(paymentOrder entity.PaymentOrder) *paymentpb.CreatePaymentRequest {
	return &paymentpb.CreatePaymentRequest{
		OrderId:    paymentOrder.OrderId,
		UserId:     paymentOrder.CustomerId,
		TotalPrice: paymentOrder.TotalPrice,
	}
}
