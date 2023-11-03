package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

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

func (c *userDatabase) FindUserByUserID(ctx context.Context, userID uint) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE id = $1`
	err = c.DB.Raw(query, userID).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserName(ctx context.Context, userName string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE user_name = $1`
	err = c.DB.Raw(query, userName).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE email = $1`
	err = c.DB.Raw(query, email).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE phone = $1`
	err = c.DB.Raw(query, phoneNumber).Scan(&user).Error

	return user, err
}
