package utils_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// Test function for AppendMessageToError.
func TestAppendMessageToError(t *testing.T) {
	t.Parallel()

	// Define input and output
	type input struct {
		err     error
		message string
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input input
		want  output
	}{
		"Normal case: Error is nil": {
			input: input{
				err:     nil,
				message: "test message",
			},
			want: output{
				wantErr: errors.New("test message"),
			},
		},
		"Normal case: Error is not nil": {
			input: input{
				err:     errors.New("existing error"),
				message: "test message",
			},
			want: output{
				wantErr: fmt.Errorf("%w \n%s", errors.New("existing error"), "test message"),
			},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			err := utils.AppendMessageToError(tt.input.err, tt.input.message)
			assert.Equal(t, tt.want.wantErr, err)
		})
	}
}
