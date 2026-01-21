package system

type Service struct {
	Repo Repository
}

func (s *Service) GetServerCount() (int, error) {
	return 0, nil
}
