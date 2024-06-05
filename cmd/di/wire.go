//go:build wireinject
// +build wireinject

package di

import (
	"log/slog"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/google/wire"
	"github.com/suzushin54/actor-based-inventory/internal/adapters"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure/server"
	"github.com/suzushin54/actor-based-inventory/internal/service"
)

func InitServer(logger *slog.Logger) (*server.Server, error) {
	wire.Build(
		NewActorSystem,

		adapters.NewOrderServiceHandler,
		//service.NewOrderService,

		adapters.NewInventoryServiceHandler,
		service.NewInventoryService,

		// server
		infrastructure.NewEventPublisher,
		provideHTTPHandler,
		server.ConfigureMux,
		server.NewServer,
	)
	return nil, nil
}

func provideHTTPHandler(mux *http.ServeMux) http.Handler {
	return mux
}

func NewActorSystem() *actor.ActorSystem {
	return actor.NewActorSystem()
}
