package repository

import (
	"context"

	"payment_service/internal/entity"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	GetPayment(ctx context.Context, id int64) (*entity.Payment, error)
	UpdatePayment(ctx context.Context, payment *entity.Payment) error
	DeletePayment(ctx context.Context, id string) error
	ValidateCard(ctx context.Context, cardNum, cvv, expDate, name string) (bool, error)
	CheckAndDeductBalance(ctx context.Context, cardNum string, amount float64) (bool, error)
	Close() error
}
