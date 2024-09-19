package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fullstack-capstone-backend/models" // Import your models package
)

var DB *gorm.DB
var JwtSecret = []byte("222b224fd5c59f7f304a999de07283015a97122aecfa4693b2469bbcd139ba4b") // Ensure this is defined

func Initialize() {
	dsn := "host=localhost user=postgres password=my-password dbname=final-project-ruangguru port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
		&models.Cart{},
	)
	if err != nil {
		log.Fatalf("failed to auto migrate models: %v", err)
	}
}
