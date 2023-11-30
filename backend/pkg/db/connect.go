package db

import (
	"errors"
	"fmt"

	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s TimeZone=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword, cfg.DBSslmode, cfg.DBTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if !errors.Is(err, nil) {
		panic(err.Error())
	}

	err = db.AutoMigrate(

		// auth
		domain.RefreshSession{},

		// user
		domain.User{},
		domain.Country{},
		domain.Address{},
		domain.UserAddress{},
		domain.PaymentMethod{},

		// product
		domain.Product{},
		domain.Brand{},
		domain.Category{},
		domain.Variation{},
		domain.VariationOption{},
		domain.ProductItem{},
		domain.ProductConfiguration{},
		domain.ProductImage{},

		// offer
		domain.Offer{},
		domain.OfferProduct{},

		// order
		domain.ShopOrder{},
		domain.ShopOrderProductItem{},
		domain.ShopOrderVariation{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database models: %w", err)
	}

	return db, nil
}
