package project

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetByID(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid project id",
		})
	}

	project, err := h.service.GetByID(
		c.UserContext(),
		id,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get project",
		})
	}

	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "project not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(project)
}
