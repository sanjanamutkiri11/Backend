package thresholds

type Thresholds struct {
	CPUHigh    float64 `json:"cpu_high"`
	CPULow     float64 `json:"cpu_low"`
	MemoryHigh float64 `json:"memory_high"`
	MemoryLow  float64 `json:"memory_low"`
}
