package scaling

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/scaling/events", h.GetEvents)
}
