package kafka

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	Producer *kafka.Producer
}

func NewProducer() (*Producer, error) {
	brokerAddr := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddr,
	})
	if err != nil {
		return nil, err
	}

	return &Producer{Producer: p}, nil
}

func (p *Producer) SendMessage(topic string, key string, message []byte) error {
	return p.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topic, Partition: kafka.PartitionAny,
			},
			Key:   []byte(key),
			Value: message,
		}, nil)
}

func (p *Producer) Close() {
	p.Producer.Close()
}
