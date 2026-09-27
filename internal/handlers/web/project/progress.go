package project

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetProgress(c *fiber.Ctx) error {

	projectID, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		return c.Status(
			fiber.StatusBadRequest,
		).JSON(fiber.Map{
			"error": "invalid project id",
		})
	}

	progress, err := h.service.GetProgress(
		c.UserContext(),
		projectID,
	)

	if err != nil {
		return c.Status(
			fiber.StatusBadRequest,
		).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(
		fiber.StatusOK,
	).JSON(progress)
}
