package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	service "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

type userUserCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(userRepo interfaces.UserRepository) service.UserUseCase {
	return &userUserCase{
		userRepo: userRepo,
	}
}

func (c *userUserCase) FindProfile(ctx echo.Context, userID string) (domain.User, error) {
	user, err := c.userRepo.FindUserByUserID(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find user details: %w", err)
	}

	return user, nil
}

func (c *userUserCase) UpdateProfile(ctx echo.Context, user domain.User) error {
	// first check any other user exist with this entered unique fields
	checkUser, err := c.userRepo.FindUserByUserNameEmailOrPhone(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if checkUser.ID != "" { // if there is an user exist with given details then make it as error
		if err := utils.CompareUserExistingDetails(user, checkUser); err != nil {
			return fmt.Errorf("failed to compare user details: %w", err)
		}
	}

	// if user password given then hash the password
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)
		if err != nil {
			return fmt.Errorf("failed to generate hash password for user: %w", err)
		}

		user.Password = string(hash)
	}

	err = c.userRepo.UpdateUser(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// adddress.
func (c *userUserCase) SaveAddress(ctx echo.Context, userID string, address domain.Address, isDefault bool) error {
	exist, err := c.userRepo.IsAddressAlreadyExistForUser(ctx, address, userID)
	if err != nil {
		return fmt.Errorf("failed to check address already exist \nerror:%v", err.Error())
	}

	if exist {
		return fmt.Errorf("given address already exist for user")
	}

	// //this address not exist then create it
	// country, err := c.userRepo.FindCountryByID(ctx, address.CountryID)
	// if err != nil {
	// 	return err
	// } else if country.ID == 0 {
	// 	return errors.New("invalid country id")
	// }

	// save the address on database
	addressID, err := c.userRepo.SaveAddress(ctx, address)
	if err != nil {
		return fmt.Errorf("failed to save address: %w", err)
	}

	userAddress := domain.UserAddress{
		UserID:    userID,
		AddressID: addressID,
		IsDefault: isDefault,
	}

	// then update the address with user
	err = c.userRepo.SaveUserAddress(ctx, userAddress)

	if err != nil {
		return fmt.Errorf("failed to save user address: %w", err)
	}

	return nil
}

func (c *userUserCase) UpdateAddress(ctx echo.Context, addressBody request.EditAddress, userID string) error {
	if exist, err := c.userRepo.IsAddressIDExist(ctx, addressBody.ID); err != nil {
		return fmt.Errorf("failed to check address ID existence: %w", err)
	} else if !exist {
		return errors.New("invalid address id")
	}

	var address domain.Address
	if err := copier.Copy(&address, &addressBody); err != nil {
		return fmt.Errorf("failed to copy address: %w", err)
	}

	if err := c.userRepo.UpdateAddress(ctx, address); err != nil {
		return fmt.Errorf("failed to update address: %w", err)
	}

	// check the user address need to set default or not if it need then set it as default
	if addressBody.IsDefault != nil && *addressBody.IsDefault {
		userAddress := domain.UserAddress{
			UserID:    userID,
			AddressID: address.ID,
			IsDefault: *addressBody.IsDefault,
		}

		err := c.userRepo.UpdateUserAddress(ctx, userAddress)
		if err != nil {
			return fmt.Errorf("failed to update user address: %w", err)
		}
	}

	log.Printf("successfully address saved for user with user_id %s", userID)

	return nil
}

// get all address.
func (c *userUserCase) FindAddresses(ctx echo.Context, userID string) (addresses []response.Address, err error) {
	addresses, err = c.userRepo.FindAllAddressByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find addresses: %w", err)
	}

	return addresses, nil
}
