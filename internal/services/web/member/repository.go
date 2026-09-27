package member

import (
	"context"

	"project-tracker-backend/internal/domain"
	memberDTO "project-tracker-backend/internal/dtos/member"
)

type MemberRepository interface {
	Create(
		ctx context.Context,
		member *domain.Member,
	) error

	ListByProjectID(
		ctx context.Context,
		projectID int64,
	) ([]domain.Member, error)

	GetByID(
		ctx context.Context,
		id int64,
	) (*domain.Member, error)

	Update(
		ctx context.Context,
		member *domain.Member,
	) error

	Delete(ctx context.Context, id int64) error

	GetContributionsByProjectID(
		ctx context.Context,
		projectID int64,
	) ([]memberDTO.ContributionResponse, error)
}
