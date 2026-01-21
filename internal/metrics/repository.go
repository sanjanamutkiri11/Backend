package metrics

type Repository interface {
	Save(m *Metric) error
	GetLatest() (*Metric, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Save(m *Metric) error        { return nil }
func (r *repository) GetLatest() (*Metric, error) { return nil, nil }
