package member

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
)

func TestListMembersByProjectIDSuccess(t *testing.T) {

	repo := &fakeMemberRepository{
		members: []domain.Member{
			{
				ID:        1,
				ProjectID: 1,
				Name:      "Noor",
				Role:      "Developer",
			},
			{
				ID:        2,
				ProjectID: 1,
				Name:      "Sara",
				Role:      "Designer",
			},
			{
				ID:        3,
				ProjectID: 2,
				Name:      "Ali",
				Role:      "Tester",
			},
		},
	}

	service := NewService(repo)

	members, err := service.ListByProjectID(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if len(members) != 2 {
		t.Fatalf(
			"expected 2 members, got %d",
			len(members),
		)
	}

	if members[0].Name != "Noor" {
		t.Errorf(
			"expected Noor, got %s",
			members[0].Name,
		)
	}

	if members[1].Name != "Sara" {
		t.Errorf(
			"expected Sara, got %s",
			members[1].Name,
		)
	}
}

func TestListMembersInvalidProjectID(t *testing.T) {

	repo := &fakeMemberRepository{}

	service := NewService(repo)

	members, err := service.ListByProjectID(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if members != nil {
		t.Fatal("expected members to be nil")
	}
}
