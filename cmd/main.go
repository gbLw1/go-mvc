package main

import (
	"go-mvc/pkg/initializers"
	"go-mvc/pkg/middleware"
	"go-mvc/pkg/routes"
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
	// load templates
	engine := html.New("./pkg/views", ".html")

	// setup app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// configure app
	app.Static("/", "./public")
	app.Use(middleware.RequireAuth)

	// register routes
	routes.HandleAppRoutes(app)

	// start server
	app.Listen(":" + os.Getenv("PORT"))
}
