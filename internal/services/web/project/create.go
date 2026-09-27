package project

import (
	"context"
	"errors"
	"strings"

	"project-tracker-backend/internal/domain"
	projectDTO "project-tracker-backend/internal/dtos/project"
)

func (s *Service) Create(
	ctx context.Context,
	req projectDTO.CreateRequest,
) (*domain.Project, error) {

	name := strings.TrimSpace(req.Name)

	if name == "" {
		return nil, errors.New("project name is required")
	}

	if req.OwnerID <= 0 {
		return nil, errors.New("owner id is required")
	}

	newProject := &domain.Project{
		Name:        name,
		Description: strings.TrimSpace(req.Description),
		OwnerID:     req.OwnerID,
	}

	err := s.repo.Create(ctx, newProject)
	if err != nil {
		return nil, err
	}

	return newProject, nil
}
