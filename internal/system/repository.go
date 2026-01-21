package system

type Repository interface {
	GetCount() (int, error)
	UpdateCount(count int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetCount() (int, error)      { return 0, nil }
func (r *repository) UpdateCount(count int) error { return nil }
