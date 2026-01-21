package user

type Repository interface {
	GetByID(id int) (*User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetByID(id int) (*User, error) {
	return nil, nil
}
