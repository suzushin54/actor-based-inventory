// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/suzushin54/actor-based-inventory/internal/adapters"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure/server"
	"github.com/suzushin54/actor-based-inventory/internal/service"
	"net/http"
	"os"
)

// Injectors from wire.go:

func InitServer() (*server.Server, error) {
	string2 := provideKafkaBrokerAddress()
	inventoryService, err := service.NewInventoryService(string2)
	if err != nil {
		return nil, err
	}
	inventoryServiceHandler := adapters.NewInventoryServiceHandler(inventoryService)
	serveMux := server.ConfigureMux(inventoryServiceHandler)
	handler := provideHTTPHandler(serveMux)
	serverServer := server.NewServer(handler)
	return serverServer, nil
}

// wire.go:

func provideHTTPHandler(mux *http.ServeMux) http.Handler {
	return mux
}

func provideKafkaBrokerAddress() string {
	return os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}
