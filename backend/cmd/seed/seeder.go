package main

import (
	"log"

	"github.com/shion0625/FYP/backend/pkg/config"
	"github.com/shion0625/FYP/backend/pkg/db"
	"github.com/shion0625/FYP/backend/pkg/db/seeds"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error to load the config: ", err)
	}

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("\x1b[31mFailed to begin transaction: %s \x1b[0m", err)
	}

	tx := gormDB.Begin()
	if tx.Error != nil {
		log.Fatalf("\x1b[31mFailed to begin transaction: %s \x1b[0m", tx.Error)
	}

	for _, seed := range seeds.All(tx, cfg) {
		if err := seed.Run(tx); err != nil {
			tx.Rollback()
			log.Fatalf("\x1b[31mFailed '%s', failed with error: %s \x1b[0m", seed.Name, err)
		}

		log.Printf("\x1b[32mSuccess '%s'\x1b[0m", seed.Name)
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("\x1b[31mFailed to commit transaction: %s \x1b[0m", err)
	}

	log.Println("All seeds executed successfully.")
}
