package project

import projectService "project-tracker-backend/internal/services/web/project"

type Handler struct {
	service *projectService.Service
}

func NewHandler(service *projectService.Service) *Handler {
	return &Handler{
		service: service,
	}
}
