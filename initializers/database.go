package initializers

import (
	"learn/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn := os.Getenv("DB_DNS")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	utils.LogError(err, "Can not connect to DB")
}
