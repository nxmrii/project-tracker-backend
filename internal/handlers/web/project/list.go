package project

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) List(c *fiber.Ctx) error {

	projects, err := h.service.List(
		c.UserContext(),
	)

	if err != nil {

		log.Printf("LIST PROJECTS ERROR: %v", err)
		return c.Status(
			fiber.StatusInternalServerError,
		).JSON(fiber.Map{
			"error": "failed to get projects",
		})
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}
