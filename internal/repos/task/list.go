package task

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Task, error) {

	tasks := make([]domain.Task, 0)

	query := `
		SELECT
			id,
			project_id,
			member_id,
			title,
			status,
			due_date,
			created_at,
			updated_at
		FROM tasks
		WHERE project_id = $1
		ORDER BY created_at DESC
	`

	err := r.db.SelectContext(
		ctx,
		&tasks,
		query,
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
