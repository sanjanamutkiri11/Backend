package auth

import "net/http"

func (h *Handler) Login(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {}

type Handler struct {
	Service *Service
}
