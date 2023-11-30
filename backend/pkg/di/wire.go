//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/shion0625/FYP/backend/pkg/api"
	"github.com/shion0625/FYP/backend/pkg/api/handler"
	"github.com/shion0625/FYP/backend/pkg/api/middleware"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/db"
	"github.com/shion0625/FYP/backend/pkg/repository"
	"github.com/shion0625/FYP/backend/pkg/service/cloud"
	"github.com/shion0625/FYP/backend/pkg/service/token"
	"github.com/shion0625/FYP/backend/pkg/usecase"
)

func InitializeApi(cfg *config.Config) (*http.ServerHTTP, error) {

	wire.Build(
		//db
		db.ConnectDatabase,

		//external
		token.NewTokenService,
		cloud.NewGCPCloudService,
		middleware.NewMiddleware,

		//repository
		repository.NewUserRepository,
		repository.NewAuthRepository,
		repository.NewProductRepository,
		repository.NewOrderRepository,

		//usecase
		usecase.NewAuthUseCase,
		usecase.NewUserUseCase,
		usecase.NewProductUseCase,
		usecase.NewOrderUseCase,

		//handler
		handler.NewAuthHandler,
		handler.NewUserHandler,
		handler.NewProductHandler,
		handler.NewOrderHandler,

		// server
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
