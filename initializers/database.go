package initializers

import (
	"gin-gorm-postg/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDB() {
	var err error

	dbUrl := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	if os.Getenv("AUTO_MIGRATE") == "Y" {
		if err := DB.AutoMigrate(models.Post{}); err != nil {
			log.Fatal("Failed to migrate database")
		}
	}

}