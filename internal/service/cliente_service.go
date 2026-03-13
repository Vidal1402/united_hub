package service

import (
	"context"
	"time"

	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/repository"
	"github.com/redis/go-redis/v9"
)

// Reexporta PageParams para ficar mais ergonômico nos services.
type PageParams = repository.PageParams

type ClienteService struct {
	clientes   repository.ClienteRepository
	kanban     repository.KanbanRepository
	relatorios repository.RelatorioRepository
	materiais  repository.MaterialRepository
	reunioes   repository.ReuniaoRepository
	financeiro repository.FinanceiroRepository
	cursos     repository.CursoRepository
	chamados   repository.ChamadoRepository
	alertas    repository.AlertaRepository
	notifs     repository.NotificacaoRepository
	cache      *redis.Client
}

func NewClienteService(
	clientes repository.ClienteRepository,
	kanban repository.KanbanRepository,
	relatorios repository.RelatorioRepository,
	materiais repository.MaterialRepository,
	reunioes repository.ReuniaoRepository,
	financeiro repository.FinanceiroRepository,
	cursos repository.CursoRepository,
	chamados repository.ChamadoRepository,
	alertas repository.AlertaRepository,
	notifs repository.NotificacaoRepository,
	cache *redis.Client,
) *ClienteService {
	return &ClienteService{
		clientes:   clientes,
		kanban:     kanban,
		relatorios: relatorios,
		materiais:  materiais,
		reunioes:   reunioes,
		financeiro: financeiro,
		cursos:     cursos,
		chamados:   chamados,
		alertas:    alertas,
		notifs:     notifs,
		cache:      cache,
	}
}

// Métodos de negócio básicos. Neste momento mantemos a implementação
// enxuta, retornando dados brutos das camadas de repositório.

// Produção
func (s *ClienteService) GetProducao(ctx context.Context, clienteUUID string) (any, error) {
	// Placeholder: aqui no futuro agregamos colunas + cards + métricas.
	cols, err := s.kanban.ListColumnsByCliente(ctx, clienteUUID)
	if err != nil {
		return nil, err
	}
	cards, _, err := s.kanban.ListCardsByCliente(ctx, clienteUUID, PageParams{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"columns": cols,
		"cards":   cards,
	}, nil
}

// Dashboard (com cache simples em Redis por alguns minutos)

func (s *ClienteService) GetDashboardChart(ctx context.Context, clienteUUID, period string) (any, error) {
	key := "client:" + clienteUUID + ":dashboard:chart:" + period
	if s.cache != nil {
		if v, err := s.cache.Get(ctx, key).Result(); err == nil {
			return map[string]any{"cached": true, "data": v}, nil
		}
	}
	// Placeholder: retornar estrutura vazia por enquanto
	result := map[string]any{
		"points": []any{},
	}
	if s.cache != nil {
		_ = s.cache.Set(ctx, key, "{}", 2*time.Minute).Err()
	}
	return result, nil
}

func (s *ClienteService) GetDashboardFunnel(ctx context.Context, clienteUUID string) (any, error) {
	key := "client:" + clienteUUID + ":dashboard:funnel"
	if s.cache != nil {
		if v, err := s.cache.Get(ctx, key).Result(); err == nil {
			return map[string]any{"cached": true, "data": v}, nil
		}
	}
	result := map[string]any{
		"stages": []any{},
	}
	if s.cache != nil {
		_ = s.cache.Set(ctx, key, "{}", 2*time.Minute).Err()
	}
	return result, nil
}

func (s *ClienteService) GetDashboardKPIs(ctx context.Context, clienteUUID string) (any, error) {
	key := "client:" + clienteUUID + ":dashboard:kpis"
	if s.cache != nil {
		if v, err := s.cache.Get(ctx, key).Result(); err == nil {
			return map[string]any{"cached": true, "data": v}, nil
		}
	}
	result := map[string]any{
		"kpis": map[string]any{},
	}
	if s.cache != nil {
		_ = s.cache.Set(ctx, key, "{}", 2*time.Minute).Err()
	}
	return result, nil
}

func (s *ClienteService) GetDashboardScore(ctx context.Context, clienteUUID string) (any, error) {
	key := "client:" + clienteUUID + ":dashboard:score"
	if s.cache != nil {
		if v, err := s.cache.Get(ctx, key).Result(); err == nil {
			return map[string]any{"cached": true, "data": v}, nil
		}
	}
	result := map[string]any{
		"score": 0,
	}
	if s.cache != nil {
		_ = s.cache.Set(ctx, key, "{}", 2*time.Minute).Err()
	}
	return result, nil
}

// Relatórios

func (s *ClienteService) ListRelatorios(ctx context.Context, clienteUUID string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.relatorios.ListByCliente(ctx, clienteUUID, page)
	if err != nil {
		return dto.Page[any]{}, err
	}
	// Aqui poderíamos mapear para DTO específico; por enquanto expomos domínio.
	out := make([]any, len(items))
	for i, it := range items {
		out[i] = it
	}
	return dto.Page[any]{
		Items:  out,
		Total:  total,
		Limit:  page.Limit,
		Offset: page.Offset,
	}, nil
}

// Materiais

func (s *ClienteService) ListMateriaisPastas(ctx context.Context, clienteUUID string) (any, error) {
	items, _, err := s.materiais.ListPastasByCliente(ctx, clienteUUID, PageParams{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ClienteService) ListMateriaisArquivos(ctx context.Context, clienteUUID, pastaUUID string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.materiais.ListArquivosByCliente(ctx, clienteUUID, page)
	if err != nil {
		return dto.Page[any]{}, err
	}
	out := make([]any, len(items))
	for i, it := range items {
		out[i] = it
	}
	return dto.Page[any]{
		Items:  out,
		Total:  total,
		Limit:  page.Limit,
		Offset: page.Offset,
	}, nil
}

func (s *ClienteService) UploadMaterial(ctx context.Context, clienteUUID string, input dto.UploadMaterialInput) error {
	// Implementação real faria persistência + contadores em cards etc.
	return nil
}

// Reuniões

func (s *ClienteService) ListReunioesProximas(ctx context.Context, clienteUUID string) (any, error) {
	items, _, err := s.reunioes.ListProximasByCliente(ctx, clienteUUID, PageParams{Limit: 50, Offset: 0})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ClienteService) ListReunioesHistorico(ctx context.Context, clienteUUID string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.reunioes.ListHistoricoByCliente(ctx, clienteUUID, page)
	if err != nil {
		return dto.Page[any]{}, err
	}
	out := make([]any, len(items))
	for i, it := range items {
		out[i] = it
	}
	return dto.Page[any]{
		Items:  out,
		Total:  total,
		Limit:  page.Limit,
		Offset: page.Offset,
	}, nil
}

// Financeiro

func (s *ClienteService) ListFaturas(ctx context.Context, clienteUUID string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.financeiro.ListFaturasByCliente(ctx, clienteUUID, page)
	if err != nil {
		return dto.Page[any]{}, err
	}
	out := make([]any, len(items))
	for i, it := range items {
		out[i] = it
	}
	return dto.Page[any]{
		Items:  out,
		Total:  total,
		Limit:  page.Limit,
		Offset: page.Offset,
	}, nil
}

func (s *ClienteService) GetPlanoFinanceiro(ctx context.Context, clienteUUID string) (any, error) {
	// Placeholder: poderia usar última fatura / plano do cliente.
	return map[string]any{}, nil
}

// Academy

func (s *ClienteService) ListCursos(ctx context.Context, clienteUUID string) (any, error) {
	key := "client:" + clienteUUID + ":academy:cursos"
	if s.cache != nil {
		if v, err := s.cache.Get(ctx, key).Result(); err == nil {
			return map[string]any{"cached": true, "data": v}, nil
		}
	}
	cursos, _, err := s.cursos.ListCursos(ctx, PageParams{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}
	if s.cache != nil {
		_ = s.cache.Set(ctx, key, "{}", 5*time.Minute).Err()
	}
	return cursos, nil
}

// Suporte

func (s *ClienteService) ListChamados(ctx context.Context, clienteUUID string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.chamados.ListByCliente(ctx, clienteUUID, page)
	if err != nil {
		return dto.Page[any]{}, err
	}
	out := make([]any, len(items))
	for i, it := range items {
		out[i] = it
	}
	return dto.Page[any]{
		Items:  out,
		Total:  total,
		Limit:  page.Limit,
		Offset: page.Offset,
	}, nil
}

func (s *ClienteService) CreateChamado(ctx context.Context, clienteUUID string, input dto.CreateChamadoInput) error {
	// Implementação real: criar domain.Chamado com UUID etc.
	return nil
}

func (s *ClienteService) ListFAQ(ctx context.Context, clienteUUID string) (any, error) {
	// Poderia ser hardcoded ou vir de coleção dedicada.
	return []any{}, nil
}

// Config

func (s *ClienteService) GetPerfil(ctx context.Context, clienteUUID string) (any, error) {
	c, err := s.clientes.GetByUUID(ctx, clienteUUID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *ClienteService) UpdatePerfil(ctx context.Context, clienteUUID string, input dto.UpdatePerfilInput) error {
	// Implementação real: carregar cliente, aplicar mudanças, salvar.
	return nil
}

func (s *ClienteService) ListUsuarios(ctx context.Context, clienteUUID string) (any, error) {
	// Poderia listar colaboradores associados ao cliente.
	return []any{}, nil
}

func (s *ClienteService) GetNotificacoesConfig(ctx context.Context, clienteUUID string) (any, error) {
	return map[string]any{}, nil
}

func (s *ClienteService) UpdateNotificacoesConfig(ctx context.Context, clienteUUID string, input dto.UpdateNotificacoesConfigInput) error {
	return nil
}

func (s *ClienteService) ListIntegracoes(ctx context.Context, clienteUUID string) (any, error) {
	return []any{}, nil
}

func (s *ClienteService) ConnectIntegracao(ctx context.Context, clienteUUID, integracaoID string) error {
	return nil
}

