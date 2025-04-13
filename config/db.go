package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Could not load .env from root (../../.env)")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("Failed to get DATABASE_URL from environment variables")
	}

	var err2 error
	DB, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		log.Fatalf("Failed to connect to database: %v", err2)
	}

	log.Println("Database connected successfully")
}
