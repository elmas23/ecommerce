package order

import (
	"context"

	"github.com/elmas23/ecommerce/order/internal/entity"
	"github.com/elmas23/ecommerce/order/internal/gateway/payment"
	"github.com/elmas23/ecommerce/order/internal/repository/order"
)

type Controller interface {
	PlaceOrder(ctx context.Context, order entity.Order) (entity.Order, error)
}

type controller struct {
	orderRepo      order.Repository
	paymentGateway payment.Gateway
}

func NewController(ctx context.Context) *controller {
	return &controller{
		orderRepo:      order.NewRepository(ctx),
		paymentGateway: payment.NewGateway(ctx),
	}
}

func (c *controller) PlaceOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	// save the order to the database
	err := c.orderRepo.SaveOrder(ctx, &order)
	if err != nil {
		return entity.Order{}, err
	}
	// Once saved, create charge the customer
	paymentOrder := entity.PaymentOrder{
		OrderId:    order.ID,
		CustomerId: order.CustomerID,
		TotalPrice: c.calculateTotalPrice(order),
	}
	paymentErr := c.paymentGateway.Charge(ctx, paymentOrder)
	if paymentErr != nil {
		return entity.Order{}, paymentErr
	}
	return order, nil
}

// calculateTotalPrice calculates the total price of an order for all the order items
func (c *controller) calculateTotalPrice(order entity.Order) float32 {
	var totalPrice float32
	for _, orderItem := range order.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}
