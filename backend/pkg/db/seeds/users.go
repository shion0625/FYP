package seeds

import (
	"fmt"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/utils"
	"gorm.io/gorm"
)

func CreateUserDomain(db *gorm.DB, cfg *config.Config, options ...func(*domain.User)) error {
	const (
		MinAge          = 20
		MaxAge          = 80
		ImageSize       = 100
		PasswordLength  = 10
		MinPincode      = 100000
		MaxPincode      = 999999
		ExpireAfterDays = 24 * 7
		YearFormat      = 100
	)

	gofakeit.Seed(0)

	// user domain
	user := domain.User{
		ID:          gofakeit.UUID(),
		Age:         uint(gofakeit.Number(MinAge, MaxAge)),
		GoogleImage: "109.jpeg",
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		UserName:    gofakeit.Username(),
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
		Password:    gofakeit.Password(true, true, true, true, true, PasswordLength),
		Verified:    gofakeit.Bool(),
		BlockStatus: gofakeit.Bool(),
		CreatedAt:   gofakeit.Date(),
		UpdatedAt:   gofakeit.Date(),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	address := domain.Address{
		Name:        gofakeit.Street(),
		PhoneNumber: gofakeit.Phone(),
		House:       gofakeit.Street(),
		Area:        gofakeit.City(),
		LandMark:    gofakeit.Street(),
		City:        gofakeit.City(),
		Pincode:     strconv.Itoa(gofakeit.Number(MinPincode, MaxPincode)),
		CountryName: gofakeit.Country(),
		CreatedAt:   gofakeit.Date(),
		UpdatedAt:   gofakeit.Date(),
	}
	if err := db.Create(&address).Error; err != nil {
		return err
	}

	userAddress := domain.UserAddress{
		UserID:    user.ID,
		User:      user,
		AddressID: address.ID,
		Address:   address,
		IsDefault: gofakeit.Bool(),
	}
	if err := db.Create(&userAddress).Error; err != nil {
		return err
	}

	creditNumber := gofakeit.CreditCardNumber(nil)
	// PaymentMethod domain
	paymentMethod := domain.PaymentMethod{
		Number: utils.Encrypt(creditNumber, user.ID+cfg.CreditCardKey),
		Expiry: fmt.Sprintf("%02d/%02d", gofakeit.Month(), gofakeit.Year()%YearFormat), Cvc: gofakeit.CreditCardCvv(),
		UserId:      user.ID,
		CardCompany: utils.GetCardIssuer(creditNumber),
		CreatedAt:   gofakeit.Date(),
		UpdatedAt:   gofakeit.Date(),
	}

	if err := db.Create(&paymentMethod).Error; err != nil {
		return err
	}
	// auth domain
	refreshSession := domain.RefreshSession{
		TokenID:      gofakeit.UUID(),
		UserID:       gofakeit.UUID(),
		RefreshToken: gofakeit.UUID(),
		ExpireAt:     time.Now().Add(time.Hour * ExpireAfterDays), // Expire after 7 days
		IsBlocked:    gofakeit.Bool(),
	}
	if err := db.Create(&refreshSession).Error; err != nil {
		return err
	}

	return nil
}

func CreateUsersDomain(db *gorm.DB, count int, cfg *config.Config) error {
	for i := 0; i < count; i++ {
		if err := CreateUserDomain(db, cfg); err != nil {
			return err
		}
	}

	return nil
}
