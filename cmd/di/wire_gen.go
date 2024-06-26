// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/suzushin54/actor-based-inventory/internal/adapters"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure/server"
	"github.com/suzushin54/actor-based-inventory/internal/service"
	"log/slog"
	"net/http"
)

// Injectors from wire.go:

func InitServer(logger *slog.Logger) (*server.Server, error) {
	actorSystem := NewActorSystem()
	eventPublisher, err := infrastructure.NewEventPublisher()
	if err != nil {
		return nil, err
	}
	inventoryService := service.NewInventoryService(logger, actorSystem, eventPublisher)
	inventoryServiceHandler := adapters.NewInventoryServiceHandler(logger, inventoryService)
	serveMux := server.ConfigureMux(inventoryServiceHandler)
	handler := provideHTTPHandler(serveMux)
	serverServer := server.NewServer(handler)
	return serverServer, nil
}

// wire.go:

func provideHTTPHandler(mux *http.ServeMux) http.Handler {
	return mux
}

func NewActorSystem() *actor.ActorSystem {
	return actor.NewActorSystem()
}
