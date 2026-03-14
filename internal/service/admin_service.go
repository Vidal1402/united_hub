package service

import (
	"context"
	"errors"

	"backend_united_hub/internal/domain"
	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	usuarios   repository.UsuarioRepository
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
	usuarios repository.UsuarioRepository,
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
		usuarios:   usuarios,
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

func (s *AdminService) CreateCliente(ctx context.Context, input dto.CreateClienteInput) (dto.ClienteOutput, error) {
	c := &domain.Cliente{
		UUID:      uuid.New().String(),
		Nome:      input.Nome,
		Email:     input.Email,
		Segmento:  input.Segmento,
		Plano:     input.Plano,
		Status:    input.Status,
		Cidade:    input.Cidade,
		OwnerUUID: input.OwnerUUID,
	}
	if err := s.clientes.Create(ctx, c); err != nil {
		return dto.ClienteOutput{}, err
	}
	return dto.ClienteOutput{
		UUID:      c.UUID,
		Nome:      c.Nome,
		Email:     c.Email,
		Segmento:  c.Segmento,
		Plano:     c.Plano,
		Status:    c.Status,
		Cidade:    c.Cidade,
		OwnerUUID: c.OwnerUUID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (s *AdminService) GetCliente(ctx context.Context, id string) (any, error) {
	c, err := s.clientes.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AdminService) UpdateCliente(ctx context.Context, id string, input dto.UpdateClienteInput) (dto.ClienteOutput, error) {
	existing, err := s.clientes.GetByUUID(ctx, id)
	if err != nil {
		return dto.ClienteOutput{}, err
	}
	if existing == nil {
		return dto.ClienteOutput{}, errors.New("cliente not found")
	}
	existing.Nome = input.Nome
	existing.Email = input.Email
	existing.Segmento = input.Segmento
	existing.Plano = input.Plano
	existing.Status = input.Status
	existing.Cidade = input.Cidade
	existing.OwnerUUID = input.OwnerUUID
	if err := s.clientes.Update(ctx, existing); err != nil {
		return dto.ClienteOutput{}, err
	}
	return dto.ClienteOutput{
		UUID:      existing.UUID,
		Nome:      existing.Nome,
		Email:     existing.Email,
		Segmento:  existing.Segmento,
		Plano:     existing.Plano,
		Status:    existing.Status,
		Cidade:    existing.Cidade,
		OwnerUUID: existing.OwnerUUID,
		CreatedAt: existing.CreatedAt,
		UpdatedAt: existing.UpdatedAt,
	}, nil
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

// Usuários (auth/admin)

func (s *AdminService) CreateUsuario(ctx context.Context, input dto.UsuarioCreateInput) (dto.UsuarioOutput, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UsuarioOutput{}, err
	}

	u := domain.Usuario{
		UUID:           uuid.NewString(),
		ClienteUUID:    input.ClienteUUID,
		Email:          input.Email,
		SenhaHash:      string(hash),
		Role:           input.Role,
		CanProducao:    input.CanProducao,
		CanPerformance: input.CanPerformance,
	}
	if err := s.usuarios.Create(ctx, &u); err != nil {
		return dto.UsuarioOutput{}, err
	}

	return dto.UsuarioOutput{
		UUID:           u.UUID,
		ClienteUUID:    u.ClienteUUID,
		Email:          u.Email,
		Role:           u.Role,
		CanProducao:    u.CanProducao,
		CanPerformance: u.CanPerformance,
	}, nil
}

func (s *AdminService) ListUsuarios(ctx context.Context, page PageParams) (dto.Page[dto.UsuarioOutput], error) {
	items, total, err := s.usuarios.List(ctx, page)
	if err != nil {
		return dto.Page[dto.UsuarioOutput]{}, err
	}
	out := make([]dto.UsuarioOutput, len(items))
	for i, u := range items {
		out[i] = dto.UsuarioOutput{
			UUID:           u.UUID,
			ClienteUUID:    u.ClienteUUID,
			Email:          u.Email,
			Role:           u.Role,
			CanProducao:    u.CanProducao,
			CanPerformance: u.CanPerformance,
		}
	}
	return dto.Page[dto.UsuarioOutput]{
		Items:  out,
		Total:  total,
		Limit:  page.Limit,
		Offset: page.Offset,
	}, nil
}

func (s *AdminService) UpdateUsuario(ctx context.Context, id string, input dto.UsuarioUpdateInput) (dto.UsuarioOutput, error) {
	u, err := s.usuarios.GetByUUID(ctx, id)
	if err != nil {
		return dto.UsuarioOutput{}, err
	}
	if u == nil {
		return dto.UsuarioOutput{}, errors.New("usuario not found")
	}

	if input.CanProducao != nil {
		u.CanProducao = *input.CanProducao
	}
	if input.CanPerformance != nil {
		u.CanPerformance = *input.CanPerformance
	}

	if err := s.usuarios.Update(ctx, u); err != nil {
		return dto.UsuarioOutput{}, err
	}

	return dto.UsuarioOutput{
		UUID:           u.UUID,
		ClienteUUID:    u.ClienteUUID,
		Email:          u.Email,
		Role:           u.Role,
		CanProducao:    u.CanProducao,
		CanPerformance: u.CanPerformance,
	}, nil
}

