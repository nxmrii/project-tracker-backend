package task

import (
	"context"
	"errors"
	"strings"

	"project-tracker-backend/internal/domain"
	taskDTO "project-tracker-backend/internal/dtos/task"
)

func (s *Service) Create(
	ctx context.Context,
	projectID int64,
	req taskDTO.CreateRequest,
) (*domain.Task, error) {

	if projectID <= 0 {
		return nil, errors.New("invalid project id")
	}

	if req.MemberID <= 0 {
		return nil, errors.New("invalid member id")
	}

	title := strings.TrimSpace(req.Title)

	if title == "" {
		return nil, errors.New("task title is required")
	}

	status := strings.TrimSpace(req.Status)

	if status == "" {
		status = "todo"
	}

	if status != "todo" &&
		status != "in_progress" &&
		status != "done" {

		return nil, errors.New("invalid task status")
	}

	newTask := &domain.Task{
		ProjectID:   projectID,
		MemberID:    req.MemberID,
		Title:       title,
		Description: strings.TrimSpace(req.Description),
		Status:      status,
	}

	if err := s.repo.Create(ctx, newTask); err != nil {
		return nil, err
	}

	s.invalidateProjectProgressCache(
		ctx,
		projectID,
	)

	return newTask, nil
}
