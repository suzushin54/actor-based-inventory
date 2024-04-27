package di

import (
	"github.com/google/wire"
	"github.com/suzushin54/actor-based-inventory/cmd/config"
	"github.com/suzushin54/actor-based-inventory/service"
)

func Init() (*service.Service, error) {
	wire.Build(
		config.Load,
		service.NewService,
		wire.Struct(new(service.Service), "*"),
	)

	return nil, nil
}
