package task

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Task, error) {

	var tasks []domain.Task

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
