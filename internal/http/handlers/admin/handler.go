package admin

import (
	"encoding/json"
	"net/http"

	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/http/pagination"
	"backend_united_hub/internal/http/response"
	"backend_united_hub/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc       *service.AdminService
	validator *validator.Validate
}

func New(svc *service.AdminService, v *validator.Validate) *Handler {
	return &Handler{svc: svc, validator: v}
}

func (h *Handler) GetOverview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.svc.GetOverview(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetMRRMensal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.svc.GetMRRMensal(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListClientes(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListClientes(ctx, nil, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateCliente(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input dto.CreateClienteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateCliente(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) GetCliente(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	result, err := h.svc.GetCliente(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) UpdateCliente(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	var input dto.UpdateClienteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.UpdateCliente(ctx, id, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) DesativarCliente(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if err := h.svc.DesativarCliente(ctx, id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ListColaboradores(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListColaboradores(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateColaborador(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateColaborador(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) GetColaborador(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	result, err := h.svc.GetColaborador(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) UpdateColaborador(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.UpdateColaborador(ctx, id, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListReceber(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListReceber(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListPagar(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListPagar(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateLancamento(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateLancamento(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) MarcarRecebivelPago(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if err := h.svc.MarcarRecebivelPago(ctx, id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ListProdutosPorFamilia(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	familia := chi.URLParam(r, "familia")
	result, err := h.svc.ListProdutosPorFamilia(ctx, familia, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateProduto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	familia := chi.URLParam(r, "familia")
	if familia == "" {
		response.Error(w, http.StatusBadRequest, "familia is required", nil)
		return
	}
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateProduto(ctx, familia, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) UpdateProduto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.UpdateProduto(ctx, id, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) DeleteProduto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if err := h.svc.DeleteProduto(ctx, id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func (h *Handler) ListAlertas(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListAlertas(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ResolverAlerta(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if err := h.svc.ResolverAlerta(ctx, id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ListNotificacoesEnviadas(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListNotificacoesEnviadas(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) EnviarNotificacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if err := h.svc.EnviarNotificacao(ctx, input); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, map[string]string{"status": "ok"})
}

func (h *Handler) ListRelatoriosAdmin(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListRelatoriosAdmin(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateRelatorio(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateRelatorio(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) CreateMaterialPasta(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateMaterialPasta(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) CreateMaterialArquivo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateMaterialArquivo(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) ListPastasAdmin(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	clienteUUID := r.URL.Query().Get("cliente_uuid")
	if clienteUUID == "" {
		response.Error(w, http.StatusBadRequest, "cliente_uuid is required", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListPastasAdmin(ctx, clienteUUID, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListArquivosAdmin(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	clienteUUID := r.URL.Query().Get("cliente_uuid")
	if clienteUUID == "" {
		response.Error(w, http.StatusBadRequest, "cliente_uuid is required", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.ListArquivosAdmin(ctx, clienteUUID, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) ListReunioesAdmin(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 50, 100)
	ctx := r.Context()
	result, err := h.svc.ListReunioesAdmin(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateReuniao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateReuniao(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) ListChamadosAdmin(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 50, 100)
	ctx := r.Context()
	result, err := h.svc.ListChamadosAdmin(ctx, service.PageParams{
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
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateChamado(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) GetProducaoAdmin(w http.ResponseWriter, r *http.Request) {
	clienteUUID := r.URL.Query().Get("cliente_uuid")
	if clienteUUID == "" {
		response.Error(w, http.StatusBadRequest, "cliente_uuid is required", nil)
		return
	}
	ctx := r.Context()
	result, err := h.svc.GetProducaoAdmin(ctx, clienteUUID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) CreateProducaoCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.CreateProducaoCard(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) MoveProducaoCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cardID := chi.URLParam(r, "id")
	var body struct {
		ColumnID string `json:"column_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if body.ColumnID == "" {
		response.Error(w, http.StatusBadRequest, "column_id is required", nil)
		return
	}
	result, err := h.svc.MoveProducaoCard(ctx, cardID, body.ColumnID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) GetComercial(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.svc.GetComercial(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

// Usuários (admin)

func (h *Handler) CreateUsuario(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input dto.UsuarioCreateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if err := h.validator.Struct(input); err != nil {
		response.Error(w, http.StatusBadRequest, "validation error", nil)
		return
	}
	result, err := h.svc.CreateUsuario(ctx, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) ListUsuarios(w http.ResponseWriter, r *http.Request) {
	pag := pagination.Parse(r, 20, 100)
	ctx := r.Context()
	result, err := h.svc.ListUsuarios(ctx, service.PageParams{
		Limit:  pag.Limit,
		Offset: pag.Offset,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	var input dto.UsuarioUpdateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	result, err := h.svc.UpdateUsuario(ctx, id, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.JSON(w, http.StatusOK, result)
}


