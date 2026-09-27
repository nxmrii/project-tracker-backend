package task

import (
	"context"
	"errors"

	"project-tracker-backend/internal/domain"
)

func (s *Service) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Task, error) {

	if projectID <= 0 {
		return nil, errors.New("invalid project id")
	}

	tasks, err := s.repo.ListByProjectID(
		ctx,
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
