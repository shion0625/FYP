package interfaces

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type ProductRepository interface {
	Transactions(ctx echo.Context, trxFn func(repo ProductRepository) error) error

	// category
	IsCategoryNameExist(ctx echo.Context, categoryName string) (bool, error)
	FindAllMainCategories(ctx echo.Context, pagination request.Pagination) ([]response.Category, error)
	SaveCategory(ctx echo.Context, categoryName string) error

	// variation
	IsVariationNameExistForCategory(ctx echo.Context, name string, categoryID uint) (bool, error)
	SaveVariation(ctx echo.Context, categoryID uint, variationName string) error
	FindAllVariationsByCategoryID(ctx echo.Context, categoryID uint) ([]response.Variation, error)

	// variation values
	IsVariationValueExistForVariation(ctx echo.Context, value string, variationID uint) (exist bool, err error)
	SaveVariationOption(ctx echo.Context, variationID uint, variationValue string) error
	FindAllVariationOptionsByVariationID(ctx echo.Context, variationID uint) ([]response.VariationOption, error)

	FindAllVariationValuesOfProductItem(ctx echo.Context, productItemID uint) ([]response.ProductVariationValue, error)

	FindProductByID(ctx echo.Context, productID uint) (product response.Product, err error)
	IsProductNameExistForOtherProduct(ctx echo.Context, name string, productID uint) (bool, error)
	IsProductNameExist(ctx echo.Context, productName string) (exist bool, err error)

	FindAllProducts(ctx echo.Context, pagination request.Pagination, categoryID *uint, brandID *uint) ([]response.Product, error)

	SaveProduct(ctx echo.Context, product domain.Product) error
	UpdateProduct(ctx echo.Context, product domain.Product) error

	// product items
	FindProductItemByID(ctx echo.Context, productItemID uint) (domain.ProductItem, error)
	FindAllProductItems(ctx echo.Context, productID uint) ([]response.ProductItemsDB, error)
	FindVariationCountForProduct(ctx echo.Context, productID uint) (variationCount uint, err error) // to check the product config already exist
	FindAllProductItemIDsByProductIDAndVariationOptionID(ctx context.Context, productID, variationOptionID uint) ([]uint, error)
	SaveProductConfiguration(ctx echo.Context, productItemID, variationOptionID uint) error
	SaveProductItem(ctx echo.Context, productItem domain.ProductItem) (productItemID uint, err error)
	// product item image
	FindAllProductItemImages(ctx echo.Context, productItemID uint) (images []string, err error)
	SaveProductItemImage(ctx echo.Context, productItemID uint, image string) error
}
