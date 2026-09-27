package task

import (
	"context"
	"errors"
)

func (s *Service) Delete(
	ctx context.Context,
	id int64,
) error {

	if id <= 0 {
		return errors.New("invalid task id")
	}

	// نأتي بالـTask أولًا لأننا نحتاج ProjectID
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if task == nil {
		return errors.New("task not found")
	}

	// نحذفها من PostgreSQL
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	// نحذف Progress القديم من Redis
	s.invalidateProjectProgressCache(
		ctx,
		task.ProjectID,
	)

	return nil
}
