package payment

import (
	"context"
	"fmt"

	"github.com/elmas23/ecommerce/payment/internal/entity"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	GetPayment(ctx context.Context, id int64) (entity.Payment, error)
	SavePayment(ctx context.Context, payment *entity.Payment) error
}

type repository struct {
	// inject gorm db dependency into the repository
	db *gorm.DB
}

func NewRepository(ctx context.Context) *repository {
	db, openErr := gorm.Open(mysql.Open("root:verysecretpass@tcp(127.0.0.1:3307)/payment"), &gorm.Config{})
	if openErr != nil {
		panic(fmt.Errorf("error connecting to database: %v", openErr))
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("payment"))); err != nil {
		panic(fmt.Errorf("db otel plugin error: %v", err))
	}

	err := db.AutoMigrate(&entity.PaymentModel{})
	if err != nil {
		panic(fmt.Errorf("db migration error: %v", err))
	}
	return &repository{db: db}
}

func (r *repository) GetPayment(ctx context.Context, id int64) (entity.Payment, error) {
	var paymentEntity entity.PaymentModel
	res := r.db.First(&paymentEntity, id)
	payment := entity.Payment{
		ID:         int64(paymentEntity.ID),
		Status:     paymentEntity.Status,
		OrderId:    paymentEntity.OrderId,
		CustomerId: paymentEntity.CustomerId,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}
	return payment, res.Error
}

func (r *repository) SavePayment(ctx context.Context, payment *entity.Payment) error {
	paymentModel := entity.PaymentModel{
		CustomerId: payment.CustomerId,
		Status:     payment.Status,
		OrderId:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
	}
	res := r.db.WithContext(ctx).Create(&paymentModel)
	if res.Error == nil {
		payment.ID = int64(paymentModel.ID)
	}
	return res.Error
}
