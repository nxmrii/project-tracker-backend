package task

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
	taskDTO "project-tracker-backend/internal/dtos/task"
)

type fakeTaskRepository struct {
	createCalled       bool
	deleteCalled       bool
	updateStatusCalled bool
	updateCalled       bool
	tasks              []domain.Task
}

func (f *fakeTaskRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	f.deleteCalled = true

	for i, task := range f.tasks {
		if task.ID == id {
			f.tasks = append(
				f.tasks[:i],
				f.tasks[i+1:]...,
			)
			return nil
		}
	}

	return nil
}

func (f *fakeTaskRepository) Update(
	ctx context.Context,
	task *domain.Task,
) error {

	f.updateCalled = true

	for i := range f.tasks {
		if f.tasks[i].ID == task.ID {
			f.tasks[i] = *task
			return nil
		}
	}

	return nil
}

func (f *fakeTaskRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Task, error) {

	for i := range f.tasks {
		if f.tasks[i].ID == id {
			return &f.tasks[i], nil
		}
	}

	return nil, nil
}

func (f *fakeTaskRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
) error {

	f.updateStatusCalled = true

	for i := range f.tasks {
		if f.tasks[i].ID == id {
			f.tasks[i].Status = status
			return nil
		}
	}

	return nil
}

func (f *fakeTaskRepository) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Task, error) {

	var result []domain.Task

	for _, task := range f.tasks {
		if task.ProjectID == projectID {
			result = append(result, task)
		}
	}

	return result, nil
}

func (f *fakeTaskRepository) Create(
	ctx context.Context,
	task *domain.Task,
) error {

	f.createCalled = true
	task.ID = 1

	return nil
}

func TestCreateTaskInvalidMemberID(t *testing.T) {
	repo := &fakeTaskRepository{}
	service := NewService(repo)

	req := taskDTO.CreateRequest{
		Title:    "Create Login",
		MemberID: 0,
		DueDate:  "2026-10-10",
	}

	result, err := service.Create(
		context.Background(),
		1,
		req,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

func TestCreateTaskEmptyDueDate(t *testing.T) {
	repo := &fakeTaskRepository{}
	service := NewService(repo)

	req := taskDTO.CreateRequest{
		Title:    "Create Login",
		MemberID: 1,
		DueDate:  "",
	}

	result, err := service.Create(
		context.Background(),
		1,
		req,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}

func TestCreateTaskInvalidDueDate(t *testing.T) {
	repo := &fakeTaskRepository{}
	service := NewService(repo)

	req := taskDTO.CreateRequest{
		Title:    "Create Login",
		MemberID: 1,
		DueDate:  "wrong-date",
	}

	result, err := service.Create(
		context.Background(),
		1,
		req,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}

	if repo.createCalled {
		t.Fatal(
			"repository Create should not be called",
		)
	}
}
