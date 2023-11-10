package utils_test

import (
	"testing"

	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// generate random string for testing.
func TestGenerateRandomString(t *testing.T) {
	t.Parallel()

	// Define input and output.
	type input struct {
		length int
	}

	type output struct {
		want    string
		wantErr error
	}

	tests := map[string]struct {
		input input
		want  output
	}{
		"Normal case: Generate random string with length 10": {
			input: input{length: 10},
			want:  output{wantErr: nil},
		},
		"Normal case: Generate random string with length 20": {
			input: input{length: 20},
			want:  output{wantErr: nil},
		},
		"Normal case: Generate random string with length 30": {
			input: input{length: 30},
			want:  output{wantErr: nil},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got, err := utils.GenerateRandomString(tt.input.length)
			assert.Equal(t, tt.want.wantErr, err)
			assert.Len(t, got, tt.input.length*2)
		})
	}
}
