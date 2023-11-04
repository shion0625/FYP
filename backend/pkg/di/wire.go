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
	"github.com/shion0625/FYP/backend/pkg/usecase"
)

func InitializeApi(cfg *config.Config) (*http.ServerHTTP, error) {

	wire.Build(
		//external
		middleware.NewMiddleware,

		//db
		db.ConnectDatabase,

		//repository
		repository.NewUserRepository,

		//usecase
		usecase.NewAuthUseCase,

		//handler
		handler.NewAuthHandler,

		// server
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
