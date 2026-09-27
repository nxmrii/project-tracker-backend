package task

import (
	"strconv"

	taskDTO "project-tracker-backend/internal/dtos/task"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpdateStatus(c *fiber.Ctx) error {

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

	var req taskDTO.UpdateStatusRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.service.UpdateStatus(
		c.UserContext(),
		id,
		req.Status,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task_status_updated",
	})
}
