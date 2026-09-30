package project

import (
	"context"
	"testing"
	"time"

	"project-tracker-backend/internal/domain"
)

func TestGetProjectByIDSuccess(t *testing.T) {
	repo := &fakeProjectRepository{
		projects: []domain.Project{
			{
				ID:       1,
				Name:     "Project Tracker",
				Deadline: time.Now().AddDate(0, 1, 0),
			},
		},
	}

	service := NewService(repo)

	result, err := service.GetByID(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected project, got nil")
	}

	if result.ID != 1 {
		t.Errorf(
			"expected id 1, got %d",
			result.ID,
		)
	}

	if result.Name != "Project Tracker" {
		t.Errorf(
			"expected Project Tracker, got %s",
			result.Name,
		)
	}
}

func TestGetProjectByIDInvalidID(t *testing.T) {
	repo := &fakeProjectRepository{}

	service := NewService(repo)

	result, err := service.GetByID(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}
}
