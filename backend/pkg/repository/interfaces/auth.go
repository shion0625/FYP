package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
)

type AuthRepository interface {
	SaveRefreshSession(ctx echo.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx echo.Context, tokenID string) (domain.RefreshSession, error)
}
