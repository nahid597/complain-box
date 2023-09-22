package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nahidh597/complain-box/controller"
	"github.com/nahidh597/complain-box/middleware"
)

func SetUpPost(app *fiber.App) {
	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allposts", controller.AllPosts)
	app.Get("/api/post/:id", controller.DetailsPost)
	app.Put("/api/post/:id", controller.UpdatePost)
	app.Get("/api/userpost", controller.UniquePostsByUser)
	app.Delete("/api/post/:id", controller.DeletePost)
	app.Post("/api/upload-file", controller.UploadFile)
	app.Static("/api/uploads", "./uploads")
}
