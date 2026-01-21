package auth

type Repository interface {
	GetUserByUsername(username string) (*User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetUserByUsername(username string) (*User, error) {
	return nil, nil
}
