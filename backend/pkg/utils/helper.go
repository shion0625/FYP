package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	suffixLength     = 4
	numbersLength    = 10
	couponCodeLength = 30
	skuLength        = 10
	bcryptCost       = 10
)

func StringToUint(str string) (uint, error) {
	val, err := strconv.Atoi(str)

	return uint(val), err
}

// generate userName.
func GenerateRandomUserName(firstName string) string {
	suffix, err := rand.Int(rand.Reader, big.NewInt(int64(numbersLength)))
	if err != nil {
		panic(err)
	}

	userName := (firstName + suffix.String())

	return strings.ToLower(userName)
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

// random coupons.
func GenerateCouponCode(couponCodeLength int) string {
	// letter for coupons
	couponCode, err := rand.Int(rand.Reader, big.NewInt(int64(couponCodeLength)))
	if err != nil {
		panic(err)
	}
	// convert into string and return the random letter array
	return couponCode.String()
}

func StringToTime(timeString string) (timeValue time.Time, err error) {
	// parse the string time to time
	timeValue, err = time.Parse(time.RFC3339Nano, timeString)

	if err != nil {
		return timeValue, fmt.Errorf("faild to parse given time %v to time variable \nivalid input", timeString)
	}

	return timeValue, fmt.Errorf("failed to parse time: %w", err)
}

func GenerateRandomString(length int) (string, error) {
	sku := make([]byte, length)

	_, err := rand.Read(sku)
	if err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}

	return hex.EncodeToString(sku), nil
}

func RandomInt(min, max int) int {
	// 乱数の範囲は[min, max)となるように調整する
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}

	return int(num.Int64()) + min
}

func GetHashedPassword(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return hashedPassword, fmt.Errorf("failed to generate hashed password: %w", err)
	}

	hashedPassword = string(hash)

	return hashedPassword, nil
}

func ComparePasswordWithHashedPassword(actualpassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(actualpassword))

	return fmt.Errorf("failed to compare password with hashed password: %w", err)
}
