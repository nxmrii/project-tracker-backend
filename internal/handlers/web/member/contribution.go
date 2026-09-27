package member

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetContributions(c *fiber.Ctx) error {

	projectID, err := strconv.ParseInt(
		c.Params("projectId"),
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

	contributions, err :=
		h.service.GetContributions(
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
	).JSON(contributions)
}
