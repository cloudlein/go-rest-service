package database

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func buildDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

func DatabaseInit() {
	var err error
	dsn := buildDSN()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("DSN:", dsn)

	if err != nil {
		log.Fatal("Failed to Connect to database...")
	}

	fmt.Println("Connecting to database...")
}
