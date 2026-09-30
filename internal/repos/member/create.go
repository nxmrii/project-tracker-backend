package member

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) Create(
	ctx context.Context,
	member *domain.Member,
) error {

	query := `
		INSERT INTO members (
			project_id,
			name,
			role
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		member.ProjectID,
		member.Name,
		member.Role,
	).Scan(
		&member.ID,
		&member.CreatedAt,
	)

	return err
}
