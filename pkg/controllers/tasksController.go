package controllers

import (
	"go-mvc/pkg/initializers"
	"go-mvc/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func FetchTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	initializers.DB.Order("created_at desc").Find(&tasks)

	return c.Render("tasks/index", fiber.Map{
		"Tasks": tasks,
	})
}

func FetchTask(c *fiber.Ctx) error {
	var task models.Task
	taskId := c.Params("id")

	result := initializers.DB.First(&task, taskId)

	if result.Error != nil {
		return c.Render("error", fiber.Map{
			"Error": result.Error,
		})
	}

	return c.Render("tasks/details", fiber.Map{
		"Task": task,
	})
}

func CreateTaskPage(c *fiber.Ctx) error {
	return c.Render("tasks/create", nil)
}

func CreateTask(c *fiber.Ctx) error {
	var body struct {
		Title string
		Body  string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	if body.Title == "" || body.Body == "" {
		return c.Render("error", fiber.Map{
			"Error": "Title and Body are required",
		})
	}

	task := models.Task{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&task)

	if result.Error != nil {
		return result.Error
	}

	return c.Redirect("/tasks")
}

func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("id")

	var task models.Task
	result := initializers.DB.First(&task, taskId)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	result = initializers.DB.Delete(&models.Task{}, taskId)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to delete task")
	}

	return c.SendStatus(204)
}
