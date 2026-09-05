package database

import (
	"belajargolang/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto migrate models
	err = DB.AutoMigrate(&models.Project{})
	if err != nil {
		return err
	}

	return nil
}

