package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authDatabase struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authDatabase{
		DB: db,
	}
}

func (c *authDatabase) SaveRefreshSession(ctx echo.Context, refreshSession domain.RefreshSession) error {
	if err := c.DB.Table("refresh_sessions").Create(&refreshSession).Error; err != nil {
		return err
	}

	return nil
}

func (c *authDatabase) FindRefreshSessionByTokenID(ctx echo.Context, tokenID string) (refreshSession domain.RefreshSession, err error) {
	if err := c.DB.Table("refresh_sessions").Where("token_id = ?", tokenID).First(&refreshSession).Error; err != nil {
		return refreshSession, err
	}

	return refreshSession, nil
}
