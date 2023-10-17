package routes

import (
	"go-mvc/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func HandleAppRoutes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex)
}
