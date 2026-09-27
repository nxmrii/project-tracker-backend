package project

import (
	"context"
	"errors"
)

func (s *Service) Delete(
	ctx context.Context,
	id int64,
) error {

	if id <= 0 {
		return errors.New("invalid project id")
	}

	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if project == nil {
		return errors.New("project not found")
	}

	return s.repo.Delete(ctx, id)
}
