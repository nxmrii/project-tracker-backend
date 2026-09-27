package task

import (
	"strconv"

	taskDTO "project-tracker-backend/internal/dtos/task"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Update(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id",
		})
	}

	var req taskDTO.UpdateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.service.Update(
		c.UserContext(),
		id,
		req,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task_updated",
	})
}
