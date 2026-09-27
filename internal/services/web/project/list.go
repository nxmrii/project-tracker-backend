package project

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (s *Service) List(
	ctx context.Context,
) ([]domain.Project, error) {

	projects, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
