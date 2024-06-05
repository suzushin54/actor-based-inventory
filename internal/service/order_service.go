package service

import (
	"context"
	"log/slog"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure"
)

type OrderProduct struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Order struct {
	OrderID   string
	Products  []OrderProduct
	Status    string
	CreatedAt string
}

type OrderService interface {
	CreateOrder(ctx context.Context, orderID string, products []OrderProduct) error
	QueryOrder(ctx context.Context, orderID string) (*Order, error)
}

type orderService struct {
	logger         *slog.Logger
	actorSystem    *actor.ActorSystem
	orderActorPID  *actor.PID
	eventPublisher *infrastructure.EventPublisher
}

func NewOrderService(
	logger *slog.Logger,
	actorSystem *actor.ActorSystem,
	orderActorPID *actor.PID,
	eventPublisher *infrastructure.EventPublisher,
) OrderService {
	return &orderService{
		logger:         logger,
		actorSystem:    actorSystem,
		orderActorPID:  orderActorPID,
		eventPublisher: eventPublisher,
	}
}

func (o orderService) CreateOrder(ctx context.Context, orderID string, products []OrderProduct) error {
	//TODO implement me
	panic("implement me")
}

func (o orderService) QueryOrder(ctx context.Context, orderID string) (*Order, error) {
	//TODO implement me
	panic("implement me")
}
