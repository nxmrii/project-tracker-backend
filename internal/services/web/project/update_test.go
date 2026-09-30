package project

import (
	"context"
	"testing"
	"time"

	"project-tracker-backend/internal/domain"
	projectDTO "project-tracker-backend/internal/dtos/project"
)

func TestUpdateProjectSuccess(t *testing.T) {

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

	req := projectDTO.UpdateRequest{
		Name: "Project Management System",
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

	if repo.projects[0].Name != "Project Management System" {
		t.Errorf(
			"expected updated project name, got %s",
			repo.projects[0].Name,
		)
	}
}

func TestUpdateProjectEmptyName(t *testing.T) {

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

	req := projectDTO.UpdateRequest{
		Name: "",
	}

	err := service.Update(
		context.Background(),
		1,
		req,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if repo.updateCalled {
		t.Fatal("repository Update should not be called")
	}
}
