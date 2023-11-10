package usecase

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/domain"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

const (
	countryCode       = "+91"
	otpExpireDuration = time.Minute * 2
)

type authUseCase struct {
	userRepo repoInterfaces.UserRepository
}

func NewAuthUseCase(
	userRepo repoInterfaces.UserRepository,
) interfaces.AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
	}
}

func (c *authUseCase) UserLogin(ctx echo.Context, loginInfo request.Login) (string, error) {
	var (
		user domain.User
		err  error
	)

	switch {
	case loginInfo.Email != "":
		user, err = c.userRepo.FindUserByEmail(ctx, loginInfo.Email)
	case loginInfo.UserName != "":
		user, err = c.userRepo.FindUserByUserName(ctx, loginInfo.UserName)
	case loginInfo.Phone != "":
		user, err = c.userRepo.FindUserByPhoneNumber(ctx, loginInfo.Phone)
	default:
		return "", ErrEmptyLoginCredentials
	}

	if err != nil {
		return "", fmt.Errorf("failed to find user from database: %w", err)
	}

	if user.ID == "" {
		return "", ErrUserNotExist
	}

	if user.BlockStatus {
		return "", ErrUserBlocked
	}

	err = utils.ComparePasswordWithHashedPassword(loginInfo.Password, user.Password)
	if err != nil {
		return "", ErrWrongPassword
	}

	return user.ID, nil
}

func (c *authUseCase) UserSignUp(ctx echo.Context, signUpDetails domain.User) (string, error) {
	existUser, err := c.userRepo.FindUserByUserNameEmailOrPhoneNotID(ctx, signUpDetails)
	if err != nil {
		return "", fmt.Errorf("failed to check user details already exist: %w", err)
	}

	// if user credentials already exist and  verified then return it as errors
	if existUser.ID != "" && existUser.Verified {
		err = utils.CompareUserExistingDetails(existUser, signUpDetails)

		return "", fmt.Errorf("failed to check user details already exist: %w", err)
	}

	userID := existUser.ID

	if userID == "" { // if user not exist then save user on database
		hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
		if err != nil {
			return "", fmt.Errorf("failed to hash the password: %w", err)
		}

		signUpDetails.Password = hashPass
		_, err = c.userRepo.SaveUser(ctx, signUpDetails)

		if err != nil {
			return "", fmt.Errorf("failed to save user details: %w", err)
		}
	}

	return "success", nil
}
