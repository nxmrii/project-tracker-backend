package member

import (
	"context"
	"errors"
	"strings"

	"project-tracker-backend/internal/domain"
	memberDTO "project-tracker-backend/internal/dtos/member"
)

func (s *Service) Create(
	ctx context.Context,
	projectID int64,
	req memberDTO.CreateRequest,
) (*domain.Member, error) {

	if projectID <= 0 {
		return nil, errors.New("invalid project id")
	}

	name := strings.TrimSpace(req.Name)
	role := strings.TrimSpace(req.Role)

	if name == "" {
		return nil, errors.New("member name is required")
	}

	if role == "" {
		return nil, errors.New("member role is required")
	}

	newMember := &domain.Member{
		ProjectID: projectID,
		Name:      name,
		Role:      role,
	}

	if err := s.repo.Create(ctx, newMember); err != nil {
		return nil, err
	}

	return newMember, nil
}
