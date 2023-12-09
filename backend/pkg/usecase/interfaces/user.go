package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type UserUseCase interface {
	// profile side
	FindProfile(ctx echo.Context, userId string) (domain.User, error)
	UpdateProfile(ctx echo.Context, user domain.User) error

	SaveAddress(ctx echo.Context, userID string, address domain.Address, isDefault bool) error // save address
	UpdateAddress(ctx echo.Context, addressBody request.EditAddress, userID string) error
	FindAddresses(ctx echo.Context, userID string) ([]response.Address, error) // to get all address of a user

	FindPaymentMethods(ctx echo.Context, userID string) ([]response.PaymentMethod, error)
	SavePaymentMethod(ctx echo.Context, userID string, paymentMethod request.PaymentMethod) error
}
