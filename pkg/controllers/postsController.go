package controllers

import "github.com/gofiber/fiber/v2"

func PostsGetAll(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{})
}
