package metrics

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/metrics", h.GetMetrics)
}
