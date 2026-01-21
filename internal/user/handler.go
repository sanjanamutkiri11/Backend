package user

import "net/http"

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {}

type Handler struct {
	Service *Service
}
