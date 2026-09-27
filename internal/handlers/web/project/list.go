package project

import "github.com/gofiber/fiber/v2"

func (h *Handler) List(c *fiber.Ctx) error {

	projects, err := h.service.List(
		c.UserContext(),
	)

	if err != nil {
		return c.Status(
			fiber.StatusInternalServerError,
		).JSON(fiber.Map{
			"error": "failed to get projects",
		})
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}
