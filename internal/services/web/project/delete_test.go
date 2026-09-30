package project

import (
	"context"
	"testing"
	"time"

	"project-tracker-backend/internal/domain"
)

// test for deleting a project successfully
func TestDeleteProjectSuccess(t *testing.T) {

	repo := &fakeProjectRepository{
		projects: []domain.Project{
			{
				ID:       1,
				Name:     "Project Tracker",
				Deadline: time.Now().AddDate(0, 1, 0), // 1 month from now
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

	if len(repo.projects) != 0 {
		t.Fatalf(
			"expected 0 projects, got %d",
			len(repo.projects),
		)
	}
}

// test for deleting a project with an invalid ID
func TestDeleteProjectInvalidID(t *testing.T) {

	repo := &fakeProjectRepository{}

	service := NewService(repo)

	err := service.Delete(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if repo.deleteCalled {
		t.Fatal("repository Delete should not be called")
	}
}
