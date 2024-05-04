//go:build wireinject
// +build wireinject

package di

import (
	"net/http"
	"os"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/google/wire"
	"github.com/suzushin54/actor-based-inventory/internal/adapters"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure/server"
	"github.com/suzushin54/actor-based-inventory/internal/service"
)

func InitServer() (*server.Server, error) {
	wire.Build(
		NewActorSystem,

		service.NewInventoryService,
		adapters.NewInventoryServiceHandler,
		infrastructure.NewEventPublisher,
		// server
		provideHTTPHandler,
		server.ConfigureMux,
		server.NewServer,
	)
	return nil, nil
}

func provideHTTPHandler(mux *http.ServeMux) http.Handler {
	return mux
}

func provideKafkaBrokerAddress() string {
	return os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}

func NewActorSystem() *actor.ActorSystem {
	return actor.NewActorSystem()
}
