package member

import (
	"context"
	"database/sql"
	"errors"

	"project-tracker-backend/internal/domain"
)

func (r *Repository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Member, error) {

	var member domain.Member

	query := `
		SELECT
			id,
			project_id,
			name,
			email,
			role,
			created_at
		FROM members
		WHERE id = $1
	`

	err := r.db.GetContext(
		ctx,
		&member,
		query,
		id,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &member, nil
}
