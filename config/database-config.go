package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to loan env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed Connection Database")
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close connection DB")
	}
	dbSQL.Close()
}
