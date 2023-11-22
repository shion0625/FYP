package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type ProductUseCase interface {
	FindAllCategories(ctx echo.Context, pagination request.Pagination) ([]response.Category, error)
	SaveCategory(ctx echo.Context, categoryName string) error

	// variations
	SaveVariation(ctx echo.Context, categoryID uint, variationNames []string) error
	SaveVariationOption(ctx echo.Context, variationID uint, variationOptionValues []string) error

	FindAllVariationsAndItsValues(ctx echo.Context, categoryID uint) ([]response.Variation, error)

	// products
	FindAllProducts(ctx echo.Context, pagination request.Pagination, categoryID *uint, brandID *uint) (products []response.Product, err error)
	SaveProduct(ctx echo.Context, product request.Product) error
	UpdateProduct(ctx echo.Context, product domain.Product) error

	SaveProductItem(ctx echo.Context, productID uint, productItem request.ProductItem) error
	FindAllProductItems(ctx echo.Context, productID uint) ([]response.ProductItems, error)
}
