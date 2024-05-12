package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/internal/actors"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure"
)

type InventoryService interface {
	AddInventoryItem(ctx context.Context, item *actors.InventoryItem) error
	RemoveInventoryItem(ctx context.Context, itemID ulid.ULID)
	UpdateInventoryItemCount(ctx context.Context, itemID ulid.ULID, count int)
	QueryInventoryItem(ctx context.Context, itemID ulid.ULID) (*actors.InventoryItem, error)
}

type inventoryService struct {
	logger            *slog.Logger
	actorSystem       *actor.ActorSystem
	inventoryActorPID *actor.PID
	eventPublisher    *infrastructure.EventPublisher
}

func NewInventoryService(l *slog.Logger, as *actor.ActorSystem, ep *infrastructure.EventPublisher) InventoryService {
	l = l.With("component", "InventoryService")

	props := actor.PropsFromProducer(
		func() actor.Actor {
			return actors.NewInventoryActor()
		},
	)
	pid := as.Root.Spawn(props)
	return &inventoryService{
		logger:            l,
		actorSystem:       as,
		inventoryActorPID: pid,
		eventPublisher:    ep,
	}
}

func (s *inventoryService) AddInventoryItem(ctx context.Context, item *actors.InventoryItem) error {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.AddInventoryItem{Item: item})

	message := fmt.Sprintf("Added item %s to inventory", item.ID)
	if err := s.eventPublisher.Publish("inventory", item.ID, []byte(message)); err != nil {
		return err
	}

	s.logger.InfoContext(ctx, "successfully added item to inventory", slog.String("item_id", item.ID), slog.Int("count", item.Count))
	return nil
}

func (s *inventoryService) RemoveInventoryItem(ctx context.Context, itemID ulid.ULID) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.RemoveInventoryItem{ItemID: itemID.String()})

	s.logger.InfoContext(ctx, "removed item from inventory", slog.String("item_id", itemID.String()))
}

func (s *inventoryService) UpdateInventoryItemCount(ctx context.Context, itemID ulid.ULID, count int) {
	s.actorSystem.Root.Send(s.inventoryActorPID, &actors.UpdateInventoryItemCount{
		ItemID: itemID.String(),
		Count:  count,
	})

	s.logger.InfoContext(ctx, "updated item count in inventory", slog.String("item_id", itemID.String()), slog.Int("count", count))
}

func (s *inventoryService) QueryInventoryItem(ctx context.Context, itemID ulid.ULID) (*actors.InventoryItem, error) {
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

	s.logger.InfoContext(ctx, "queried item from inventory", slog.String("item_id", itemID.String()), slog.Int("count", res.Count))
	return res, nil
}
