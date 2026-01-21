package metrics

import "time"

type Scheduler struct {
	Service *Service
}

func (s *Scheduler) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			s.Service.CollectMetrics()
		}
	}()
}
