package user

type Service struct {
	Repo Repository
}

func (s *Service) GetUserByID(id int) (*User, error) {
	return nil, nil
}
