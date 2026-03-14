package cliente

import (
	"encoding/json"
	"net/http"

	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/http/middleware"
	"backend_united_hub/internal/http/pagination"
	"backend_united_hub/internal/http/response"
	"backend_united_hub/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc       *service.ClienteService
	validator *validator.Validate
}

func New(svc *service.ClienteService, v *validator.Validate) *Handler {
	return &Handler{svc: svc, validator: v}
}

func (h *Handler) GetProducao(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	if !claims.CanProducao {
		response.Error(w, http.StatusForbidden, "forbidden", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetProducao(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateSolicitacao(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	if !claims.CanProducao {
		response.Error(w, http.StatusForbidden, "forbidden", nil)
		return
	}
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateSolicitacao(ctx, claims.ClienteID.String(), input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) GetDashboardChart(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	if !claims.CanPerformance {
		response.Error(w, http.StatusForbidden, "forbidden", nil)
		return
	}
	period := r.URL.Query().Get("period")
	ctx := r.Context()
	result, err := h.svc.GetDashboardChart(ctx, claims.ClienteID.String(), period)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetDashboardFunnel(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	if !claims.CanPerformance {
		response.Error(w, http.StatusForbidden, "forbidden", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetDashboardFunnel(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetDashboardKPIs(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	if !claims.CanPerformance {
		response.Error(w, http.StatusForbidden, "forbidden", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetDashboardKPIs(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetDashboardScore(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	if !claims.CanPerformance {
		response.Error(w, http.StatusForbidden, "forbidden", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetDashboardScore(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListRelatorios(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListRelatorios(ctx, claims.ClienteID.String(), service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListMateriaisPastas(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListMateriaisPastas(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListMateriaisArquivos(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListMateriaisArquivos(ctx, claims.ClienteID.String(), "", service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) UploadMaterial(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	var input dto.UploadMaterialInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if err := h.validator.Struct(input); err != nil {
		response.Error(w, http.StatusBadRequest, "validation error", nil)
		return
	}
	ctx := r.Context()
	if err := h.svc.UploadMaterial(ctx, claims.ClienteID.String(), input); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]string{"status": "ok"})
}

func (h *Handler) ListReunioesProximas(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListReunioesProximas(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListReunioesHistorico(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListReunioesHistorico(ctx, claims.ClienteID.String(), service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListFaturas(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListFaturas(ctx, claims.ClienteID.String(), service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetPlanoFinanceiro(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetPlanoFinanceiro(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListCursos(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListCursos(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListChamados(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListChamados(ctx, claims.ClienteID.String(), service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateChamado(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	var input dto.CreateChamadoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if err := h.validator.Struct(input); err != nil {
		response.Error(w, http.StatusBadRequest, "validation error", nil)
		return
	}
	ctx := r.Context()
	if err := h.svc.CreateChamado(ctx, claims.ClienteID.String(), input); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]string{"status": "ok"})
}

func (h *Handler) ListFAQ(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListFAQ(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetPerfil(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetPerfil(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) UpdatePerfil(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	var input dto.UpdatePerfilInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if err := h.validator.Struct(input); err != nil {
		response.Error(w, http.StatusBadRequest, "validation error", nil)
		return
	}
	ctx := r.Context()
	if err := h.svc.UpdatePerfil(ctx, claims.ClienteID.String(), input); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ListUsuarios(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListUsuarios(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetNotificacoesConfig(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetNotificacoesConfig(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) UpdateNotificacoesConfig(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	var input dto.UpdateNotificacoesConfigInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	ctx := r.Context()
	if err := h.svc.UpdateNotificacoesConfig(ctx, claims.ClienteID.String(), input); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ListIntegracoes(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListIntegracoes(ctx, claims.ClienteID.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ConnectIntegracao(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	// Pode vir como path param no futuro; por enquanto query param.
	integracaoID := r.URL.Query().Get("id")
	ctx := r.Context()
	if err := h.svc.ConnectIntegracao(ctx, claims.ClienteID.String(), integracaoID); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

