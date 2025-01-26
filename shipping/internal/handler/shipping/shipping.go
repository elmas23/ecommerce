package shipping

import (
	"context"

	shippingpb "github.com/elmas23/ecommerce-idl/golang/shipping"
	"github.com/elmas23/ecommerce/shipping/internal/controller/shipping"
	"github.com/elmas23/ecommerce/shipping/internal/entity"
)

type handler struct {
	shippingpb.UnimplementedShippingServer
	shippingController shipping.Controller
}

func NewHandler(ctx context.Context) *handler {
	return &handler{
		shippingController: shipping.NewController(ctx),
	}
}

func (h *handler) Create(ctx context.Context, req *shippingpb.CreateShippingRequest) (*shippingpb.CreateShippingResponse, error) {
	newShipping := entity.NewShipping(req.Address)
	_, err := h.shippingController.CreateShipping(ctx, newShipping)
	if err != nil {
		return nil, err
	}
	return &shippingpb.CreateShippingResponse{}, nil
}
