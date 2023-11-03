package main

import (
	"log"

	"github.com/shion0625/FYP/backend/pkg/di"
)

func main() {

	server, err := di.InitializeApi()
	if err != nil {
		log.Fatal("Failed to initialize the api: ", err)
	}

	if server.Start(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
