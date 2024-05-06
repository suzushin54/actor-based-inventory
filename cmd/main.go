package main

import (
	"log/slog"
	"os"

	_ "github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/cmd/di"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("main process started")

	svr, err := di.InitServer(logger)
	if err != nil {
		logger.Error("failed to initialize service", err)
	}

	if err := svr.Start(); err != nil {
		logger.Error("failed to start server", err)
	}
}
