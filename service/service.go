package service

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/actors"
)

type Service struct {
	actorSystem       *actor.ActorSystem
	inventoryActorPID *actor.PID
}

func NewService() *Service {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(
		func() actor.Actor {
			return actors.NewInventoryActor()
		},
	)

	pid := system.Root.Spawn(props)

	return &Service{
		actorSystem:       system,
		inventoryActorPID: pid,
	}
}

func (s *Service) AddInventoryItem(item *actors.InventoryItem) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.AddInventoryItem{Item: item})
}

func (s *Service) RemoveInventoryItem(itemID ulid.ULID) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.RemoveInventoryItem{ItemID: itemID.String()})
}

func (s *Service) UpdateInventoryItemCount(itemID ulid.ULID, count int) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.UpdateInventoryItemCount{ItemID: itemID.String(), Count: count})
}

func (s *Service) QueryInventoryItem(itemID ulid.ULID) *actors.InventoryItem {
	future := s.actorSystem.Root.RequestFuture(s.inventoryActorPID, &actors.QueryInventoryItem{ItemID: itemID.String()}, 1000)
	result, _ := future.Result()
	return result.(*actors.InventoryItem)
}
