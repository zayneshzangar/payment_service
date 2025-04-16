package grpc_service

import (
	"context"
	"fmt"

	"payment_service/internal/entity"
	"payment_service/internal/paymentpb"
	"payment_service/internal/repository"
	"payment_service/internal/service"
)

type paymentService struct {
	repo repository.PaymentRepository
}

func NewGrpcService(repo repository.PaymentRepository) service.PaymentServiceServer {
	return &paymentService{repo: repo}
}

func (s *paymentService) GeneratePaymentLink(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {
	// Создаем запись о платеже в БД (в статусе "pending")
	payment := &entity.Payment{
		UserID:  req.UserId,
		OrderID: req.OrderId,
		Amount:  req.TotalPrice,
		Status:  "pending",
	}

	createdPayment, err := s.repo.CreatePayment(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment record: %v", err)
	}

	// Генерируем ссылку
	paymentURL := fmt.Sprintf("http://localhost:3000/pay?payment_id=%d", createdPayment.ID)

	return &paymentpb.PaymentResponse{
		PaymentUrl: paymentURL,
	}, nil
}

// Template for generating a payment link
// import (
// 	"github.com/stripe/stripe-go/v78"
// 	"github.com/stripe/stripe-go/v78/checkout/session"
// )

// func (s *paymentService) GeneratePaymentLink(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {
// 	stripe.Key = "sk_test_your_secret_key" // ⚠️ Укажи свой ключ Stripe

// 	// Создаем сессию для оплаты
// 	params := &stripe.CheckoutSessionParams{
// 		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
// 		LineItems: []*stripe.CheckoutSessionLineItemParams{
// 			{
// 				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
// 					Currency: stripe.String("usd"),
// 					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
// 						Name: stripe.String(fmt.Sprintf("Order #%d", req.OrderId)),
// 					},
// 					UnitAmount: stripe.Int64(int64(req.TotalPrice * 100)), // Цена в центах
// 				},
// 				Quantity: stripe.Int64(1),
// 			},
// 		},
// 		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
// 		SuccessURL: stripe.String(fmt.Sprintf("https://your-site.com/success?order_id=%d", req.OrderId)), // Куда перенаправить после оплаты
// 		CancelURL:  stripe.String(fmt.Sprintf("https://your-site.com/cancel?order_id=%d", req.OrderId)),  // Куда перенаправить при отмене
// 	}

// 	// Создаем сессию оплаты
// 	session, err := session.New(params)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create payment session: %v", err)
// 	}

// 	// Возвращаем ссылку на оплату
// 	return &paymentpb.PaymentResponse{
// 		PaymentUrl: session.URL,
// 	}, nil
// }
