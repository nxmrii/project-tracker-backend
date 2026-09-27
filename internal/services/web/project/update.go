package project

import (
	"context"
	"errors"
	"strings"

	projectDTO "project-tracker-backend/internal/dtos/project"
)

func (s *Service) Update(
	ctx context.Context,
	id int64,
	req projectDTO.UpdateRequest,
) error {

	if id <= 0 {
		return errors.New("invalid project id")
	}

	name := strings.TrimSpace(req.Name)

	if name == "" {
		return errors.New("project name is required")
	}

	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if project == nil {
		return errors.New("project not found")
	}

	project.Name = name
	project.Description = strings.TrimSpace(req.Description)

	return s.repo.Update(ctx, project)
}
