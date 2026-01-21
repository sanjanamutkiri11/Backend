package thresholds

type Service struct {
	Repo Repository
}

func (s *Service) GetThresholds() (*Thresholds, error) {
	return nil, nil
}
