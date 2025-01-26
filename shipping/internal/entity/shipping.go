package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Shipping entity
type Shipping struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
}

type ShippingModel struct {
	gorm.Model
	Address string
}

func NewShipping(address string) Shipping {
	return Shipping{
		ID:      int64(uuid.New().ID()),
		Address: address,
	}
}
