package web

import (
	projectHandler "project-tracker-backend/internal/handlers/web/project"

	"github.com/gofiber/fiber/v2"
)

func RegisterProjectRoutes(
	app *fiber.App,
	handler *projectHandler.Handler,
) {
	projects := app.Group("/projects")

	projects.Post("/", handler.Create)
	projects.Get("/", handler.List)
	projects.Get("/:id", handler.GetByID)
	projects.Patch("/:id", handler.Update)
	projects.Delete("/:id", handler.Delete)
	projects.Get(
		"/:id/progress",
		handler.GetProgress,
	)
}
