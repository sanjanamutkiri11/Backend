package thresholds

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/thresholds", h.GetThresholds)
}
