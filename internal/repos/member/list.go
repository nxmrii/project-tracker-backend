package member

import (
	"context"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Member, error) {

	var members []domain.Member

	query := `
		SELECT
			id,
			project_id,
			name,
			role,
			created_at
		FROM members
		WHERE project_id = $1
		ORDER BY created_at DESC
	`

	err := r.db.SelectContext(
		ctx,
		&members,
		query,
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return members, nil
}
