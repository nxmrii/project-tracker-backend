package task

import (
	"context"
	"errors"
	"strings"

	taskDTO "project-tracker-backend/internal/dtos/task"
)

func (s *Service) Update(
	ctx context.Context,
	id int64,
	req taskDTO.UpdateRequest,
) error {

	if id <= 0 {
		return errors.New("invalid task id")
	}

	title := strings.TrimSpace(req.Title)

	if title == "" {
		return errors.New("task title is required")
	}

	if req.MemberID <= 0 {
		return errors.New("invalid member id")
	}

	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if task == nil {
		return errors.New("task not found")
	}

	task.Title = title
	task.Description = strings.TrimSpace(req.Description)
	task.MemberID = req.MemberID

	if err := s.repo.Update(ctx, task); err != nil {
		return err
	}

	s.invalidateProjectProgressCache(
		ctx,
		task.ProjectID,
	)

	return nil
}
