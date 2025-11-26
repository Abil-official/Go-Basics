package migrations

import (
	"log"

	"gorm.io/gorm"
)

// Example model
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}

// RunMigrations runs all database migrations
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("✅ Migration completed successfully!")
}
