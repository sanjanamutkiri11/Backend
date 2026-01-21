package thresholds

import "net/http"

func (h *Handler) GetThresholds(w http.ResponseWriter, r *http.Request) {}

type Handler struct {
	Service *Service
}
