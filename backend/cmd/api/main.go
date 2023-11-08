package main

import (
	"log"

	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error to load the config: ", err)
	}

	server, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatal("Failed to initialize the api: ", err)
	}

	if server.Start(cfg); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
