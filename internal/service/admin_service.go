package service

import (
	"context"

	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/repository"
)

type AdminService struct {
	clientes   repository.ClienteRepository
	colabs     repository.ColaboradorRepository
	financeiro repository.FinanceiroRepository
	produtos   repository.ProdutoRepository
	alertas    repository.AlertaRepository
	notifs     repository.NotificacaoRepository
	relatorios repository.RelatorioRepository
	kanban     repository.KanbanRepository
}

func NewAdminService(
	clientes repository.ClienteRepository,
	colabs repository.ColaboradorRepository,
	financeiro repository.FinanceiroRepository,
	produtos repository.ProdutoRepository,
	alertas repository.AlertaRepository,
	notifs repository.NotificacaoRepository,
	relatorios repository.RelatorioRepository,
	kanban repository.KanbanRepository,
) *AdminService {
	return &AdminService{
		clientes:   clientes,
		colabs:     colabs,
		financeiro: financeiro,
		produtos:   produtos,
		alertas:    alertas,
		notifs:     notifs,
		relatorios: relatorios,
		kanban:     kanban,
	}
}

// Overview

func (s *AdminService) GetOverview(ctx context.Context) (any, error) {
	// Placeholder básico.
	return map[string]any{}, nil
}

func (s *AdminService) GetMRRMensal(ctx context.Context) (any, error) {
	return map[string]any{}, nil
}

// Clientes

func (s *AdminService) ListClientes(ctx context.Context, filtro any, page PageParams) (dto.Page[any], error) {
	items, total, err := s.clientes.List(ctx, page)
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

func (s *AdminService) CreateCliente(ctx context.Context, input dto.CreateClienteInput) (any, error) {
	return map[string]any{}, nil
}

func (s *AdminService) GetCliente(ctx context.Context, id string) (any, error) {
	c, err := s.clientes.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AdminService) UpdateCliente(ctx context.Context, id string, input dto.UpdateClienteInput) (any, error) {
	return map[string]any{}, nil
}

func (s *AdminService) DesativarCliente(ctx context.Context, id string) error {
	return s.clientes.Desativar(ctx, id)
}

// Colaboradores

func (s *AdminService) ListColaboradores(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.colabs.List(ctx, page)
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

func (s *AdminService) CreateColaborador(ctx context.Context, input any) (any, error) {
	return map[string]any{}, nil
}

func (s *AdminService) GetColaborador(ctx context.Context, id string) (any, error) {
	c, err := s.colabs.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AdminService) UpdateColaborador(ctx context.Context, id string, input any) (any, error) {
	return map[string]any{}, nil
}

// Financeiro

func (s *AdminService) ListReceber(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.financeiro.ListRecebiveis(ctx, page)
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

func (s *AdminService) ListPagar(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.financeiro.ListPagamentos(ctx, page)
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

func (s *AdminService) CreateLancamento(ctx context.Context, input any) (any, error) {
	return map[string]any{}, nil
}

func (s *AdminService) MarcarRecebivelPago(ctx context.Context, id string) error {
	return s.financeiro.MarkRecebivelPago(ctx, id)
}

// Produtos

func (s *AdminService) ListProdutosPorFamilia(ctx context.Context, familia string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.produtos.ListByFamilia(ctx, familia, page)
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

func (s *AdminService) CreateProduto(ctx context.Context, input any) (any, error) {
	return map[string]any{}, nil
}

func (s *AdminService) UpdateProduto(ctx context.Context, id string, input any) (any, error) {
	return map[string]any{}, nil
}

func (s *AdminService) DeleteProduto(ctx context.Context, id string) error {
	return s.produtos.Delete(ctx, id)
}

// Alertas

func (s *AdminService) ListAlertas(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.alertas.List(ctx, page)
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

func (s *AdminService) ResolverAlerta(ctx context.Context, id string) error {
	return s.alertas.Resolver(ctx, id)
}

// Notificações

func (s *AdminService) ListNotificacoesEnviadas(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.notifs.ListEnviadas(ctx, page)
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

func (s *AdminService) EnviarNotificacao(ctx context.Context, input any) error {
	return nil
}

// Relatórios / Comercial

func (s *AdminService) ListRelatoriosAdmin(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.relatorios.ListAdmin(ctx, page)
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

func (s *AdminService) GetComercial(ctx context.Context) (any, error) {
	return map[string]any{}, nil
}

