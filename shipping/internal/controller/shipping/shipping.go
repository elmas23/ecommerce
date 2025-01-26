package shipping

import (
	"context"

	"github.com/elmas23/ecommerce/shipping/internal/entity"
	"github.com/elmas23/ecommerce/shipping/internal/repository/shipping"
)

type Controller interface {
	CreateShipping(ctx context.Context, shipping entity.Shipping) (entity.Shipping, error)
}

type controller struct {
	shippingRepo shipping.Repository
}

func NewController(ctx context.Context) *controller {
	return &controller{
		shippingRepo: shipping.NewRepository(ctx),
	}
}

func (c *controller) CreateShipping(ctx context.Context, shipping entity.Shipping) (entity.Shipping, error) {
	err := c.shippingRepo.Create(ctx, &shipping)
	if err != nil {
		return entity.Shipping{}, err
	}
	return shipping, nil
}
