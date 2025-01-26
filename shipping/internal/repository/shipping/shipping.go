package shipping

import (
	"context"
	"fmt"

	"github.com/elmas23/ecommerce/shipping/internal/entity"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, shipping *entity.Shipping) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(ctx context.Context) *repository {
	db, openErr := gorm.Open(mysql.Open("root:verysecretpass@tcp(127.0.0.1:3308)/shipping"), &gorm.Config{})
	if openErr != nil {
		panic(fmt.Errorf("error connecting to database: %v", openErr))
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("shipping"))); err != nil {
		panic(fmt.Errorf("db otel plugin error: %v", err))
	}

	err := db.AutoMigrate(&entity.ShippingModel{})
	if err != nil {
		panic(fmt.Errorf("db migration error: %v", err))
	}
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, shipping *entity.Shipping) error {
	shippingModel := entity.ShippingModel{
		Address: shipping.Address,
	}
	res := r.db.WithContext(ctx).Create(&shippingModel)
	return res.Error
}
