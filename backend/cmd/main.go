package main

import (
	"log"

	"github.com/shion0625/FYP/backend/pkg/di"
)

func main() {

	server, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatal("Failed to initialize the api: ", err)
	}

	d, _ := dig.BuildDigDependencies()
	err := d.Invoke(func(r *resolver.Resolver) error {
		e.GET("/", Playground())
		g := e.Group("/api")
		g.Use(echo.WrapMiddleware(auth.AuthMiddleware))
		return nil
	})

	if server.Start(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
