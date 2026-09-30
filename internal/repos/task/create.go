package task

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) Create(
	ctx context.Context,
	task *domain.Task,
) error {

	query := `
		INSERT INTO tasks (
			project_id,
			member_id,
			title,
			status,
			due_date
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			created_at,
			updated_at
	`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		task.ProjectID,
		task.MemberID,
		task.Title,
		task.Status,
		task.DueDate,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return err
}
