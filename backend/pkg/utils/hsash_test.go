package utils_test

import (
	"testing"

	"github.com/shion0625/FYP/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateHashFromPassword(t *testing.T) {
	t.Parallel()

	password := "testpassword"

	hashedPassword, err := utils.GenerateHashFromPassword(password)
	assert.Nil(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestVerifyHashAndPassword(t *testing.T) {
	t.Parallel()

	password := "testpassword"
	hashedPassword, _ := utils.GenerateHashFromPassword(password)

	isValid := utils.VerifyHashAndPassword(hashedPassword, password)
	assert.True(t, isValid)
}
