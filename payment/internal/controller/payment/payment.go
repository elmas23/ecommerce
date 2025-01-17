package payment

import (
	"context"

	"github.com/elmas23/ecommerce/payment/internal/entity"
	"github.com/elmas23/ecommerce/payment/internal/repository/payment"
)

type Controller interface {
	Charge(ctx context.Context, payment entity.Payment) (entity.Payment, error)
}

type controller struct {
	paymentRepo payment.Repository
}

func NewController(ctx context.Context) *controller {
	return &controller{
		paymentRepo: payment.NewRepository(ctx),
	}
}

func (c *controller) Charge(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
	err := c.paymentRepo.SavePayment(ctx, &payment)
	if err != nil {
		return entity.Payment{}, err
	}
	return payment, nil
}
