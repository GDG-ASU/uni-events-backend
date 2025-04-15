package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("Could not load .env")
    }

    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatalf("Failed to get DATABASE_URL from environment variables")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Database connected successfully")
    return db
}