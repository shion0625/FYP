package repository

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
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
		userDetails.UserName, userDetails.Email, userDetails.Phone, userDetails.ID).First(&user).Error
	return
}

func (c *userDatabase) SaveUser(ctx echo.Context, user domain.User) (userID uint, err error) {
	// save the user details
	query := `INSERT INTO users (user_name, first_name,
		last_name, age, email, phone, password, google_image, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 ) RETURNING id`

	createdAt := time.Now()
	err = c.DB.Raw(query, user.UserName, user.FirstName, user.LastName,
		user.Age, user.Email, user.Phone, user.Password, user.GoogleImage, createdAt).Scan(&userID).Error

	return userID, err
}
