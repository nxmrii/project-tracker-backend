package task

import taskService "project-tracker-backend/internal/services/web/task"

type Handler struct {
	service *taskService.Service
}

func NewHandler(service *taskService.Service) *Handler {
	return &Handler{
		service: service,
	}
}
