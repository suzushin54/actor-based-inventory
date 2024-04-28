package service

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/internal/actors"
	"github.com/suzushin54/actor-based-inventory/pkg/kafka"
)

type InventoryService struct {
	actorSystem       *actor.ActorSystem
	inventoryActorPID *actor.PID
	kafkaProducer     *kafka.Producer
}

func NewInventoryService(brokers string) (*InventoryService, error) {

	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(
		func() actor.Actor {
			return actors.NewInventoryActor()
		},
	)

	pid := system.Root.Spawn(props)

	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		return nil, err
	}

	return &InventoryService{
		actorSystem:       system,
		inventoryActorPID: pid,
		kafkaProducer:     kafkaProducer,
	}, nil
}

func (s *InventoryService) AddInventoryItem(item *actors.InventoryItem) error {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.AddInventoryItem{Item: item})

	message := fmt.Sprintf("Added item %s to inventory", item.ID)
	if err := s.kafkaProducer.SendMessage("inventory", item.ID, []byte(message)); err != nil {
		return err
	}

	fmt.Printf("Added item %s to inventory\n", item.ID)
	return nil
}

func (s *InventoryService) RemoveInventoryItem(itemID ulid.ULID) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.RemoveInventoryItem{ItemID: itemID.String()})
}

func (s *InventoryService) UpdateInventoryItemCount(itemID ulid.ULID, count int) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.UpdateInventoryItemCount{
		ItemID: itemID.String(),
		Count:  count,
	})
}

func (s *InventoryService) QueryInventoryItem(itemID ulid.ULID) (*actors.InventoryItem, error) {
	future := s.actorSystem.Root.RequestFuture(s.inventoryActorPID, &actors.QueryInventoryItem{
		ItemID: itemID.String(),
	}, 5*time.Second)

	result, err := future.Result()
	if err != nil {
		return nil, err
	}

	res, ok := result.(*actors.InventoryItem)
	if !ok {
		return nil, fmt.Errorf("unexpected response: %v", result)
	}

	return res, nil
}
