package task

import (
	"context"

	"project-tracker-backend/internal/domain"
)

type TaskRepository interface {
	Create(
		ctx context.Context,
		task *domain.Task,
	) error

	ListByProjectID(
		ctx context.Context,
		projectID int64,
	) ([]domain.Task, error)

	GetByID(
		ctx context.Context,
		id int64,
	) (*domain.Task, error)

	UpdateStatus(
		ctx context.Context,
		id int64,
		status string,
	) error

	Update(
		ctx context.Context,
		task *domain.Task,
	) error

	Delete(
		ctx context.Context,
		id int64,
	) error
}
