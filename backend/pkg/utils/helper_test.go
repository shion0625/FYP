package utils_test

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var (
	suffixLength     = 4
	numbersLength    = 10
	couponCodeLength = 30
	skuLength        = 10
	bcryptCost       = 10
)

func TestGetUserIdFromContext(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		ctx  echo.Context
		want string
		err  error
	}{
		"Normal case: userId set in context": {
			ctx: func() echo.Context {
				e := echo.New()
				req := e.NewContext(nil, nil)
				req.Set("userId", "testUser")

				return req
			}(),
			want: "testUser",
			err:  nil,
		},
		"Abnormal case: userId not set in context": {
			ctx:  echo.New().NewContext(nil, nil),
			want: "",
			err:  fmt.Errorf("failed to get userID from context"),
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			// Mock rand.Read to return error
			patch := monkey.Patch(rand.Read, func(b []byte) (n int, err error) {
				return 0, fmt.Errorf("mock error")
			})
			defer patch.Unpatch()
			result, err := utils.GetUserIdFromContext(tt.ctx)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestParseStringToUint32(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  uint
		err   error
	}{
		"Normal case: numeric string": {
			input: "123",
			want:  123,
			err:   nil,
		},
		"Abnormal case: non-numeric string": {
			input: "abc",
			want:  0,
			err:   &strconv.NumError{Func: "ParseUint", Num: "abc", Err: strconv.ErrSyntax},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			result, err := utils.ParseStringToUint32(tt.input)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGenerateSKU(t *testing.T) {
	t.Parallel()

	mockError := fmt.Errorf("mock error")

	tests := map[string]struct {
		wantLen int
		err     error
	}{
		"Normal case": {
			wantLen: 20, // 10 bytes are encoded into 20-character string
			err:     nil,
		},
		"Abnormal case: error generating SKU": {
			wantLen: 0,
			err:     fmt.Errorf("failed to generate SKU: %w", mockError),
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			// Mock rand.Read to return error only for the abnormal case
			if tt.err != nil {
				patch := monkey.Patch(rand.Read, func(b []byte) (n int, err error) {
					return 0, mockError
				})
				defer patch.Unpatch()
			}

			result, err := utils.GenerateSKU()
			assert.Equal(t, tt.err, err)
			assert.Len(t, result, tt.wantLen)
		})
	}
}

func TestComparePasswordWithHashedPassword(t *testing.T) {
	t.Parallel()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcryptCost)

	tests := map[string]struct {
		actualPassword string
		hashedPassword string
		err            error
	}{
		"Normal case: passwords match": {
			actualPassword: "password",
			hashedPassword: string(hashedPassword),
			err:            nil,
		},
		"Abnormal case: passwords do not match": {
			actualPassword: "password1",
			hashedPassword: string(hashedPassword),

			err: fmt.Errorf("failed to compare password with hashed password: %w", bcrypt.ErrMismatchedHashAndPassword),
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := utils.ComparePasswordWithHashedPassword(tt.actualPassword, tt.hashedPassword)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGenerateUniqueString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		wantLen int
	}{
		"Normal case": {
			wantLen: 36, // UUID string length is 36
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			result := utils.GenerateUniqueString()
			assert.Len(t, result, tt.wantLen)
		})
	}
}
