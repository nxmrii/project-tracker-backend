package member

import memberService "project-tracker-backend/internal/services/web/member"

type Handler struct {
	service *memberService.Service
}

func NewHandler(service *memberService.Service) *Handler {
	return &Handler{
		service: service,
	}
}
