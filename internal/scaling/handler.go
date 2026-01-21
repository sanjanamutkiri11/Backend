package scaling

import "net/http"

func (h *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {}

type Handler struct {
	Service *Service
}
