package service

import (
	"github.com/asynkron/protoactor-go/actor"
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

func (s *Service) RemoveInventoryItem(itemID string) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.RemoveInventoryItem{ItemID: itemID})
}
