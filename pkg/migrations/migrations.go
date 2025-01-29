package migrations

import (
	"github.com/abdealt/videomaker/models"
	"gorm.io/gorm"
)

// Migrate runs the database migrations.
func MigrateAll(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Video{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Platform{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Status{}); err != nil {
		return err
	}

	return nil
}
