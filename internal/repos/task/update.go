package task

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) Update(
	ctx context.Context,
	task *domain.Task,
) error {

	query := `
		UPDATE tasks
		SET
			title = $1,
			member_id = $2,
			due_date = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		task.Title,
		task.MemberID,
		task.DueDate,
		task.ID,
	)

	return err
}
