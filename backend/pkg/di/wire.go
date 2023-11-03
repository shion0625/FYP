package dig

import (
	"github.com/google/wire"
	http "github.com/shion0625/FYP/backend/pkg/api"
	"github.com/shion0625/FYP/backend/pkg/db"

)

func InitializeApi() (*http.ServerHTTP, error) {

	wire.Build(
		db.ConnectDatabase,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
