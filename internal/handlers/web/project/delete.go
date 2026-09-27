package project

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Delete(c *fiber.Ctx) error {

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

	err = h.service.Delete(
		c.UserContext(),
		id,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "project_deleted",
	})
}
