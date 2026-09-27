package web

import (
	memberHandler "project-tracker-backend/internal/handlers/web/member"

	"github.com/gofiber/fiber/v2"
)

func RegisterMemberRoutes(
	app *fiber.App,
	handler *memberHandler.Handler,
) {
	app.Post(
		"/projects/:projectId/members",
		handler.Create,
	)

	app.Get(
		"/projects/:projectId/members",
		handler.ListByProjectID,
	)

	app.Patch(
		"/members/:id",
		handler.Update,
	)

	app.Delete(
		"/members/:id",
		handler.Delete,
	)

	app.Get(
		"/projects/:projectId/members/contribution",
		handler.GetContributions,
	)
}
