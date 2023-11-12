package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/success", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.SendString(`{"message": "success"}`)
	})

	app.Get("/bad_request", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(http.StatusBadRequest).SendString(`{"error": "The new password cannot be the same as the previous one"}`)
	})

	app.Get("/not_found", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(http.StatusNotFound).SendString(`{"error":"This resource is not found at database 'vuln'"`)
	})

	app.Get("/internal_server_error", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(http.StatusInternalServerError).SendString(`{"error":"Error log with private information"`)
	})

	app.Listen(":8000")
}
