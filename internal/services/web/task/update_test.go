package task

import (
	"context"
	"testing"
	"time"

	"project-tracker-backend/internal/domain"
	taskDTO "project-tracker-backend/internal/dtos/task"
)

func TestUpdateTaskSuccess(t *testing.T) {

	oldDueDate, _ := time.Parse(
		"2006-01-02",
		"2026-10-01",
	)

	repo := &fakeTaskRepository{
		tasks: []domain.Task{
			{
				ID:        1,
				ProjectID: 1,
				MemberID:  1,
				Title:     "Old Task",
				Status:    "todo",
				DueDate:   oldDueDate,
			},
		},
	}

	service := NewService(repo)

	req := taskDTO.UpdateRequest{
		Title:    "New Task",
		MemberID: 1,
		DueDate:  "2026-10-15",
	}

	err := service.Update(
		context.Background(),
		1,
		req,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if !repo.updateCalled {
		t.Fatal(
			"expected repository Update to be called",
		)
	}

	if repo.tasks[0].Title != "New Task" {
		t.Errorf(
			"expected New Task, got %s",
			repo.tasks[0].Title,
		)
	}

	if repo.tasks[0].MemberID != 1 {
		t.Errorf(
			"expected member id 1, got %d",
			repo.tasks[0].MemberID,
		)
	}

	if repo.tasks[0].DueDate.Format(
		"2006-01-02",
	) != "2026-10-15" {
		t.Errorf(
			"expected due date 2026-10-15, got %s",
			repo.tasks[0].DueDate.Format(
				"2006-01-02",
			),
		)
	}
}
