package task

import (
	"context"
)

func (r *Repository) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
) error {

	query := `
		UPDATE tasks
		SET
			status = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		status,
		id,
	)

	return err
}
