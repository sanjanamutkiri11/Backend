package auth

type Service struct {
	Repo Repository
}

func (s *Service) Authenticate(username, password string) (string, error) {
	return "", nil
}
