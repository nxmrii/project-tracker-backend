package member

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
)

func TestDeleteMemberSuccess(t *testing.T) {

	repo := &fakeMemberRepository{
		members: []domain.Member{
			{
				ID:        1,
				ProjectID: 1,
				Name:      "Test Member",
			},
		},
	}

	service := NewService(repo)

	err := service.Delete(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.deleteCalled {
		t.Fatal("expected repository Delete to be called")
	}

	if len(repo.members) != 0 {
		t.Fatalf(
			"expected 0 members, got %d",
			len(repo.members),
		)
	}
}
