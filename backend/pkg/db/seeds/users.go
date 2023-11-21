package seeds

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"gorm.io/gorm"
)

func CreateUserDomain(db *gorm.DB, options ...func(*domain.User)) error {
	const (
		MinAge          = 20
		MaxAge          = 80
		ImageSize       = 100
		PasswordLength  = 10
		MinPincode      = 100000
		MaxPincode      = 999999
		ExpireAfterDays = 24 * 7
	)

	gofakeit.Seed(0)

	// user domain
	user := domain.User{
		ID:          gofakeit.UUID(),
		Age:         uint(gofakeit.Number(MinAge, MaxAge)),
		GoogleImage: gofakeit.ImageURL(ImageSize, ImageSize),
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

	country := domain.Country{
		CountryName: gofakeit.Country(),
	}
	if err := db.Create(&country).Error; err != nil {
		return err
	}

	address := domain.Address{
		Name:        gofakeit.Street(),
		PhoneNumber: gofakeit.Phone(),
		House:       gofakeit.Street(),
		Area:        gofakeit.City(),
		LandMark:    gofakeit.Street(),
		City:        gofakeit.City(),
		Pincode:     uint(gofakeit.Number(MinPincode, MaxPincode)),
		CountryID:   country.ID,
		Country:     country,
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

func CreateUsersDomain(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		if err := CreateUserDomain(db); err != nil {
			return err
		}
	}

	return nil
}
