package project

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) List(
	ctx context.Context,
) ([]domain.Project, error) {

	var projects []domain.Project //هي تأخذ نتيجة SQL وتضعها مباشرة داخل:

	query := `
		SELECT
			id,
			name,
			deadline,
			created_at,
			updated_at
		FROM projects
		ORDER BY created_at DESC
	`

	err := r.db.SelectContext(
		ctx,
		&projects,
		query,
	)

	if err != nil {
		return nil, err
	}

	return projects, nil
}
