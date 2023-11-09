package utils_test

import (
	"errors"
	"testing"

	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestCompareUserExistingDetails(t *testing.T) {
	t.Parallel()

	user1 := domain.User{
		Email:    "test1@example.com",
		UserName: "testuser1",
		Phone:    "1234567890",
	}
	user2 := domain.User{
		Email:    "test2@example.com",
		UserName: "testuser2",
		Phone:    "0987654321",
	}

	// 入出力を定義
	type input struct {
		user1 domain.User
		user2 domain.User
	}

	type output struct {
		wantErr error
	}

	req := input{
		user1: user1,
		user2: user2,
	}

	tests := map[string]struct {
		input input
		want  output
	}{
		"正常系: 既にユーザが存在しない": {
			input: req,
			want:  output{errors.New("failed to find existing details")},
		},
		"異常系: 同一のメールアドレスが存在": {
			input: input{
				user1: req.user1,
				user2: domain.User{
					Email:    req.user1.Email,
					UserName: req.user2.UserName,
					Phone:    req.user2.Phone,
				},
			},
			want: output{errors.New("user already exist with this email")},
		}, "異常系: 同一のユーザ名が存在": {
			input: input{
				user1: req.user1,
				user2: domain.User{
					Email:    req.user2.Email,
					UserName: req.user1.UserName,
					Phone:    req.user2.Phone,
				},
			},
			want: output{errors.New("user already exist with this user name")},
		}, "異常系: 同一の電話番号が存在": {
			input: input{
				user1: req.user1,
				user2: domain.User{
					Email:    req.user2.Email,
					UserName: req.user2.UserName,
					Phone:    req.user1.Phone,
				},
			},
			want: output{errors.New("user already exist with this phone")},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := utils.CompareUserExistingDetails(tt.input.user1, tt.input.user2)
			assert.Equal(t, tt.want.wantErr, err)
		})
	}
}
