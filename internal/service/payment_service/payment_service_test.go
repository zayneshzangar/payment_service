package payment_service

import (
	"context"
	"errors"
	"testing"

	"payment_service/internal/entity"
	KafkaMocks "payment_service/internal/kafka/mocks"
	RepositoryMocks "payment_service/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPaymentService_PaymentProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := RepositoryMocks.NewMockPaymentRepository(ctrl)
	mockProducer := KafkaMocks.NewMockKafkaProducer(ctrl)
	ctx := context.Background()

	// Создаём service напрямую с моками, избегая NewPaymentService
	service := &paymentService{
		repo:     mockRepo,
		producer: mockProducer,
	}

	payment := &entity.Payment{
		ID:      1,
		UserID:  1,
		OrderID: 1,
		Amount:  100.0,
		Status:  "pending",
	}

	paymentRequest := entity.PaymentRequest{
		CardNum: "1234567890123456",
		CVV:     "123",
		ExpDate: "12/25",
		Name:    "John Doe",
	}

	t.Run("AlreadyPaid", func(t *testing.T) {
		// Подготовка
		paidPayment := &entity.Payment{
			ID:      1,
			UserID:  1,
			OrderID: 1,
			Amount:  100.0,
			Status:  "paid",
		}
		mockRepo.EXPECT().GetPayment(ctx, int64(1)).Return(paidPayment, nil)

		// Выполнение
		err := service.PaymentProcess(ctx, 1, paymentRequest)

		// Проверка
		assert.ErrorIs(t, err, errorAlreadyPaidFor)
	})

	t.Run("InvalidCard", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().GetPayment(ctx, int64(1)).Return(payment, nil)
		mockRepo.EXPECT().ValidateCard(ctx, paymentRequest.CardNum, paymentRequest.CVV, paymentRequest.ExpDate, paymentRequest.Name).
			Return(false, nil)

		// Выполнение
		err := service.PaymentProcess(ctx, 1, paymentRequest)

		// Проверка
		assert.EqualError(t, err, "invalid card details")
	})

	t.Run("InsufficientFunds", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().GetPayment(ctx, int64(1)).Return(payment, nil)
		mockRepo.EXPECT().ValidateCard(ctx, paymentRequest.CardNum, paymentRequest.CVV, paymentRequest.ExpDate, paymentRequest.Name).
			Return(true, nil)
		mockRepo.EXPECT().CheckAndDeductBalance(ctx, paymentRequest.CardNum, payment.Amount).Return(false, nil)

		// Выполнение
		err := service.PaymentProcess(ctx, 1, paymentRequest)

		// Проверка
		assert.EqualError(t, err, "insufficient funds")
	})

	t.Run("GetPaymentError", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().GetPayment(ctx, int64(1)).Return(nil, errors.New("database error"))

		// Выполнение
		err := service.PaymentProcess(ctx, 1, paymentRequest)

		// Проверка
		assert.EqualError(t, err, "database error")
	})

	t.Run("UpdatePaymentError", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().GetPayment(ctx, int64(1)).Return(payment, nil)
		mockRepo.EXPECT().ValidateCard(ctx, paymentRequest.CardNum, paymentRequest.CVV, paymentRequest.ExpDate, paymentRequest.Name).
			Return(true, nil)
		mockRepo.EXPECT().CheckAndDeductBalance(ctx, paymentRequest.CardNum, payment.Amount).Return(true, nil)
		mockRepo.EXPECT().UpdatePayment(ctx, &entity.Payment{
			ID:     payment.ID,
			Status: "paid",
		}).Return(errors.New("update error"))

		// Выполнение
		err := service.PaymentProcess(ctx, 1, paymentRequest)

		// Проверка
		assert.EqualError(t, err, "update error")
	})

	t.Run("KafkaProduceError", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().GetPayment(ctx, int64(1)).Return(payment, nil)
		mockRepo.EXPECT().ValidateCard(ctx, paymentRequest.CardNum, paymentRequest.CVV, paymentRequest.ExpDate, paymentRequest.Name).
			Return(true, nil)
		mockRepo.EXPECT().CheckAndDeductBalance(ctx, paymentRequest.CardNum, payment.Amount).Return(true, nil)
		mockRepo.EXPECT().UpdatePayment(ctx, &entity.Payment{
			ID:     payment.ID,
			Status: "paid",
		}).Return(nil)

		// Настройка Kafka
		mockProducer.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(errors.New("kafka error"))

		// Выполнение
		err := service.PaymentProcess(ctx, 1, paymentRequest)

		// Проверка
		assert.EqualError(t, err, "kafka error")
	})
}

func TestPaymentService_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := RepositoryMocks.NewMockPaymentRepository(ctrl)
	mockProducer := KafkaMocks.NewMockKafkaProducer(ctrl)

	// Создаём service напрямую с моками
	service := &paymentService{
		repo:     mockRepo,
		producer: mockProducer,
	}

	// Подготовка
	mockProducer.EXPECT().Flush(flushTimeout).Return(0)
	mockProducer.EXPECT().Close()

	// Выполнение
	service.Close()

	// Проверка: метод Close не возвращает ошибок, проверяем вызовы через моки
}
