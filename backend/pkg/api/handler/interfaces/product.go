package interfaces

import "github.com/labstack/echo/v4"

type ProductHandler interface {
	GetAllCategories(ctx echo.Context) error
	SaveCategory(ctx echo.Context) error
	SaveSubCategory(ctx echo.Context) error
	SaveVariation(ctx echo.Context) error
	SaveVariationOption(ctx echo.Context) error
	GetAllVariations(ctx echo.Context) error

	GetAllProductsAdmin() func(ctx echo.Context) error
	GetAllProductsUser() func(ctx echo.Context) error

	SaveProduct(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error

	SaveProductItem(ctx echo.Context) error
	GetAllProductItemsAdmin() func(ctx echo.Context) error
	GetAllProductItemsUser() func(ctx echo.Context) error
}
