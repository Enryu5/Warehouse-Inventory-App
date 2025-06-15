package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database connection used throughout the app
var DB *gorm.DB

// InitPostgres initializes the PostgreSQL connection without auto migration
func InitPostgres() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set up connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance from GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Connected to PostgreSQL successfully (migrations not run)")
	return DB
}

// Migrate performs database migrations for all models
func Migrate(db *gorm.DB) error {
	// Add all your models here
	err := db.AutoMigrate(
		&domain.Warehouse{},
		&domain.Item{},
		&domain.Admin{},
		&domain.Stock{},
		// Add other models as needed
	)
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	log.Println("Database tables migrated successfully")
	return nil
}
