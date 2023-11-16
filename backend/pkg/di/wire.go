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
	"github.com/shion0625/FYP/backend/pkg/service/token"
	"github.com/shion0625/FYP/backend/pkg/usecase"
)

func InitializeApi(cfg *config.Config) (*http.ServerHTTP, error) {

	wire.Build(
		//db
		db.ConnectDatabase,

		//external
		token.NewTokenService,
		middleware.NewMiddleware,

		//repository
		repository.NewUserRepository,
		repository.NewAuthRepository,

		//usecase
		usecase.NewAuthUseCase,
		usecase.NewUserUseCase,

		//handler
		handler.NewAuthHandler,
		handler.NewUserHandler,

		// server
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
