package api

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handlerInterfaces "github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	apiMiddleware "github.com/shion0625/FYP/backend/pkg/api/middleware"
	"github.com/shion0625/FYP/backend/pkg/api/router"
	"github.com/shion0625/FYP/backend/pkg/config"
)

type ServerHTTP struct {
	Engine *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func NewServerHTTP(
	cfg *config.Config,
	apiMiddleware apiMiddleware.Middleware,
	authHandler handlerInterfaces.AuthHandler,
	userHandler handlerInterfaces.UserHandler,
	productHandler handlerInterfaces.ProductHandler,
	orderHandler handlerInterfaces.OrderHandler,
) *ServerHTTP {
	engine := echo.New()
	engine.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}

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
	engine.Use(apiMiddleware.AccessControlExposeHeaders)

	router.UserRoutes(engine.Group("/api"), apiMiddleware, authHandler, userHandler, productHandler, orderHandler)

	return &ServerHTTP{Engine: engine}
}

func (s *ServerHTTP) Start(cfg *config.Config) error {
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	if err := s.Engine.Start("127.0.0.1:" + port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
