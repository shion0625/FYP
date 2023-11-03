package db

import (
	"fmt"
	"log"

	"github.com/shion0625/FYP/backend/pkg/config"
	// "github.com/shion0625/FYP/backend/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func to connect data base using config(database config) and return address of a new instnce of gorm DB
func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	// migrate the database tables
	err = db.AutoMigrate(

		//auth
		domain.RefreshSession{},
		domain.OtpSession{},
		//user
		domain.User{},
		domain.Country{},
		domain.Address{},
		domain.UserAddress{},

		//admin
		domain.Admin{},

		//product
		domain.Category{},
		domain.Product{},
		domain.Variation{},
		domain.VariationOption{},
		domain.ProductItem{},
		domain.ProductConfiguration{},
		domain.ProductImage{},

		// wish list
		domain.WishList{},

		// cart
		domain.Cart{},
		domain.CartItem{},

		// order
		domain.OrderStatus{},
		domain.ShopOrder{},
		domain.OrderLine{},
		domain.OrderReturn{},

		//offer
		domain.Offer{},
		domain.OfferCategory{},
		domain.OfferProduct{},

		// coupon
		domain.Coupon{},
		domain.CouponUses{},

		//wallet
		domain.Wallet{},
		domain.Transaction{},
	)

	if err != nil {
		log.Printf("failed to migrate database models")
		return nil, err
	}
	return db, err
}
