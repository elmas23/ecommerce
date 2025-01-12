package order

import (
	"context"

	"github.com/elmas23/ecommerce/order/internal/entity"
	"github.com/elmas23/ecommerce/order/internal/repository/order"
)

type Controller interface {
	PlaceOrder(ctx context.Context, order entity.Order) (entity.Order, error)
}

type controller struct {
	orderRepo order.Repository
}

func NewController(ctx context.Context) *controller {
	return &controller{
		orderRepo: order.NewRepository(ctx),
	}
}

func (c *controller) PlaceOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	err := c.orderRepo.SaveOrder(ctx, &order)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
