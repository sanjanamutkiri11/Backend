package metrics

import "time"

type Collector interface {
	Collect() (*Metric, error)
}

type collector struct{}

func NewCollector() Collector {
	return &collector{}
}

func (c *collector) Collect() (*Metric, error) {
	return &Metric{Timestamp: time.Now()}, nil
}
