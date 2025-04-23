package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// KafkaProducer определяет методы, необходимые для отправки сообщений в Kafka.
type KafkaProducer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	Flush(timeoutMs int) int
	Close()
}