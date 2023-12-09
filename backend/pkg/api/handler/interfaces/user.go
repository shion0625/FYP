package interfaces

import "github.com/labstack/echo/v4"

type UserHandler interface {
	GetProfile(ctx echo.Context) error
	UpdateProfile(ctx echo.Context) error

	SaveAddress(ctx echo.Context) error
	GetAllAddresses(ctx echo.Context) error
	UpdateAddress(ctx echo.Context) error

	SavePaymentMethod(ctx echo.Context) error
	GetAllPaymentMethods(ctx echo.Context) error
}
