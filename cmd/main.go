package main

import (
	"fmt"
	"log"

	_ "github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/cmd/di"
)

func main() {
	log.Printf("main process started\n")

	svr, err := di.InitServer()
	if err != nil {
		log.Fatalf("failed to initialize service: %v", err)
	}

	if err := svr.Start(); err != nil {
		fmt.Printf("failed to terminate server: %v", err)
	}
}
