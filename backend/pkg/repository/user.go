package repository

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: db}
}

func (c *userDatabase) FindUserByUserID(ctx echo.Context, userID uint) (user domain.User, err error) {
	err = c.DB.Find(&user, userID).Error

	return user, err
}

func (c *userDatabase) FindUserByUserName(ctx echo.Context, userName string) (user domain.User, err error) {
	err = c.DB.Where("user_name = ?", userName).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByEmail(ctx echo.Context, email string) (user domain.User, err error) {
	err = c.DB.Where("email = ?", email).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByPhoneNumber(ctx echo.Context, phoneNumber string) (user domain.User, err error) {
	err = c.DB.Where("phone = ?", phoneNumber).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserNameEmailOrPhoneNotID(ctx echo.Context,
	userDetails domain.User,
) (user domain.User, err error) {
	err = c.DB.Where("(user_name = ? OR email = ? OR phone = ?) AND id != ?",
		userDetails.UserName, userDetails.Email, userDetails.Phone, userDetails.ID).Find(&user).Error

	return user, err
}

func (c *userDatabase) SaveUser(ctx echo.Context, user domain.User) (userID string, err error) {
	// save the user details
	user = domain.User{
		Age:         user.Age,
		GoogleImage: user.GoogleImage,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		UserName:    user.UserName,
		Email:       user.Email,
		Phone:       user.Phone,
		Password:    user.Password,
		CreatedAt:   time.Now(),
	}
	result := c.DB.Create(&user)
	fmt.Print(user.ID)
	fmt.Print(result.Error)

	return user.ID, result.Error
}
