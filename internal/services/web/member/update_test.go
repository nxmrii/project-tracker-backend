package member

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
	memberDTO "project-tracker-backend/internal/dtos/member"
)

func TestUpdateMemberSuccess(t *testing.T) {

	repo := &fakeMemberRepository{
		members: []domain.Member{
			{
				ID:        1,
				ProjectID: 1,
				Name:      "Noor",
				Email:     "noor@example.com",
				Role:      "Developer",
			},
		},
	}

	service := NewService(repo)

	req := memberDTO.UpdateRequest{
		Name:  "Noor Al Amri",
		Email: "noor@example.com",
		Role:  "Team Lead",
	}

	err := service.Update(
		context.Background(),
		1,
		req,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected repository Update to be called")
	}

	if repo.members[0].Role != "Team Lead" {
		t.Errorf(
			"expected Team Lead, got %s",
			repo.members[0].Role,
		)
	}
}
