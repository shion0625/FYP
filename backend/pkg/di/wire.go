//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/shion0625/FYP/backend/pkg/api"
	// "github.com/shion0625/FYP/backend/pkg/db"
	// "github.com/shion0625/FYP/backend/pkg/service/token"
)

func InitializeApi() (*http.ServerHTTP, error) {

	wire.Build(
		// db.ConnectDatabase,
		//external
		// token.NewTokenService,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
