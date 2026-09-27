package project

import (
	"context"

	"project-tracker-backend/internal/domain"
)

// interface ProjectRepository defines the methods that a project repository should implement.
type ProjectRepository interface {
	Create(ctx context.Context, project *domain.Project) error
	List(ctx context.Context) ([]domain.Project, error)
	GetByID(ctx context.Context, id int64) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id int64) error
}
