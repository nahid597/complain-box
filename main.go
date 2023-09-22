package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nahidh597/complain-box/database"
	"github.com/nahidh597/complain-box/routes"
)

func main() {
	database.Connect()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Can not connect .env file , main file")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	routes.SetUp(app)
	routes.SetUpPost(app)
	app.Listen(":" + port)

}
