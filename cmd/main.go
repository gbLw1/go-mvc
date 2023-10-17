package main

import (
	"go-mvc/pkg/controllers"
	"go-mvc/pkg/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	engine := html.New("./pkg/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", controllers.PostsGetAll)

	app.Listen(":" + os.Getenv("PORT"))
}
