package repository_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAuthRepository(t *testing.T) {
	t.Parallel()

	t.Run("SaveRefreshSession: Normal case", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
		}

		authRepo := repository.NewAuthRepository(gormDB)
		refreshSession := domain.RefreshSession{
			TokenID:      "token_id",
			UserID:       "user_id",
			RefreshToken: "refresh_token",
			ExpireAt:     time.Now(),
			IsBlocked:    false,
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"refresh_sessions\" (\"token_id\",\"user_id\",\"refresh_token\",\"expire_at\",\"is_blocked\") VALUES ($1,$2,$3,$4,$5)")).
			WithArgs(refreshSession.TokenID, refreshSession.UserID, refreshSession.RefreshToken, refreshSession.ExpireAt, refreshSession.IsBlocked).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = authRepo.SaveRefreshSession(echo.New().AcquireContext(), refreshSession)
		require.NoError(t, err)
	})

	t.Run("FindRefreshSessionByTokenID: Normal case", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
		}

		authRepo := repository.NewAuthRepository(gormDB)
		tokenID := "testTokenID"

		rows := sqlmock.NewRows([]string{"token_id"}).AddRow(tokenID)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"refresh_sessions\" WHERE token_id = $1 ORDER BY \"refresh_sessions\".\"token_id\" LIMIT 1")).WithArgs(tokenID).WillReturnRows(rows)

		_, err = authRepo.FindRefreshSessionByTokenID(echo.New().AcquireContext(), tokenID)
		require.NoError(t, err)
	})

	t.Run("FindRefreshSessionByTokenID: Error case", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
		}

		authRepo := repository.NewAuthRepository(gormDB)
		tokenID := "testTokenID"

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"refresh_sessions\" WHERE token_id = $1 ORDER BY \"refresh_sessions\".\"token_id\" LIMIT 1")).WithArgs(tokenID).WillReturnError(errors.New("error"))

		_, err = authRepo.FindRefreshSessionByTokenID(echo.New().AcquireContext(), tokenID)
		require.Error(t, err)
	})
}
