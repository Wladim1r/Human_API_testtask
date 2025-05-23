package db

import (
	"fmt"
	"os"

	"github.com/Wladim1r/testtask/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	cfg := &gorm.Config{}

	if os.Getenv("DEBUG") == "true" {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	} else {
		cfg.Logger = logger.Default.LogMode(logger.Silent)
	}

	var err error

	db, err = gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		return nil, fmt.Errorf("could not connect to DB %w", err)
	}

	if err := db.AutoMigrate(&models.Human{}); err != nil {
		return nil, fmt.Errorf("error when creation DB %w", err)
	}

	return db, nil
}
