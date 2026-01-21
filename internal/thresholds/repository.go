package thresholds

type Repository interface {
	Get() (*Thresholds, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Get() (*Thresholds, error) { return nil, nil }
