package project

type Service struct {
	repo     ProjectRepository
	taskRepo TaskRepository
	cache    Cache
}

// ينشئ service
func NewService(
	repo ProjectRepository,
	taskRepo ...TaskRepository,
) *Service {

	service := &Service{
		repo: repo,
	}

	if len(taskRepo) > 0 {
		service.taskRepo = taskRepo[0]
	}

	return service
}

// يعطي redis
func (s *Service) SetCache(
	cache Cache,
) {
	s.cache = cache
}
