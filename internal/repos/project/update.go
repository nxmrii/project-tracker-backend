package project

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) Update(
	ctx context.Context,
	project *domain.Project,
) error {

	query := `
		UPDATE projects
		SET
			name = $1,
			description = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING updated_at
	`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		project.Name,
		project.Description,
		project.ID,
	).Scan(
		&project.UpdatedAt,
	)

	return err
}
