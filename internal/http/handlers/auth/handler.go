package authhandler

import (
	"encoding/json"
	"net/http"

	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/http/middleware"
	"backend_united_hub/internal/http/response"
	"backend_united_hub/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc       *service.AuthService
	validator *validator.Validate
}

func New(svc *service.AuthService, v *validator.Validate) *Handler {
	return &Handler{svc: svc, validator: v}
}

// POST /api/auth/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var in dto.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if err := h.validator.Struct(in); err != nil {
		response.Error(w, http.StatusBadRequest, "validation error", nil)
		return
	}
	ctx := r.Context()
	resp, err := h.svc.Login(ctx, in)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid credentials", nil)
		return
	}
	response.JSON(w, http.StatusOK, resp)
}

// GET /api/auth/me
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	info, err := h.svc.Me(ctx, claims)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, info)
}

