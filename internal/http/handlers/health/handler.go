package health

import (
  "net/http"

  "backend_united_hub/internal/http/response"
)

type Handler struct{}

func New() *Handler { return &Handler{} }

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
  response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}