package system

import "net/http"

func (h *Handler) GetState(w http.ResponseWriter, r *http.Request) {}

type Handler struct {
	Service *Service
}
