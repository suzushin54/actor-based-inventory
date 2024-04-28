package server

import (
	"net/http"

	"github.com/suzushin54/actor-based-inventory/gen/inventory/v1/inventoryv1connect"
	"github.com/suzushin54/actor-based-inventory/internal/adapters"
)

// ConfigureMux sets up the HTTP routing for the application.
func ConfigureMux(ish *adapters.InventoryServiceHandler) *http.ServeMux {
	mux := http.NewServeMux()
	path, handler := inventoryv1connect.NewInventoryServiceHandler(ish)
	mux.Handle(path, handler)

	return mux
}
