package main

import (
	"fmt"
	"github.com/suzushin54/actor-based-inventory/actors"
	"log"
	"net/http"
	"os"

	_ "github.com/oklog/ulid/v2"
	inventoryv1 "github.com/suzushin54/actor-based-inventory/gen/inventory/v1"
	"github.com/suzushin54/actor-based-inventory/gen/inventory/v1/inventoryv1connect"
	"github.com/suzushin54/actor-based-inventory/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const httpServerAddr = "localhost:8080"

func main() {
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

	if err := http.ListenAndServe(
		httpServerAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		return
	}

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

func (s *InventoryServiceServer) CreateInventoryItem(
	r *inventoryv1.CreateInventoryRequest,
) (*inventoryv1.CreateInventoryResponse, error) {
	item := &actors.InventoryItem{
		ID:    r.Inventory.ProductId,
		Count: int(r.Inventory.Quantity),
	}

	if err := s.service.AddInventoryItem(item); err != nil {
		return nil, err
	}

	fmt.Printf("CreateInventoryItem: %v\n", r.Inventory)
	return nil, nil
}
