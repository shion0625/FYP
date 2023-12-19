package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	suffixLength     = 4
	numbersLength    = 10
	couponCodeLength = 30
	skuLength        = 10
	bcryptCost       = 10
)

// take userId from context.
func GetUserIdFromContext(ctx echo.Context) (string, error) {
	userID, ok := ctx.Get("userId").(string)
	if !ok {
		return "", fmt.Errorf("failed to get userID from context")
	}

	return userID, nil
}

func ParseStringToUint32(param string) (uint, error) {
	value, err := strconv.ParseUint(param, 10, 32)

	return uint(value), err
}

// generate unique string for sku.
func GenerateSKU() (string, error) {
	sku := make([]byte, skuLength)

	_, err := rand.Read(sku)
	if err != nil {
		return "", fmt.Errorf("failed to generate SKU: %w", err)
	}

	return hex.EncodeToString(sku), nil
}

func ComparePasswordWithHashedPassword(actualPassword, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(actualPassword)); err != nil {
		return fmt.Errorf("failed to compare password with hashed password: %w", err)
	}

	return nil
}

func GenerateUniqueString() string {
	return uuid.NewString()
}
