//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/shion0625/backend/pkg/api"
	"github.com/shion0625/backend/pkg/api/handler"
	"github.com/shion0625/backend/pkg/api/middleware"
	"github.com/shion0625/backend/pkg/config"
	"github.com/shion0625/backend/pkg/db"
	"github.com/shion0625/backend/pkg/repository"
	"github.com/shion0625/backend/pkg/service/cloud"
	"github.com/shion0625/backend/pkg/service/otp"
	"github.com/shion0625/backend/pkg/service/token"
	"github.com/shion0625/backend/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {

	wire.Build(db.ConnectDatabase,
		//external
		token.NewTokenService,
		otp.NewOtpAuth,
		cloud.NewAWSCloudService,

		// repository

		middleware.NewMiddleware,
		repository.NewAuthRepository,
		repository.NewPaymentRepository,
		repository.NewAdminRepository,
		repository.NewUserRepository,
		repository.NewCartRepository,
		repository.NewProductRepository,
		repository.NewOrderRepository,
		repository.NewCouponRepository,
		repository.NewOfferRepository,
		repository.NewStockRepository,
		repository.NewBrandDatabaseRepository,

		//usecase
		usecase.NewAuthUseCase,
		usecase.NewAdminUseCase,
		usecase.NewUserUseCase,
		usecase.NewCartUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewProductUseCase,
		usecase.NewOrderUseCase,
		usecase.NewCouponUseCase,
		usecase.NewOfferUseCase,
		usecase.NewStockUseCase,
		usecase.NewBrandUseCase,
		// handler
		handler.NewAuthHandler,
		handler.NewAdminHandler,
		handler.NewUserHandler,
		handler.NewCartHandler,
		handler.NewPaymentHandler,
		handler.NewProductHandler,
		handler.NewOrderHandler,
		handler.NewCouponHandler,
		handler.NewOfferHandler,
		handler.NewStockHandler,
		handler.NewBrandHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
