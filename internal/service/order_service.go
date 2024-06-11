package service

import (
	"context"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/internal/actors"
	"log/slog"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure"
)

type OrderProduct struct {
	ProductID string
	Name      string
	Price     int
	Quantity  int
}

type Order struct {
	OrderID   string
	Products  []OrderProduct
	Status    string
	CreatedAt string
}

type OrderService interface {
	CreateOrder(ctx context.Context, products []OrderProduct) (string, error)
	QueryOrder(ctx context.Context, orderID string) (*Order, error)
}

type orderService struct {
	logger         *slog.Logger
	actorSystem    *actor.ActorSystem
	orderActorPID  *actor.PID
	eventPublisher *infrastructure.EventPublisher
}

func NewOrderService(
	l *slog.Logger,
	as *actor.ActorSystem,
	ep *infrastructure.EventPublisher,
) OrderService {
	l = l.With("component", "OrderService")

	props := actor.PropsFromProducer(
		func() actor.Actor {
			return actors.NewInventoryActor()
		},
	)
	pid := as.Root.Spawn(props)
	return &orderService{
		logger:         l,
		actorSystem:    as,
		orderActorPID:  pid,
		eventPublisher: ep,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, products []OrderProduct) (string, error) {
	orderID, err := ulid.New(ulid.Now(), nil)
	if err != nil {
		return "", err
	}

	// service.OrderProduct -> actors.OrderProduct
	var aops []*actors.OrderProduct
	for _, p := range products {
		aops = append(aops, &actors.OrderProduct{
			ProductID: p.ProductID,
			Name:      p.Name,
			Price:     p.Price,
			Quantity:  p.Quantity,
		})
	}

	order := &actors.Order{
		ID:        actors.OrderID(orderID),
		Products:  aops,
		Status:    "created",
		CreatedAt: time.Now().String(),
	}

	future := s.actorSystem.Root.RequestFuture(s.orderActorPID, &actors.CreateOrder{Order: order}, 5*time.Second)

	_, err = future.Result()
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("Order created: %s", orderID.String())
	if err := s.eventPublisher.Publish("order_events", orderID.String(), []byte(message)); err != nil {
		return "", err
	}

	s.logger.InfoContext(ctx, "successfully created order", slog.String("order_id", orderID.String()))

	return orderID.String(), nil
}

func (o orderService) QueryOrder(ctx context.Context, orderID string) (*Order, error) {
	//TODO implement me
	panic("implement me")
}
