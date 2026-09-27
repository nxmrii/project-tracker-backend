package task

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
)

func TestUpdateTaskStatusSuccess(t *testing.T) {

	repo := &fakeTaskRepository{
		tasks: []domain.Task{
			{
				ID:        1,
				ProjectID: 1,
				MemberID:  1,
				Title:     "Create Login",
				Status:    "todo",
			},
		},
	}

	service := NewService(repo)

	err := service.UpdateStatus(
		context.Background(),
		1,
		"done",
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if !repo.updateStatusCalled {
		t.Fatal(
			"expected repository UpdateStatus to be called",
		)
	}

	if repo.tasks[0].Status != "done" {
		t.Errorf(
			"expected done, got %s",
			repo.tasks[0].Status,
		)
	}
}

func TestUpdateTaskStatusInvalidStatus(t *testing.T) {

	repo := &fakeTaskRepository{
		tasks: []domain.Task{
			{
				ID:     1,
				Status: "todo",
			},
		},
	}

	service := NewService(repo)

	err := service.UpdateStatus(
		context.Background(),
		1,
		"finished",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if repo.updateStatusCalled {
		t.Fatal(
			"repository UpdateStatus should not be called",
		)
	}
}
