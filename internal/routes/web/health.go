package web

import "github.com/gofiber/fiber/v2"

func RegisterHealthRoute(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
