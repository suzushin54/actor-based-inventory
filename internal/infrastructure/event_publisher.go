package infrastructure

import (
	"github.com/suzushin54/actor-based-inventory/pkg/kafka"
)

type EventPublisher struct {
	producer *kafka.Producer
}

func NewEventPublisher() (*EventPublisher, error) {
	producer, err := kafka.NewProducer()
	if err != nil {
		return nil, err
	}
	return &EventPublisher{producer: producer}, nil
}

func (ep *EventPublisher) Publish(topic, key string, message []byte) error {
	return ep.producer.SendMessage(topic, key, message)
}

func (ep *EventPublisher) Close() {
	ep.producer.Close()
}
