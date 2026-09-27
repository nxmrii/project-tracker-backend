package member

import "context"

func (r *Repository) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `
		DELETE FROM members
		WHERE id = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)

	return err
}
