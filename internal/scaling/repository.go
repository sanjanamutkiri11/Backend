package scaling

type Repository interface {
	SaveEvent(e *ScalingEvent) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) SaveEvent(e *ScalingEvent) error { return nil }
