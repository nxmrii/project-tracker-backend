package task

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
	taskDTO "project-tracker-backend/internal/dtos/task"
)

func TestUpdateTaskSuccess(t *testing.T) {

	repo := &fakeTaskRepository{
		tasks: []domain.Task{
			{
				ID:          1,
				ProjectID:   1,
				MemberID:    1,
				Title:       "Old Task",
				Description: "Old Description",
				Status:      "todo",
			},
		},
	}

	service := NewService(repo)

	req := taskDTO.UpdateRequest{
		Title:       "New Task",
		Description: "New Description",
		MemberID:    1,
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

	if repo.tasks[0].Title != "New Task" {
		t.Errorf(
			"expected New Task, got %s",
			repo.tasks[0].Title,
		)
	}
}
