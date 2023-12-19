package utils_test

import (
	"testing"

	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	t.Parallel()

	// Define input and output
	type input struct {
		stringToEncrypt string
		keyString       string
	}

	type output struct {
		decryptedString string
	}

	tests := map[string]struct {
		input input
		want  output
	}{
		"Normal case: Encrypt and then Decrypt": {
			input: input{
				stringToEncrypt: "testString",
				keyString:       "testKey",
			},
			want: output{
				decryptedString: "testString",
			},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			encryptedString := utils.Encrypt(tt.input.stringToEncrypt, tt.input.keyString)
			decryptedString := utils.Decrypt(encryptedString, tt.input.keyString)
			assert.Equal(t, tt.want.decryptedString, decryptedString)
		})
	}
}

func TestGetCardIssuer(t *testing.T) {
	t.Parallel()

	// Define input and output
	type input struct {
		cardNumber string
	}

	type output struct {
		cardIssuer string
	}

	tests := map[string]struct {
		input input
		want  output
	}{
		"Normal case: Visa": {
			input: input{
				cardNumber: "4",
			},
			want: output{
				cardIssuer: "Visa",
			},
		},
		"Normal case: MasterCard": {
			input: input{
				cardNumber: "5",
			},
			want: output{
				cardIssuer: "MasterCard",
			},
		},
		"Normal case: American Express": {
			input: input{
				cardNumber: "34",
			},
			want: output{
				cardIssuer: "American Express",
			},
		},
		"Normal case: Discover": {
			input: input{
				cardNumber: "6",
			},
			want: output{
				cardIssuer: "Discover",
			},
		},
		"Normal case: Diners Club": {
			input: input{
				cardNumber: "36",
			},
			want: output{
				cardIssuer: "Diners Club",
			},
		},
		"Normal case: Unknown": {
			input: input{
				cardNumber: "0",
			},
			want: output{
				cardIssuer: "Unknown",
			},
		},
	}

	for testName, tt := range tests {
		tt := tt

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			cardIssuer := utils.GetCardIssuer(tt.input.cardNumber)
			assert.Equal(t, tt.want.cardIssuer, cardIssuer)
		})
	}
}
