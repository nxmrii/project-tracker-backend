package member

import (
	"context"
	"errors"
	"strings"

	memberDTO "project-tracker-backend/internal/dtos/member"
)

func (s *Service) Update(
	ctx context.Context,
	id int64,
	req memberDTO.UpdateRequest,
) error {

	if id <= 0 {
		return errors.New("invalid member id")
	}

	name := strings.TrimSpace(req.Name)
	email := strings.TrimSpace(req.Email)
	role := strings.TrimSpace(req.Role)

	if name == "" {
		return errors.New("member name is required")
	}

	if email == "" {
		return errors.New("member email is required")
	}

	if role == "" {
		return errors.New("member role is required")
	}

	member, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if member == nil {
		return errors.New("member not found")
	}

	member.Name = name
	member.Email = email
	member.Role = role

	return s.repo.Update(ctx, member)
}
