package member

import (
	"context"
	"errors"
)

func (s *Service) Delete(
	ctx context.Context,
	id int64,
) error {

	if id <= 0 {
		return errors.New("invalid member id")
	}

	member, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if member == nil {
		return errors.New("member not found")
	}

	return s.repo.Delete(ctx, id)
}
