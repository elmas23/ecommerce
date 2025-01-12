package order

import (
	"context"

	orderpb "github.com/elmas23/ecommerce-idl/golang/order"
	"github.com/elmas23/ecommerce/order/internal/controller/order"
	"github.com/elmas23/ecommerce/order/internal/entity"
)

type handler struct {
	orderpb.UnimplementedOrderServer
	orderController order.Controller
}

func NewHandler(ctx context.Context) *handler {
	return &handler{
		orderController: order.NewController(ctx),
	}
}

func (h *handler) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	var orderItems []entity.OrderItem
	for _, orderItem := range req.OrderItems {
		orderItems = append(orderItems, entity.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := entity.NewOrder(req.UserId, orderItems)
	order, err := h.orderController.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	return &orderpb.CreateOrderResponse{
		OrderId: order.ID,
	}, nil
}
