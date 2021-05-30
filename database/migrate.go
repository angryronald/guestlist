package database

import (
	guestRepo "github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	"gorm.io/gorm"
)

func generateModels() []interface{} {
	result := []interface{}{}
	result = append(result, &guestRepo.Guest{})

	return result
}

// Migrate migrates the database up
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(generateModels()...)
}
