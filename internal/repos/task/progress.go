package task

import "context"

func (r *Repository) GetProjectTaskCounts(
	ctx context.Context,
	projectID int64,
) (int, int, error) {

	var totalTasks int
	var completedTasks int

	query := `
		SELECT
			COUNT(*) AS total_tasks,
			COUNT(*) FILTER (WHERE status = 'done') AS completed_tasks
		FROM tasks
		WHERE project_id = $1
	`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		projectID,
	).Scan(
		&totalTasks,
		&completedTasks,
	)

	if err != nil {
		return 0, 0, err
	}

	return totalTasks, completedTasks, nil
}
