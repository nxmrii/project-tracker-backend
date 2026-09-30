package task

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
)

func TestListTasksByProjectIDSuccess(t *testing.T) {

	repo := &fakeTaskRepository{
		tasks: []domain.Task{
			{
				ID:        1,
				ProjectID: 1,
				MemberID:  1,
				Title:     "Create Login",
				Status:    "todo",
			},
			{
				ID:        2,
				ProjectID: 1,
				MemberID:  1,
				Title:     "Create Dashboard",
				Status:    "in-progress",
			},
			{
				ID:        3,
				ProjectID: 2,
				MemberID:  2,
				Title:     "Testing",
				Status:    "done",
			},
		},
	}

	service := NewService(repo)

	tasks, err := service.ListByProjectID(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf(
			"expected 2 tasks, got %d",
			len(tasks),
		)
	}

	if tasks[0].Title != "Create Login" {
		t.Errorf(
			"expected Create Login, got %s",
			tasks[0].Title,
		)
	}

	if tasks[1].Title != "Create Dashboard" {
		t.Errorf(
			"expected Create Dashboard, got %s",
			tasks[1].Title,
		)
	}
}

func TestListTasksInvalidProjectID(t *testing.T) {

	repo := &fakeTaskRepository{}

	service := NewService(repo)

	tasks, err := service.ListByProjectID(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if tasks != nil {
		t.Fatal("expected tasks to be nil")
	}
}
