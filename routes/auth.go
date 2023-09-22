package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nahidh597/complain-box/controller"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	// this check for api guard
	// app.Use(middleware.IsAuthenticated)
}
