package member

import (
	"context"
	"testing"

	memberDTO "project-tracker-backend/internal/dtos/member"
)

func TestGetMemberContributions(t *testing.T) {

	repo := &fakeMemberRepository{
		contributions: []memberDTO.ContributionResponse{
			{
				MemberID:       1,
				MemberName:     "Noor",
				TotalTasks:     3,
				CompletedTasks: 2,
				Contribution:   66.67,
			},
			{
				MemberID:       2,
				MemberName:     "Sara",
				TotalTasks:     2,
				CompletedTasks: 1,
				Contribution:   50,
			},
		},
	}

	service := NewService(repo)

	result, err := service.GetContributions(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if len(result) != 2 {
		t.Fatalf(
			"expected 2 members, got %d",
			len(result),
		)
	}

	if result[0].Contribution != 66.67 {
		t.Errorf(
			"expected 66.67, got %f",
			result[0].Contribution,
		)
	}

	if result[1].Contribution != 50 {
		t.Errorf(
			"expected 50, got %f",
			result[1].Contribution,
		)
	}
}
