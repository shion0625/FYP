package interfaces

import "github.com/labstack/echo/v4"

type UserHandler interface {
	GetProfile(ctx echo.Context)
	UpdateProfile(ctx echo.Context)

	SaveAddress(ctx echo.Context)
	GetAllAddresses(ctx echo.Context)
	UpdateAddress(ctx echo.Context)
	// SaveToWishList(ctx echo.Context)
	// RemoveFromWishList(ctx echo.Context)
	// GetWishList(ctx echo.Context)
}
