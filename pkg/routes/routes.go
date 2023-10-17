package routes

import (
	"go-mvc/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func HandleAppRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)

	app.Get("/tasks", controllers.FetchTasks)
	app.Get("/tasks/:id<int>", controllers.FetchTask)
	app.Get("/tasks/create", controllers.CreateTaskPage)
	app.Post("/tasks", controllers.CreateTask)
	app.Delete("/tasks/:id<int>", controllers.DeleteTask)
}
