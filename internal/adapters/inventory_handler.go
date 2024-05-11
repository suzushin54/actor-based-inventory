package adapters

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	inventoryv1 "github.com/suzushin54/actor-based-inventory/gen/inventory/v1"
	"github.com/suzushin54/actor-based-inventory/gen/inventory/v1/inventoryv1connect"
	"github.com/suzushin54/actor-based-inventory/internal/actors"
	"github.com/suzushin54/actor-based-inventory/internal/service"
)

// InventoryServiceHandler implements the inventoryv1connect.InventoryServiceHandler interface.
type InventoryServiceHandler struct {
	logger  *slog.Logger
	service service.InventoryService
	inventoryv1connect.UnimplementedInventoryServiceHandler
}

func NewInventoryServiceHandler(
	l *slog.Logger,
	s service.InventoryService,
) *InventoryServiceHandler {
	return &InventoryServiceHandler{
		logger:  l,
		service: s,
	}
}

func (s *InventoryServiceHandler) CreateInventory(
	ctx context.Context,
	req *connect.Request[inventoryv1.CreateInventoryRequest],
) (*connect.Response[inventoryv1.CreateInventoryResponse], error) {
	s.logger.InfoContext(ctx, "CreateInventory", slog.Any("req", req))
	item := &actors.InventoryItem{
		ID:    req.Msg.Inventory.ProductId,
		Count: int(req.Msg.Inventory.Quantity),
	}

	if err := s.service.AddInventoryItem(ctx, item); err != nil {
		return nil, err
	}

	return connect.NewResponse(
		&inventoryv1.CreateInventoryResponse{},
	), nil
}
