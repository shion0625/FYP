package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type UserRepository interface {
	FindUserByUserID(ctx echo.Context, userID string) (user domain.User, err error)
	FindUserByEmail(ctx echo.Context, email string) (user domain.User, err error)
	FindUserByUserName(ctx echo.Context, userName string) (user domain.User, err error)
	FindUserByPhoneNumber(ctx echo.Context, phoneNumber string) (user domain.User, err error)
	FindUserByUserNameEmailOrPhone(ctx echo.Context, user domain.User) (domain.User, error)

	SaveUser(ctx echo.Context, user domain.User) (userID string, err error)
	// UpdateVerified(ctx echo.Context, userID string) error
	UpdateUser(ctx echo.Context, user domain.User) (err error)
	// UpdateBlockStatus(ctx echo.Context, userID string, blockStatus bool) error

	// //address
	// FindAddressByID(ctx echo.Context, addressID uint) (response.Address, error)
	IsAddressIDExist(ctx echo.Context, addressID uint) (exist bool, err error)
	IsAddressAlreadyExistForUser(ctx echo.Context, address domain.Address, userID string) (bool, error)
	FindAllAddressByUserID(ctx echo.Context, userID string) ([]response.Address, error)
	FindAddressByUserIDAndAddressID(ctx echo.Context, userID string, addressID uint) (domain.Address, error)
	SaveAddress(ctx echo.Context, address domain.Address) (addressID uint, err error)
	UpdateAddress(ctx echo.Context, address domain.Address) error
	// // address join table
	SaveUserAddress(ctx echo.Context, userAdress domain.UserAddress) error
	UpdateUserAddress(ctx echo.Context, userAddress domain.UserAddress) error

	FindAllPaymentMethodsByUserID(ctx echo.Context, userID string) ([]response.PaymentMethod, error)
	SavePaymentMethod(ctx echo.Context, paymentMethod domain.PaymentMethod) (uint, error)
	IsPaymentMethodIDExist(ctx echo.Context, paymentMethodID uint) (exist bool, err error)
	UpdatePaymentMethod(ctx echo.Context, paymentMethod domain.PaymentMethod) error

}
