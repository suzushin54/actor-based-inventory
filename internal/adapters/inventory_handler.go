package adapters

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/oklog/ulid/v2"
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

func NewInventoryServiceHandler(l *slog.Logger, s service.InventoryService) *InventoryServiceHandler {
	l = l.With("component", "InventoryServiceHandler")

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
	itemID, err := ulid.New(ulid.Now(), nil)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	item := &actors.InventoryItem{
		ID:    actors.InventoryID(itemID),
		Count: int(req.Msg.Quantity),
	}

	if err := s.service.AddInventoryItem(ctx, item); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(
		&inventoryv1.CreateInventoryResponse{
			ProductId: item.ID.String(),
		},
	), nil
}

func (s *InventoryServiceHandler) GetInventory(
	ctx context.Context,
	req *connect.Request[inventoryv1.GetInventoryRequest],
) (*connect.Response[inventoryv1.GetInventoryResponse], error) {
	s.logger.InfoContext(ctx, "GetInventory", slog.Any("req", req))
	itemID, err := ulid.Parse(req.Msg.ProductId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid ULID: %v", err))
	}

	item, err := s.service.QueryInventoryItem(ctx, actors.OrderID(itemID))
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("item not found: %v", err))
	}

	return connect.NewResponse(
		&inventoryv1.GetInventoryResponse{
			Inventory: &inventoryv1.Inventory{
				ProductId: item.ID.String(),
				Quantity:  int32(item.Count),
			},
		},
	), nil
}

func (s *InventoryServiceHandler) UpdateInventory(
	ctx context.Context,
	req *connect.Request[inventoryv1.UpdateInventoryRequest],
) (*connect.Response[inventoryv1.UpdateInventoryResponse], error) {
	s.logger.InfoContext(ctx, "UpdateInventory", slog.Any("req", req))
	itemID, err := ulid.Parse(req.Msg.Inventory.ProductId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid ULID: %v", err))
	}

	item := &actors.InventoryItem{
		ID:    actors.InventoryID(itemID),
		Count: int(req.Msg.Inventory.Quantity),
	}

	if err := s.service.UpdateInventoryItem(ctx, item); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(
		&inventoryv1.UpdateInventoryResponse{},
	), nil
}

func (s *InventoryServiceHandler) DeleteInventory(
	ctx context.Context,
	req *connect.Request[inventoryv1.DeleteInventoryRequest],
) (*connect.Response[inventoryv1.DeleteInventoryResponse], error) {
	s.logger.InfoContext(ctx, "DeleteInventory", slog.Any("req", req))
	itemID, err := ulid.Parse(req.Msg.ProductId)
	if err != nil {
		return nil, err
	}

	if err := s.service.RemoveInventoryItem(ctx, actors.OrderID(itemID)); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(
		&inventoryv1.DeleteInventoryResponse{},
	), nil
}
