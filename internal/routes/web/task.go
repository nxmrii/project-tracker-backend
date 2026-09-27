package web

import (
	taskHandler "project-tracker-backend/internal/handlers/web/task"

	"github.com/gofiber/fiber/v2"
)

func RegisterTaskRoutes(
	app *fiber.App,
	handler *taskHandler.Handler,
) {

	app.Post(
		"/projects/:projectId/tasks",
		handler.Create,
	)

	app.Get(
		"/projects/:projectId/tasks",
		handler.ListByProjectID,
	)

	app.Patch(
		"/tasks/:id/status",
		handler.UpdateStatus,
	)

	app.Patch(
		"/tasks/:id",
		handler.Update,
	)

	app.Delete(
		"/tasks/:id",
		handler.Delete,
	)
}
