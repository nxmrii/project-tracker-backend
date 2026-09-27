package task

type Service struct {
	repo  TaskRepository
	cache Cache
}

func NewService(repo TaskRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SetCache(cache Cache) {
	s.cache = cache
}
