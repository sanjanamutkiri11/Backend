package scaling

import "time"

type ScalingEvent struct {
	ID        int       `json:"id"`
	Action    string    `json:"action"` // "UP" or "DOWN"
	Timestamp time.Time `json:"timestamp"`
}
