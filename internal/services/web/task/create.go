package task

import (
	"context"
	"errors"
	"strings"
	"time"

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

	if strings.TrimSpace(req.DueDate) == "" {
		return nil, errors.New("due date is required")
	}

	dueDate, err := time.Parse(
		"2006-01-02",
		req.DueDate,
	)

	if err != nil {
		return nil, errors.New(
			"invalid due date format",
		)
	}

	newTask := &domain.Task{
		ProjectID: projectID,
		MemberID:  req.MemberID,
		Title:     title,
		Status:    "todo",
		DueDate:   dueDate,
	}

	if err := s.repo.Create(
		ctx,
		newTask,
	); err != nil {
		return nil, err
	}

	s.invalidateProjectProgressCache(
		ctx,
		projectID,
	)

	return newTask, nil
}
