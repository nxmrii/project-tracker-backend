package project

import "context"

type TaskRepository interface {
	GetProjectTaskCounts(
		ctx context.Context,
		projectID int64,
	) (int, int, error)
}
