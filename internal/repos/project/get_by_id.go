package project

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Project, error) {

	var project domain.Project

	query := `
		SELECT
			id,
			name,
			description,
			owner_id,
			created_at,
			updated_at
		FROM projects
		WHERE id = $1
	`

	err := r.db.GetContext(
		ctx,
		&project,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &project, nil
}
