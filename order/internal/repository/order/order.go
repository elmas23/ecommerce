package order

import (
	"context"
	"fmt"

	"github.com/elmas23/ecommerce/order/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	// GetOrder returns the order with the specified ID
	GetOrder(ctx context.Context, id string) (entity.Order, error)
	// SaveOrder saves a new order
	SaveOrder(ctx context.Context, order *entity.Order) error
}

type repository struct {
	// inject gorm db dependency into the repository
	db *gorm.DB
}

func NewRepository(ctx context.Context) *repository {
	db, openErr := gorm.Open(mysql.Open("root:verysecretpass@tcp(127.0.0.1:3306)/order"), &gorm.Config{})
	if openErr != nil {
		panic(fmt.Errorf("error connecting to database: %v", openErr))
	}
	err := db.AutoMigrate(&entity.OrderModel{}, &entity.OrderItemModel{})
	if err != nil {
		panic(fmt.Errorf("error migrating database: %v", err))
	}
	return &repository{db: db}
}

// GetOrder returns the order with the specified ID
func (r *repository) GetOrder(ctx context.Context, id string) (entity.Order, error) {
	var orderEntity entity.Order
	res := r.db.First(&orderEntity, id)
	var orderItems []entity.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, entity.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := entity.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt,
	}
	return order, res.Error
}

// SaveOrder saves a new order
func (r *repository) SaveOrder(ctx context.Context, order *entity.Order) error {
	var orderItems []entity.OrderItemModel
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, entity.OrderItemModel{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	orderModel := entity.OrderModel{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := r.db.Create(&orderModel)
	if res.Error != nil {
		return res.Error
	}
	order.ID = int64(orderModel.ID)
	return nil
}
