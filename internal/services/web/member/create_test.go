package member

import (
	"context"
	"testing"

	"project-tracker-backend/internal/domain"
	memberDTO "project-tracker-backend/internal/dtos/member"
)

type fakeMemberRepository struct {
	createCalled  bool
	updateCalled  bool
	deleteCalled  bool
	members       []domain.Member
	contributions []memberDTO.ContributionResponse
}

func (f *fakeMemberRepository) GetContributionsByProjectID(
	ctx context.Context,
	projectID int64,
) ([]memberDTO.ContributionResponse, error) {

	return f.contributions, nil
}

func (f *fakeMemberRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	f.deleteCalled = true

	for i, member := range f.members {
		if member.ID == id {
			f.members = append(
				f.members[:i],
				f.members[i+1:]...,
			)

			return nil
		}
	}

	return nil
}

func (f *fakeMemberRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Member, error) {

	for i := range f.members {
		if f.members[i].ID == id {
			return &f.members[i], nil
		}
	}

	return nil, nil
}

func (f *fakeMemberRepository) Update(
	ctx context.Context,
	member *domain.Member,
) error {

	f.updateCalled = true

	for i := range f.members {
		if f.members[i].ID == member.ID {
			f.members[i] = *member
			return nil
		}
	}

	return nil
}

func (f *fakeMemberRepository) ListByProjectID(
	ctx context.Context,
	projectID int64,
) ([]domain.Member, error) {

	var result []domain.Member

	for _, member := range f.members {
		if member.ProjectID == projectID {
			result = append(result, member)
		}
	}

	return result, nil
}

func (f *fakeMemberRepository) Create(
	ctx context.Context,
	member *domain.Member,
) error {

	f.createCalled = true
	member.ID = 1

	return nil
}

func TestCreateMemberSuccess(t *testing.T) {

	repo := &fakeMemberRepository{}
	service := NewService(repo)

	req := memberDTO.CreateRequest{
		Name:  "Noor",
		Email: "noor@example.com",
		Role:  "Developer",
	}

	result, err := service.Create(
		context.Background(),
		1,
		req,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}

	if result.ProjectID != 1 {
		t.Errorf(
			"expected project id 1, got %d",
			result.ProjectID,
		)
	}

	if result.Name != "Noor" {
		t.Errorf(
			"expected Noor, got %s",
			result.Name,
		)
	}
}

// check if the repository Create method is not called when the project ID is invalid
func TestCreateMemberInvalidProjectID(t *testing.T) {

	repo := &fakeMemberRepository{}
	service := NewService(repo)

	req := memberDTO.CreateRequest{
		Name:  "Noor",
		Email: "noor@example.com",
		Role:  "Developer",
	}

	result, err := service.Create(
		context.Background(),
		0,
		req,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Fatal("expected result to be nil")
	}

	if repo.createCalled {
		t.Fatal("repository Create should not be called")
	}
}
