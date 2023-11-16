// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/shion0625/FYP/backend/pkg/api"
	"github.com/shion0625/FYP/backend/pkg/api/handler"
	"github.com/shion0625/FYP/backend/pkg/api/middleware"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/db"
	"github.com/shion0625/FYP/backend/pkg/repository"
	"github.com/shion0625/FYP/backend/pkg/service/token"
	"github.com/shion0625/FYP/backend/pkg/usecase"
)

// Injectors from wire.go:

func InitializeApi(cfg *config.Config) (*api.ServerHTTP, error) {
	tokenService := token.NewTokenService(cfg)
	middlewareMiddleware := middleware.NewMiddleware(tokenService)
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	authRepository := repository.NewAuthRepository(gormDB)
	authUseCase := usecase.NewAuthUseCase(userRepository, tokenService, authRepository)
	authHandler := handler.NewAuthHandler(authUseCase, cfg)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := api.NewServerHTTP(cfg, middlewareMiddleware, authHandler, userHandler)
	return serverHTTP, nil
}
