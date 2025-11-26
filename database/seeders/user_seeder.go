package seeders

import (
	"go/basics/database/migrations"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []migrations.User{
		{Name: "Abil", Email: "abil@example.com"},
		{Name: "Jane Doe", Email: "jane@example.com"},
		{Name: "John Smith", Email: "john@example.com"},
	}

	for _, user := range users {
		// Use FirstOrCreate to avoid duplicates if rerun
		if err := db.FirstOrCreate(&migrations.User{}, user).Error; err != nil {
			return err
		}
	}
	return nil
}
