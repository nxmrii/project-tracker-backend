package member

import (
	"context"
	"errors"

	memberDTO "project-tracker-backend/internal/dtos/member"
)

func (s *Service) GetContributions(
	ctx context.Context,
	projectID int64,
) ([]memberDTO.ContributionResponse, error) {

	if projectID <= 0 {
		return nil, errors.New("invalid project id")
	}

	contributions, err :=
		s.repo.GetContributionsByProjectID(
			ctx,
			projectID,
		)

	if err != nil {
		return nil, err
	}

	return contributions, nil
}
