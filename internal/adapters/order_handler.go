package adapters

import (
	"connectrpc.com/connect"
	"context"
	"github.com/oklog/ulid/v2"
	order "github.com/suzushin54/actor-based-inventory/gen/order/v1"
	"github.com/suzushin54/actor-based-inventory/gen/order/v1/orderv1connect"
	"github.com/suzushin54/actor-based-inventory/internal/service"
	"log/slog"
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

	orderID, err := ulid.New(ulid.Now(), nil)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// サービスレイヤに注文を作成させる
	var orderProducts []service.OrderProduct
	for _, p := range req.Msg.Products {
		orderProducts = append(orderProducts, service.OrderProduct{
			ProductID: p.ProductId,
			Name:      p.Name,
			Price:     float64(p.Price),
			Quantity:  int(p.Quantity),
		})
	}

	if err = s.service.CreateOrder(ctx, orderID.String(), orderProducts); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(
		&order.CreateOrderResponse{
			OrderId: orderID.String(),
		},
	), nil
}
