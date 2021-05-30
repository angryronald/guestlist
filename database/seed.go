package database

import (
	"gorm.io/gorm"
)

func seedUp() []interface{} {
	result := []interface{}{}
	return result
}

func Seed(db *gorm.DB) error {
	seeders := seedUp()
	for _, seed := range seeders {
		err := db.Create(seed).Error
		if err != nil {
			return err
		}
	}
	return nil
}
