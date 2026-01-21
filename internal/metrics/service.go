package metrics

type Service struct {
	Repo Repository
}

func (s *Service) CollectMetrics() error {
	return nil
}
