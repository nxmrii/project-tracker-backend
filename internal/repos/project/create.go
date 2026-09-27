package project

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) Create(
	ctx context.Context,
	project *domain.Project,
) error {

	query := `
		INSERT INTO projects (
			name,
			description,
			owner_id
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		project.Name,
		project.Description,
		project.OwnerID,
	).Scan(
		&project.ID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	return err
}
