package interfaces

import (
	"context"

	"github.com/shion0625/backend/pkg/api/handler/request"
	"github.com/shion0625/backend/pkg/api/handler/response"
)

type StockRepository interface {
	FindAll(ctx context.Context, pagination request.Pagination) (stocks []response.Stock, err error)
	Update(ctx context.Context, updateValues request.UpdateStock) error
}
