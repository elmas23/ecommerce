package order

import (
	"context"

	"github.com/elmas23/ecommerce/order/internal/entity"
	"github.com/elmas23/ecommerce/order/internal/gateway/payment"
	"github.com/elmas23/ecommerce/order/internal/repository/order"
	errorlib "github.com/elmas23/ecommerce/order/internal/utils/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		// if the payment fails, return an error with the details
		badReq := errorlib.HandleBadRequest(paymentErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return entity.Order{}, statusWithDetails.Err()
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
