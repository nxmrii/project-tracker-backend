package project

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
	projectDTO "project-tracker-backend/internal/dtos/project"
)

func TestUpdateProjectSuccess(t *testing.T) {

	repo := &fakeProjectRepository{
		projects: []domain.Project{
			{
				ID:          1,
				Name:        "Project Tracker",
				Description: "Old description",
				OwnerID:     1,
			},
		},
	}

	service := NewService(repo)

	req := projectDTO.UpdateRequest{
		Name:        "Project Management System",
		Description: "Updated description",
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
				ID:      1,
				Name:    "Project Tracker",
				OwnerID: 1,
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
