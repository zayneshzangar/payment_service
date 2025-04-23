package payment_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"payment_service/internal/entity"
	LocalKafka "payment_service/internal/kafka"
	"payment_service/internal/repository"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var errorUnknownType = errors.New("unknown event type")
var errorAlreadyPaidFor = errors.New("already paid for")

const (
	flushTimeout = 5000
)

type paymentService struct {
	repo     repository.PaymentRepository
	producer LocalKafka.KafkaProducer
}

func NewPaymentService(repo repository.PaymentRepository, kafkaBootstrapServers []string) (*paymentService, error) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(kafkaBootstrapServers, ","),
	}

	producer, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, fmt.Errorf("error with new producer: %w", err)
	}

	return &paymentService{repo: repo, producer: producer}, nil
}

func (s *paymentService) PaymentProcess(ctx context.Context, paymentID int64, req entity.PaymentRequest) error {
	// Получаем payment из БД
	payment, err := s.repo.GetPayment(ctx, paymentID)
	if err != nil {
		return err
	}

	if payment.Status == "paid" {
		return errorAlreadyPaidFor
	}

	// Проверяем карту
	valid, err := s.repo.ValidateCard(ctx, req.CardNum, req.CVV, req.ExpDate, req.Name)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid card details")
	}

	// Проверяем баланс
	hasFunds, err := s.repo.CheckAndDeductBalance(ctx, req.CardNum, payment.Amount)
	if err != nil {
		return err
	}
	if !hasFunds {
		return errors.New("insufficient funds")
	}

	// Обновляем только статус
	err = s.repo.UpdatePayment(ctx, &entity.Payment{
		ID:     payment.ID,
		Status: "paid",
	})
	if err != nil {
		return err
	}

	// Отправляем уведомление в Order Service
	err = s.notifyOrderService(payment)
	if err != nil {
		return err
	}

	return nil
}

func (s *paymentService) notifyOrderService(payment *entity.Payment) error {
	topic := "payment_events"

	notification := entity.PaymentNotification{
		OrderID: payment.OrderID,
		Status:  "paid",
	}

	// Сериализуем структуру в JSON
	messageJSON, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	kafkamsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     []byte(messageJSON),
		Timestamp: time.Now(),
		Headers: []kafka.Header{
			{
				Key:   "content-type",
				Value: []byte("text/plain"),
			},
		},
	}

	kafkaChan := make(chan kafka.Event)
	if err := s.producer.Produce(kafkamsg, kafkaChan); err != nil {
		return err
	}

	e := <-kafkaChan
	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return errorUnknownType
	}
}

func (s *paymentService) Close() {
	s.producer.Flush(flushTimeout)
	s.producer.Close()
}