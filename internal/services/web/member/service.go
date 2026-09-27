package member

type Service struct {
	repo MemberRepository
}

func NewService(repo MemberRepository) *Service {
	return &Service{
		repo: repo,
	}
}
