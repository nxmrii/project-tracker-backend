package project

import (
	"context"
	"testing"
)

type fakeTaskRepository struct {
	totalTasks     int
	completedTasks int
}

func (f *fakeTaskRepository) GetProjectTaskCounts(
	ctx context.Context,
	projectID int64,
) (int, int, error) {

	return f.totalTasks, f.completedTasks, nil
}

func TestGetProjectProgress(t *testing.T) {

	projectRepo := &fakeProjectRepository{}

	taskRepo := &fakeTaskRepository{
		totalTasks:     4,
		completedTasks: 2,
	}

	service := NewService(
		projectRepo,
		taskRepo,
	)

	result, err := service.GetProgress(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if result.TotalTasks != 4 {
		t.Errorf(
			"expected 4 total tasks, got %d",
			result.TotalTasks,
		)
	}

	if result.CompletedTasks != 2 {
		t.Errorf(
			"expected 2 completed tasks, got %d",
			result.CompletedTasks,
		)
	}

	if result.Progress != 50 {
		t.Errorf(
			"expected progress 50, got %f",
			result.Progress,
		)
	}
}

func TestGetProjectProgressNoTasks(t *testing.T) {

	projectRepo := &fakeProjectRepository{}

	taskRepo := &fakeTaskRepository{
		totalTasks:     0,
		completedTasks: 0,
	}

	service := NewService(
		projectRepo,
		taskRepo,
	)

	result, err := service.GetProgress(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if result.Progress != 0 {
		t.Errorf(
			"expected progress 0, got %f",
			result.Progress,
		)
	}
}
