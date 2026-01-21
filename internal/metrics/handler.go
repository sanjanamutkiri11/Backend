package metrics

import "net/http"

func (h *Handler) GetMetrics(w http.ResponseWriter, r *http.Request) {}

type Handler struct {
	Service *Service
}
