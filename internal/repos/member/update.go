package member

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) Update(
	ctx context.Context,
	member *domain.Member,
) error {

	query := `
		UPDATE members
		SET
			name = $1,
			role = $2,
		WHERE id = $3
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		member.Name,
		member.Role,
		member.ID,
	)

	return err
}
