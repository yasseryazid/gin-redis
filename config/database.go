package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if os.Getenv("DB_HOST") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("[!] Warning: .env file not found, using system environment variables")
		} else {
			log.Println("[V] .env file loaded successfully")
		}
	}

	requiredEnvs := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "DB_SSLMODE"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("[X] Missing required environment variable: %s", env)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[X] Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("[V] Connected to the database")
}
