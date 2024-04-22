package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	_ "github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/actors"
	inventoryv1 "github.com/suzushin54/actor-based-inventory/gen/inventory/v1"
	"github.com/suzushin54/actor-based-inventory/gen/inventory/v1/inventoryv1connect"
	"github.com/suzushin54/actor-based-inventory/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// const httpServerAddr = "localhost:8080"
const httpServerAddr = "0.0.0.0:8080"

func main() {
	log.Printf("main process started\n")

	brokerAddr := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	s, err := service.NewService(brokerAddr, "inventory")
	if err != nil {
		log.Fatal(err)
	}
	//id := ulid.Make()

	mux := http.NewServeMux()
	iss := NewInventoryServiceServer(s)
	path, handler := inventoryv1connect.NewInventoryServiceHandler(iss)
	mux.Handle(path, handler)

	log.Printf("HTTP server starting on %s", httpServerAddr)
	if err := http.ListenAndServe(
		httpServerAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
	log.Printf("HTTP server listening on %s", httpServerAddr)
}

// InventoryServiceServer implements the inventoryv1connect.InventoryServiceServer interface.
type InventoryServiceServer struct {
	service *service.Service
	inventoryv1connect.UnimplementedInventoryServiceHandler
}

func NewInventoryServiceServer(s *service.Service) *InventoryServiceServer {
	return &InventoryServiceServer{
		service: s,
	}
}

func (s *InventoryServiceServer) CreateInventory(
	ctx context.Context,
	req *connect.Request[inventoryv1.CreateInventoryRequest],
) (*connect.Response[inventoryv1.CreateInventoryResponse], error) {
	log.Printf("CreateInventory: %v\n", req.Msg)

	item := &actors.InventoryItem{
		ID:    req.Msg.Inventory.ProductId,
		Count: int(req.Msg.Inventory.Quantity),
	}

	if err := s.service.AddInventoryItem(item); err != nil {
		return nil, err
	}

	res := &inventoryv1.CreateInventoryResponse{}
	return connect.NewResponse(res), nil
}
