package service

import (
	"context"
	"payment_service/internal/entity"
	"payment_service/internal/paymentpb"
)

type PaymentServiceServer interface {
	GeneratePaymentLink(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error)
	// CreatePayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
}

type PaymentService interface {
	PaymentProcess(ctx context.Context, paymentID int64, req entity.PaymentRequest) error
}
