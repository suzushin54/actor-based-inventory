package infrastructure

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type EventPublisher struct {
	producer *kafka.Producer
}

// NewEventPublisher creates a new instance of EventPublisher with its own Kafka producer.
func NewEventPublisher() (*EventPublisher, error) {
	brokerAddr := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddr,
	})
	if err != nil {
		return nil, err
	}
	return &EventPublisher{producer: producer}, nil
}

// Publish sends a message to the specified Kafka topic.
func (ep *EventPublisher) Publish(topic, key string, message []byte) error {
	return ep.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          message,
	}, nil)
}

// Close shuts down the Kafka producer.
func (ep *EventPublisher) Close() {
	ep.producer.Close()
}
