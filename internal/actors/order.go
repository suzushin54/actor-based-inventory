package actors

import (
	"log/slog"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/oklog/ulid/v2"
)

// OrderID represents a unique identifier.
type OrderID ulid.ULID

func (id OrderID) String() string {
	return ulid.ULID(id).String()
}

// OrderProduct represents a product in an order.
type OrderProduct struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

// Order represents an order.
type Order struct {
	ID        OrderID
	Products  []OrderProduct
	Status    string
	CreatedAt string
}

// OrderActor is an actor that manages orders.
type OrderActor struct {
	Orders map[OrderID]*Order
}

type QueryOrder struct {
	OrderID OrderID
}

type CreateOrder struct {
	Order *Order
}

func NewOrderActor() *OrderActor {
	return &OrderActor{
		Orders: make(map[OrderID]*Order),
	}
}

func (actor *OrderActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *CreateOrder:
		actor.Orders[msg.Order.ID] = msg.Order
		slog.Info("Created order in order actor", "order_id", msg.Order.ID)
	case *QueryOrder:
		order, ok := actor.Orders[msg.OrderID]
		if !ok {
			slog.Warn("Order not found", "order_id", msg.OrderID)
			return
		}
		slog.Info("Order found in order actor", "order_id", msg.OrderID)
		ctx.Respond(order)
	}
}
