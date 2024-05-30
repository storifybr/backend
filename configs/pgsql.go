package configs

import (
	"fmt"
	"os"

	"github.com/storify/backend/internal/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Create database and connect to it
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to postgres: %v", err)
		return nil, err
	}

	// Migrate schema
	err = db.AutoMigrate(&schemas.Company{})
	if err != nil {
		logger.Errorf("failed to migrate schema: %v", err)
		return nil, err
	}

	return db, nil
}
