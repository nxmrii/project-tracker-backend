package project

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
	projectDTO "project-tracker-backend/internal/dtos/project"
)

type fakeProjectRepository struct {
	createCalled bool
	updateCalled bool
	deleteCalled bool
	projects     []domain.Project
}

func (f *fakeProjectRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	f.deleteCalled = true

	for i, project := range f.projects {
		if project.ID == id {
			f.projects = append(
				f.projects[:i],
				f.projects[i+1:]...,
			)

			return nil
		}
	}

	return nil
}

func (f *fakeProjectRepository) Update(
	ctx context.Context,
	project *domain.Project,
) error {

	f.updateCalled = true

	for i := range f.projects {
		if f.projects[i].ID == project.ID {
			f.projects[i] = *project
			return nil
		}
	}

	return nil
}

func (f *fakeProjectRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Project, error) {

	for _, project := range f.projects {
		if project.ID == id {
			return &project, nil
		}
	}

	return nil, nil
}

func (f *fakeProjectRepository) List(
	ctx context.Context,
) ([]domain.Project, error) {
	return f.projects, nil
}

func (f *fakeProjectRepository) Create(
	ctx context.Context,
	project *domain.Project,
) error {
	f.createCalled = true

	project.ID = 1

	return nil
}

func TestCreateProjectEmptyName(t *testing.T) {
	repo := &fakeProjectRepository{}

	service := NewService(repo)

	req := projectDTO.CreateRequest{
		Name:        "",
		Description: "My backend project",
		OwnerID:     1,
	}

	result, err := service.Create(
		context.Background(),
		req,
	)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}

	if repo.createCalled {
		t.Fatal("repository should not be called")
	}
}

func TestCreateProjectInvalidOwnerID(t *testing.T) {
	repo := &fakeProjectRepository{}

	service := NewService(repo)

	req := projectDTO.CreateRequest{
		Name:        "Project Tracker",
		Description: "My backend project",
		OwnerID:     0,
	}

	result, err := service.Create(
		context.Background(),
		req,
	)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}

	if repo.createCalled {
		t.Fatal("repository should not be called")
	}
}

func TestCreateProjectSuccess(t *testing.T) {
	repo := &fakeProjectRepository{}

	service := NewService(repo)

	req := projectDTO.CreateRequest{
		Name:        "Project Tracker",
		Description: "My backend project",
		OwnerID:     1,
	}

	result, err := service.Create(
		context.Background(),
		req,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}

	if result.Name != "Project Tracker" {
		t.Errorf(
			"expected project name Project Tracker, got %s",
			result.Name,
		)
	}

	if result.OwnerID != 1 {
		t.Errorf(
			"expected owner id 1, got %d",
			result.OwnerID,
		)
	}
}
