package adapters

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	order "github.com/suzushin54/actor-based-inventory/gen/order/v1"
	"github.com/suzushin54/actor-based-inventory/gen/order/v1/orderv1connect"
	"github.com/suzushin54/actor-based-inventory/internal/service"
)

// OrderServiceHandler implements the orderv1connect.OrderServiceHandler interface.
type OrderServiceHandler struct {
	logger  *slog.Logger
	service service.OrderService
	orderv1connect.UnimplementedOrderServiceHandler
}

func NewOrderServiceHandler(l *slog.Logger, s service.OrderService) *OrderServiceHandler {
	l = l.With("component", "OrderServiceHandler")

	return &OrderServiceHandler{
		logger:  l,
		service: s,
	}
}

func (s *OrderServiceHandler) CreateOrder(
	ctx context.Context,
	req *connect.Request[order.CreateOrderRequest],
) (*connect.Response[order.CreateOrderResponse], error) {
	s.logger.InfoContext(ctx, "CreateOrder", slog.Any("req", req))

	var orderProducts []service.OrderProduct
	for _, p := range req.Msg.Products {
		orderProducts = append(orderProducts, service.OrderProduct{
			ProductID: p.ProductId,
			Name:      p.Name,
			Price:     int(p.Price),
			Quantity:  int(p.Quantity),
		})
	}

	oid, err := s.service.CreateOrder(ctx, orderProducts)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(
		&order.CreateOrderResponse{
			OrderId: oid,
		},
	), nil
}
