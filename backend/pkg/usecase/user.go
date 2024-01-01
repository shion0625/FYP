package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	service "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost            = 10
	MinCreditNumberLength = 4
)

type userUserCase struct {
	userRepo      interfaces.UserRepository
	creditCardKey string
}

func NewUserUseCase(cfg *config.Config, userRepo interfaces.UserRepository) service.UserUseCase {
	return &userUserCase{
		userRepo:      userRepo,
		creditCardKey: cfg.CreditCardKey,
	}
}

func (u *userUserCase) FindProfile(ctx echo.Context, userID string) (domain.User, error) {
	user, err := u.userRepo.FindUserByUserID(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find user details: %w", err)
	}

	return user, nil
}

func (u *userUserCase) UpdateProfile(ctx echo.Context, user domain.User) error {
	// first check any other user exist with this entered unique fields
	checkUser, err := u.userRepo.FindUserByUserNameEmailOrPhone(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if checkUser.ID != "" { // if there is an user exist with given details then make it as error
		if err := utils.CompareUserExistingDetails(user, checkUser); err == nil {
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

	err = u.userRepo.UpdateUser(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// adddress.
func (u *userUserCase) SaveAddress(ctx echo.Context, userID string, address domain.Address, isDefault bool) error {
	exist, err := u.userRepo.IsAddressAlreadyExistForUser(ctx, address, userID)
	if err != nil {
		return fmt.Errorf("failed to check address already exist \nerror:%v", err.Error())
	}

	if exist {
		return fmt.Errorf("given address already exist for user")
	}

	// save the address on database
	addressID, err := u.userRepo.SaveAddress(ctx, address)
	if err != nil {
		return fmt.Errorf("failed to save address: %w", err)
	}

	userAddress := domain.UserAddress{
		UserID:    userID,
		AddressID: addressID,
		IsDefault: isDefault,
	}

	// then update the address with user
	err = u.userRepo.SaveUserAddress(ctx, userAddress)

	if err != nil {
		return fmt.Errorf("failed to save user address: %w", err)
	}

	return nil
}

func (u *userUserCase) UpdateAddress(ctx echo.Context, addressBody request.EditAddress, userID string) error {
	if exist, err := u.userRepo.IsAddressIDExist(ctx, addressBody.ID); err != nil {
		return fmt.Errorf("failed to check address ID existence: %w", err)
	} else if !exist {
		return errors.New("invalid address id")
	}

	var address domain.Address
	if err := copier.Copy(&address, &addressBody); err != nil {
		return fmt.Errorf("failed to copy address: %w", err)
	}

	if err := u.userRepo.UpdateAddress(ctx, address); err != nil {
		return fmt.Errorf("failed to update address: %w", err)
	}

	// check the user address need to set default or not if it need then set it as default
	if addressBody.IsDefault != nil && *addressBody.IsDefault {
		userAddress := domain.UserAddress{
			UserID:    userID,
			AddressID: address.ID,
			IsDefault: *addressBody.IsDefault,
		}

		err := u.userRepo.UpdateUserAddress(ctx, userAddress)
		if err != nil {
			return fmt.Errorf("failed to update user address: %w", err)
		}
	}

	log.Printf("successfully address saved for user with user_id %s", userID)

	return nil
}

// get all address.
func (u *userUserCase) FindAddresses(ctx echo.Context, userID string) (addresses []response.Address, err error) {
	addresses, err = u.userRepo.FindAllAddressByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find addresses: %w", err)
	}

	return addresses, nil
}

func (u *userUserCase) FindAddress(ctx echo.Context, userID string, addressID uint) (domain.Address, error) {
	user, err := u.userRepo.FindAddressByUserIDAndAddressID(ctx, userID, addressID)
	if err != nil {
		return domain.Address{}, fmt.Errorf("failed to find user details: %w", err)
	}

	return user, nil
}

func (u *userUserCase) FindPaymentMethods(ctx echo.Context, userID string) (paymentMethods []response.PaymentMethod, err error) {
	paymentMethods, err = u.userRepo.FindAllPaymentMethodsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find payment method: %w", err)
	}

	for i, method := range paymentMethods {
		creditNumberDecrypted := utils.Decrypt(method.Number, userID+u.creditCardKey)
		if len(creditNumberDecrypted) >= MinCreditNumberLength {
			paymentMethods[i].Number = creditNumberDecrypted[len(creditNumberDecrypted)-4:]
		}
	}

	return paymentMethods, nil
}

// to add new product.
func (u *userUserCase) SavePaymentMethod(ctx echo.Context, userID string, paymentMethod request.PaymentMethod) error {
	_, err := u.userRepo.SavePaymentMethod(ctx, domain.PaymentMethod{
		Number:      utils.Encrypt(paymentMethod.Number, userID+u.creditCardKey),
		Expiry:      paymentMethod.Expiry,
		Cvc:         paymentMethod.Cvc,
		UserId:      userID,
		CardCompany: utils.GetCardIssuer(paymentMethod.Number),
	})
	if err != nil {
		return fmt.Errorf("failed to save product: %w", err)
	}

	return nil
}

func (u *userUserCase) UpdatePaymentMethod(ctx echo.Context, userID string, paymentMethodBody request.UpdatePaymentMethod) error {
	if exist, err := u.userRepo.IsPaymentMethodIDExist(ctx, paymentMethodBody.ID); err != nil {
		return fmt.Errorf("failed to check payment method ID existence: %w", err)
	} else if !exist {
		return errors.New("invalid payment method id")
	}

	paymentMethod := domain.PaymentMethod{
		ID:          paymentMethodBody.ID,
		Number:      utils.Encrypt(paymentMethodBody.Number, userID+u.creditCardKey),
		Expiry:      paymentMethodBody.Expiry,
		Cvc:         paymentMethodBody.Cvc,
		CardCompany: utils.GetCardIssuer(paymentMethodBody.Number),
	}

	if err := u.userRepo.UpdatePaymentMethod(ctx, paymentMethod); err != nil {
		return fmt.Errorf("failed to update address: %w", err)
	}

	log.Printf("successfully payment method updated for user with user_id %s", userID)

	return nil
}
