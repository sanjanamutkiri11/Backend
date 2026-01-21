package metrics

import "time"

type Metric struct {
	ID        int       `json:"id"`
	CPU       float64   `json:"cpu"`
	Memory    float64   `json:"memory"`
	Timestamp time.Time `json:"timestamp"`
}
