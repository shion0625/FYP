package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlerInterfaces "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	apiMiddleware "github.com/shion0625/FYP/backend/pkg/api/middleware"
	"github.com/shion0625/FYP/backend/pkg/api/router"
	"github.com/shion0625/FYP/backend/pkg/config"
)

var timeout = 30 * time.Second

type ServerHTTP struct {
	Engine *echo.Echo
}

func NewServerHTTP(
	cfg *config.Config,
	authHandler handlerInterfaces.AuthHandler,
	apiMiddleware apiMiddleware.Middleware,
) *ServerHTTP {
	engine := echo.New()

	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
	// cors設定
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{
			cfg.FrontendUrl,
			cfg.FrontendDevelopUrl,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
	engine.Use(apiMiddleware.Context)

	router.UserRoutes(engine.Group("/api"), authHandler)

	return &ServerHTTP{Engine: engine}
}

func (s *ServerHTTP) Start(cfg *config.Config) error {
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	return s.Engine.Start(":" + port)
}
