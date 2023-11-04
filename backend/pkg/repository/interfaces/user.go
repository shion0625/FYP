package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type UserRepository interface {
	FindUserByUserID(ctx echo.Context, userID uint) (user domain.User, err error)
	FindUserByEmail(ctx echo.Context, email string) (user domain.User, err error)
	FindUserByUserName(ctx echo.Context, userName string) (user domain.User, err error)
	FindUserByPhoneNumber(ctx echo.Context, phoneNumber string) (user domain.User, err error)
	// FindUserByUserNameEmailOrPhoneNotID(ctx echo.Context, user domain.User) (domain.User, error)

	// SaveUser(ctx echo.Context, user domain.User) (userID uint, err error)
	// UpdateVerified(ctx echo.Context, userID uint) error
	// UpdateUser(ctx echo.Context, user domain.User) (err error)
	// UpdateBlockStatus(ctx echo.Context, userID uint, blockStatus bool) error

	// //address
	// FindCountryByID(ctx echo.Context, countryID uint) (domain.Country, error)
	// FindAddressByID(ctx echo.Context, addressID uint) (response.Address, error)
	// IsAddressIDExist(ctx echo.Context, addressID uint) (exist bool, err error)
	// IsAddressAlreadyExistForUser(ctx echo.Context, address domain.Address, userID uint) (bool, error)
	// FindAllAddressByUserID(ctx echo.Context, userID uint) ([]response.Address, error)
	// SaveAddress(ctx echo.Context, address domain.Address) (addressID uint, err error)
	// UpdateAddress(ctx echo.Context, address domain.Address) error
	// // address join table
	// SaveUserAddress(ctx echo.Context, userAdress domain.UserAddress) error
	// UpdateUserAddress(ctx echo.Context, userAddress domain.UserAddress) error

	// //wishlist
	// FindWishListItem(ctx echo.Context, productID, userID uint) (domain.WishList, error)
	// FindAllWishListItemsByUserID(ctx echo.Context, userID uint) ([]response.WishListItem, error)
	// SaveWishListItem(ctx echo.Context, wishList domain.WishList) error
	// RemoveWishListItem(ctx echo.Context, userID, productItemID uint) error
}
