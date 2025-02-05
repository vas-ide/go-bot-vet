package care

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Care {
	return allEntities
}

func (s *Service) Get(idx int) (*Care, error) {
	return &allEntities[idx], nil
}
