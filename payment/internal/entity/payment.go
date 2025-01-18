package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Payment status
const (
	PaymentStatusPending  = "Pending"
	PaymentStatusPaid     = "Paid"
	PaymentStatusCanceled = "Canceled"
)

type Payment struct {
	ID         int64   `json:"id"`
	CustomerId int64   `json:"customer_id"`
	Status     string  `json:"status"`
	OrderId    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  int64   `json:"created_at"`
}

type PaymentModel struct {
	gorm.Model
	CustomerId int64
	Status     string
	OrderId    int64
	TotalPrice float32
}

func NewPayment(customerId int64, orderId int64, totalPrice float32) Payment {
	return Payment{
		ID:         int64(uuid.New().ID()),
		CreatedAt:  time.Now().Unix(),
		Status:     PaymentStatusPending,
		CustomerId: customerId,
		OrderId:    orderId,
		TotalPrice: totalPrice,
	}
}
