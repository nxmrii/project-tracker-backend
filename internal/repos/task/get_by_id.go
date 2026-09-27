package task

import (
	"context"
	"database/sql"
	"errors"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Task, error) {

	var task domain.Task

	query := `
		SELECT
			id,
			project_id,
			member_id,
			title,
			description,
			status,
			created_at,
			updated_at
		FROM tasks
		WHERE id = $1
	`

	err := r.db.GetContext(
		ctx,
		&task,
		query,
		id,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}
