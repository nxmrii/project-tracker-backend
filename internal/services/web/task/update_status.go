package task

import (
	"context"
	"errors"
	"strings"
)

func (s *Service) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
) error {

	if id <= 0 {
		return errors.New("invalid task id")
	}

	status = strings.ToLower(
		strings.TrimSpace(status),
	)

	validStatuses := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	if !validStatuses[status] {
		return errors.New("invalid task status")
	}

	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if task == nil {
		return errors.New("task not found")
	}

	err = s.repo.UpdateStatus(
		ctx,
		id,
		status,
	)

	if err != nil {
		return err
	}

	s.invalidateProjectProgressCache(
		ctx,
		task.ProjectID,
	)

	return nil
}
