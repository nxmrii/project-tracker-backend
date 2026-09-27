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
			email = $2,
			role = $3
		WHERE id = $4
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		member.Name,
		member.Email,
		member.Role,
		member.ID,
	)

	return err
}
