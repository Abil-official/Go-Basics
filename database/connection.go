package database

import (
	"fmt"
	"go/basics/database/migrations"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes the connection and runs migrations
func ConnectDB() *gorm.DB {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read values from environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// ✅ Add client_encoding=UTF8 to DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8 TimeZone=Asia/Kathmandu",
		host, user, password, dbname, port,
	)
	if os.Getenv("APP_ENV") == "local" {
		dsn += " client_encoding=UTF8"
	}

	// Connect to Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations
	migrations.RunMigrations(db)

	return db
}
