package utils_test

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetParamID(t *testing.T) {
	t.Parallel()

	// Define input and output
	type input struct {
		requestURL string
		key        string
	}

	type output struct {
		wantID  *uint
		wantErr error
	}

	e := echo.New()
	rec := httptest.NewRecorder()
	parsedID, _ := utils.ParseStringToUint32("123")

	tests := map[string]struct {
		input input
		want  output
	}{
		"Normal case: Valid parameter": {
			input: input{
				requestURL: "/?test=123",
				key:        "test",
			},
			want: output{
				wantID:  &parsedID,
				wantErr: nil,
			},
		},
		"Abnormal case: Invalid parameter": {
			input: input{
				requestURL: "/",
				key:        "",
			},
			want: output{
				wantID:  nil,
				wantErr: nil,
			},
		},
		"Abnormal case: Unparsable parameter": {
			input: input{
				requestURL: "/?test=abc",
				key:        "test",
			},
			want: output{
				wantID:  nil,
				wantErr: utils.ErrInvalidParam,
			},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(echo.GET, tt.input.requestURL, nil)
			ctx := e.NewContext(req, rec)
			gotID, err := utils.GetParamID(ctx, tt.input.key)
			assert.Equal(t, tt.want.wantID, gotID)
			assert.Equal(t, tt.want.wantErr, err)
		})
	}
}
