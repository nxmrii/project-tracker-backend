package project

import (
	"context"
	"errors"

	"project-tracker-backend/internal/domain"
)

func (s *Service) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Project, error) {

	if id <= 0 {
		return nil, errors.New("invalid project id")
	}

	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}
