package project

import (
	"context"
	"testing"
	"time"

	"project-tracker-backend/internal/domain"
)

func TestListProjectsSuccess(t *testing.T) {

	repo := &fakeProjectRepository{
		projects: []domain.Project{
			{
				ID:       1,
				Name:     "Project Tracker",
				Deadline: time.Now().AddDate(0, 1, 0),
			},
			{
				ID:       2,
				Name:     "School System",
				Deadline: time.Now().AddDate(0, 1, 0),
			},
		},
	}

	service := NewService(repo)

	projects, err := service.List(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if len(projects) != 2 {
		t.Fatalf(
			"expected 2 projects, got %d",
			len(projects),
		)
	}

	if projects[0].Name != "Project Tracker" {
		t.Errorf(
			"expected Project Tracker, got %s",
			projects[0].Name,
		)
	}

	if projects[1].Name != "School System" {
		t.Errorf(
			"expected School System, got %s",
			projects[1].Name,
		)
	}
}
