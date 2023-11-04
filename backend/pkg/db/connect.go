package db

import (
	"errors"
	"fmt"
	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s TimeZone=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword, cfg.DBSslmode, cfg.DBTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if !errors.Is(err, nil) {
		panic(err.Error())
	}

	err = autoMigrate(db)

	if err != nil {
		log.Printf("failed to migrate database models")
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	// migrate the database tables
	err := db.AutoMigrate(
		//user
		domain.User{},
		domain.Country{},
		domain.Address{},
		domain.UserAddress{},
	)
	return err
}
