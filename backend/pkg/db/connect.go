package db

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func ConnectDatabase() *SQLHandler {
	dsn := os.Getenv("DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if !errors.Is(err, nil) {
		panic(err.Error())
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}
