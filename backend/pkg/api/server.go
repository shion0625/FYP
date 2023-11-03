package api

import (
	"net/http"
	"os"
	"time"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/api/router"
)

var timeout = 30 * time.Second

type ServerHTTP struct {
	Engine *echo.Echo
}

func NewServerHTTP() *ServerHTTP {
	config.LoadEnv()

	engine := echo.New()

	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
	// cors設定
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
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

	router.UserRoutes(engine.Group("/api"))

	return &ServerHTTP{Engine: engine}
}

func (s *ServerHTTP) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return s.Engine.Start(":" + port)
}
