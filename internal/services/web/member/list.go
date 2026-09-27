package member

import (
	"context"
	"errors"

	"project-tracker-backend/internal/domain"
)

func (s *Service) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Member, error) {

	if projectID <= 0 {
		return nil, errors.New("invalid project id")
	}

	members, err := s.repo.ListByProjectID(
		ctx,
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return members, nil
}
