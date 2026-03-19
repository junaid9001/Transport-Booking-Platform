package db

import (
	"log"

	"github.com/nabeel-mp/tripneo/train-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.DB_URL), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// err = db.AutoMigrate(...models...)

	log.Println("Connected to PostgreSQL (train-service)")
	DB = db
}
