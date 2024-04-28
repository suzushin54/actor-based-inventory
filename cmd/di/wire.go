//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/suzushin54/actor-based-inventory/internal/adapters"
	"github.com/suzushin54/actor-based-inventory/internal/infrastructure/server"
	"github.com/suzushin54/actor-based-inventory/internal/service"
	"net/http"
	"os"
)

func InitServer() (*server.Server, error) {
	wire.Build(
		//config.Load,
		service.NewInventoryService,
		adapters.NewInventoryServiceHandler,
		provideHTTPHandler,
		provideKafkaBrokerAddress,
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
