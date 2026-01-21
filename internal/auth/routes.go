package auth

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/auth/login", h.Login)
	mux.HandleFunc("/auth/register", h.Register)
}
