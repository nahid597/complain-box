package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nahidh597/complain-box/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when load .env")
	}

	dsn := os.Getenv("DSN")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can not connect to database")
	} else {
		log.Println("Connected with DB successfully..")
	}

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

}
