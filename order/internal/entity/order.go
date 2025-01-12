package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order status
const (
	OrderStatusPending   = "Pending"
	OrderStatusPaid      = "Paid"
	OrderStatusShipped   = "Shipped"
	OrderStatusDelivered = "Delivered"
	OrderStatusCanceled  = "Canceled"
)

// OrderItem entity
type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

// Order entity
type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

// OrderModel is the gorm model for Order
type OrderModel struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItemModel
}

// OrderItemModel is the gorm model for OrderItem
type OrderItemModel struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

// NewOrder creates a new defautl order in pending status
func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		ID:         int64(uuid.New().ID()),
		CreatedAt:  time.Now().Unix(),
		Status:     OrderStatusPending,
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}
