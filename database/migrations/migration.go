package migrations

import (
	"fmt"

	"github.com/cloudlein/go-rest-service/database"
	"github.com/cloudlein/go-rest-service/models"
	"github.com/gofiber/fiber/v2/log"
)

func Migrations() {
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("failed to migrate database")
	}

	fmt.Println("migrated successfully")
}
