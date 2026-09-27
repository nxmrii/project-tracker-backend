package member

import (
	"context"

	memberDTO "project-tracker-backend/internal/dtos/member"
)

func (r *Repository) GetContributionsByProjectID(
	ctx context.Context,
	projectID int64,
) ([]memberDTO.ContributionResponse, error) {

	var contributions []memberDTO.ContributionResponse

	query := `
		SELECT
			m.id AS member_id,
			m.name AS member_name,

			COUNT(t.id)::int AS total_tasks,

			COUNT(t.id) FILTER (
				WHERE t.status = 'done'
			)::int AS completed_tasks,

			CASE
				WHEN COUNT(t.id) = 0 THEN 0
				ELSE ROUND(
					(
						COUNT(t.id) FILTER (
							WHERE t.status = 'done'
						)::numeric
						/
						COUNT(t.id)::numeric
					) * 100,
					2
				)
			END AS contribution

		FROM members m

		LEFT JOIN tasks t
			ON t.member_id = m.id
			AND t.project_id = m.project_id

		WHERE m.project_id = $1

		GROUP BY
			m.id,
			m.name

		ORDER BY m.id
	`

	err := r.db.SelectContext(
		ctx,
		&contributions,
		query,
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return contributions, nil
}
