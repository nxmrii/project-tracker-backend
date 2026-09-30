package project

import (
	"context"
	"errors"
	"strings"
	"time"

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

	if strings.TrimSpace(req.Deadline) == "" {
		return nil, errors.New("deadline is required")
	}

	deadline, err := time.Parse(
		"2006-01-02",
		req.Deadline,
	)

	if err != nil {
		return nil, errors.New("invalid deadline format")
	}

	newProject := &domain.Project{
		Name:     name,
		Deadline: deadline,
	}

	err = s.repo.Create(ctx, newProject)
	if err != nil {
		return nil, err
	}

	return newProject, nil
}
