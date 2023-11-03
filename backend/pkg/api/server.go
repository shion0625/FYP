package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shion0625/FYP/backend/config/auth"
	"github.com/shion0625/FYP/backend/config/dig"
	"github.com/shion0625/FYP/backend/graphql/directives"
	"github.com/shion0625/FYP/backend/graphql/generated"
	"github.com/shion0625/FYP/backend/graphql/resolver"
	"github.com/shion0625/FYP/backend/util"
)

var timeout = 30 * time.Second

type ServerHTTP struct {
	Engine *echo.Engine
}

func NewServerHTTP() *ServerHTTP {
	util.LoadEnv()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// cors設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{
			os.Getenv("FRONTEND_URL"),
			os.Getenv("FRONTEND_URL"),
			os.Getenv("FRONTEND_DEVELOP_URL"),
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	d, _ := dig.BuildDigDependencies()
	err := d.Invoke(func() error {
		g := e.Group("/api")
		g.Use(echo.WrapMiddleware(auth.AuthMiddleware))

		return nil
	})

	if !errors.Is(err, nil) {
		panic(err)
	}

	port := util.GetPort()
	errPort := e.Start(":" + port)

	if !errors.Is(errPort, nil) {
		log.Fatalln(errPort)
	}

	engine := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: timeout,
	}

	return &ServerHTTP{Engine: engine}
}

func (s *ServerHTTP) Start() error {
	return engine.ListenAndServe()
}
