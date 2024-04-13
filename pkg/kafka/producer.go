package kafka

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Producer struct {
	Producer *kafka.Producer
}

func NewProducer(brokers string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
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
