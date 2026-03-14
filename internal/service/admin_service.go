package service

import (
	"context"
	"errors"
	"time"

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
	materiais  repository.MaterialRepository
	reunioes   repository.ReuniaoRepository
	chamados   repository.ChamadoRepository
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
	materiais repository.MaterialRepository,
	reunioes repository.ReuniaoRepository,
	chamados repository.ChamadoRepository,
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
		materiais:  materiais,
		reunioes:   reunioes,
		chamados:   chamados,
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

func (s *AdminService) CreateCliente(ctx context.Context, input dto.CreateClienteInput) (any, error) {
	id := uuid.New().String()
	c := &domain.Cliente{
		UUID:      id,
		Nome:      input.Nome,
		Email:     input.Email,
		Segmento:  input.Segmento,
		Plano:     input.Plano,
		Status:    input.Status,
		Cidade:    input.Cidade,
		OwnerUUID: input.OwnerUUID,
	}
	if err := s.clientes.Create(ctx, c); err != nil {
		return nil, err
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

func (s *AdminService) UpdateCliente(ctx context.Context, id string, input dto.UpdateClienteInput) (any, error) {
	existing, err := s.clientes.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	existing.Nome = input.Nome
	existing.Email = input.Email
	existing.Segmento = input.Segmento
	existing.Plano = input.Plano
	existing.Status = input.Status
	existing.Cidade = input.Cidade
	existing.OwnerUUID = input.OwnerUUID
	if err := s.clientes.Update(ctx, existing); err != nil {
		return nil, err
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
	m, _ := input.(map[string]any)
	str := func(k string) string {
		if v, ok := m[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	id := uuid.New().String()
	c := &domain.Colaborador{
		UUID:   id,
		Nome:   str("nome"),
		Email:  str("email"),
		Cargo:  str("cargo"),
		Role:   str("role"),
		Status: str("status"),
	}
	if c.Role == "" {
		c.Role = "Colaborador"
	}
	if c.Status == "" {
		c.Status = "Ativo"
	}
	if err := s.colabs.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AdminService) GetColaborador(ctx context.Context, id string) (any, error) {
	c, err := s.colabs.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AdminService) UpdateColaborador(ctx context.Context, id string, input any) (any, error) {
	existing, err := s.colabs.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("colaborador não encontrado")
	}
	m, _ := input.(map[string]any)
	str := func(k string) string {
		if v, ok := m[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	existing.Nome = str("nome")
	existing.Email = str("email")
	existing.Cargo = str("cargo")
	existing.Role = str("role")
	existing.Status = str("status")
	if existing.Role == "" {
		existing.Role = "Colaborador"
	}
	if existing.Status == "" {
		existing.Status = "Ativo"
	}
	if err := s.colabs.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
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
	m, _ := input.(map[string]any)
	str := func(k string) string {
		if v, ok := m[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	valorCentavos := int64(0)
	if v, ok := m["valor_centavos"]; ok && v != nil {
		switch n := v.(type) {
		case float64:
			valorCentavos = int64(n)
		case int:
			valorCentavos = int64(n)
		case int64:
			valorCentavos = n
		}
	}
	if valorCentavos == 0 && m["valor"] != nil {
		if n, ok := m["valor"].(float64); ok {
			valorCentavos = int64(n * 100)
		}
	}
	var vencimento time.Time
	if s := str("vencimento"); len(s) >= 10 {
		vencimento, _ = time.Parse("2006-01-02", s[:10])
	}
	if vencimento.IsZero() {
		vencimento = time.Now().UTC().AddDate(0, 1, 0)
	}
	rec := &domain.Recebivel{
		UUID:        uuid.New().String(),
		ClienteUUID: str("cliente_uuid"),
		Descricao:   str("descricao"),
		Valor:       valorCentavos,
		Vencimento:  vencimento,
		Status:      "pendente",
		Plano:       str("plano"),
	}
	if rec.Descricao == "" {
		rec.Descricao = "Lançamento"
	}
	if err := s.financeiro.CreateRecebivel(ctx, rec); err != nil {
		return nil, err
	}
	return rec, nil
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

func (s *AdminService) CreateProduto(ctx context.Context, familia string, input any) (any, error) {
	m, _ := input.(map[string]any)
	str := func(k string) string {
		if v, ok := m[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	precoCentavos := int64(0)
	if v, ok := m["price"]; ok && v != nil {
		switch n := v.(type) {
		case float64:
			precoCentavos = int64(n * 100)
		case int:
			precoCentavos = int64(n) * 100
		case int64:
			precoCentavos = n * 100
		}
	}
	if precoCentavos == 0 && m["preco_centavos"] != nil {
		if n, ok := m["preco_centavos"].(float64); ok {
			precoCentavos = int64(n)
		}
	}
	featuresStr := str("features")
	var features []string
	if featuresStr != "" {
		for _, line := range splitLines(featuresStr) {
			if line != "" {
				features = append(features, line)
			}
		}
	}
	id := uuid.New().String()
	slug := str("slug")
	if slug == "" {
		slug = id[:8]
	}
	p := &domain.Produto{
		UUID:      id,
		Familia:   familia,
		Slug:      slug,
		Nome:     str("name"),
		Preco:     precoCentavos,
		Badge:     str("badge"),
		Descricao: str("descricao"),
		Features:  features,
	}
	if p.Nome == "" {
		p.Nome = str("nome")
	}
	if err := s.produtos.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func splitLines(s string) []string {
	var out []string
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '\n' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	return out
}

func (s *AdminService) UpdateProduto(ctx context.Context, id string, input any) (any, error) {
	existing, err := s.produtos.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("produto não encontrado")
	}
	m, _ := input.(map[string]any)
	str := func(k string) string {
		if v, ok := m[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	if v := str("name"); v != "" {
		existing.Nome = v
	} else if v := str("nome"); v != "" {
		existing.Nome = v
	}
	if v, ok := m["price"]; ok && v != nil {
		if n, ok := v.(float64); ok {
			existing.Preco = int64(n * 100)
		}
	}
	if m["preco_centavos"] != nil {
		if n, ok := m["preco_centavos"].(float64); ok {
			existing.Preco = int64(n)
		}
	}
	existing.Badge = str("badge")
	existing.Descricao = str("descricao")
	featuresStr := str("features")
	if featuresStr != "" {
		var features []string
		for _, line := range splitLines(featuresStr) {
			if line != "" {
				features = append(features, line)
			}
		}
		existing.Features = features
	}
	if err := s.produtos.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
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
	m, _ := input.(map[string]any)
	str := func(k string) string {
		if v, ok := m[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	n := &domain.Notificacao{
		UUID:     uuid.New().String(),
		Titulo:   str("titulo"),
		Target:   str("target"),
		Canal:    str("canal"),
		Tipo:     "sistema",
		Conteudo: str("conteudo"),
	}
	if n.Titulo == "" {
		n.Titulo = "Notificação"
	}
	return s.notifs.Enviar(ctx, n)
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

func (s *AdminService) CreateRelatorio(ctx context.Context, input map[string]any) (any, error) {
	str := func(k string) string {
		if v, ok := input[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	paginas := 0
	if v, ok := input["paginas"]; ok && v != nil {
		switch n := v.(type) {
		case float64:
			paginas = int(n)
		case int:
			paginas = n
		}
	}
	r := &domain.Relatorio{
		UUID:        uuid.New().String(),
		ClienteUUID: str("cliente_uuid"),
		Titulo:      str("titulo"),
		Tipo:        str("tipo"),
		Periodo:     str("periodo"),
		FileURL:     str("file_url"),
		Paginas:     paginas,
		OwnerUUID:   str("owner_uuid"),
	}
	if r.Tipo == "" {
		r.Tipo = "Mensal"
	}
	if err := s.relatorios.Create(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Materiais (pastas e arquivos)

func (s *AdminService) CreateMaterialPasta(ctx context.Context, input map[string]any) (any, error) {
	str := func(k string) string {
		if v, ok := input[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	p := &domain.MaterialPasta{
		UUID:        uuid.New().String(),
		ClienteUUID: str("cliente_uuid"),
		ParentUUID:  str("parent_uuid"), // opcional: se presente, pasta é subpasta
		Nome:        str("nome"),
		Icone:       str("icone"),
	}
	if err := s.materiais.CreatePasta(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *AdminService) CreateMaterialArquivo(ctx context.Context, input map[string]any) (any, error) {
	str := func(k string) string {
		if v, ok := input[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	tamanho := int64(0)
	if v, ok := input["tamanho"]; ok && v != nil {
		switch n := v.(type) {
		case float64:
			tamanho = int64(n)
		case int:
			tamanho = int64(n)
		case int64:
			tamanho = n
		}
	}
	a := &domain.MaterialArquivo{
		UUID:        uuid.New().String(),
		ClienteUUID: str("cliente_uuid"),
		PastaUUID:   str("pasta_uuid"),
		Nome:        str("nome"),
		URL:         str("url"),
		Extensao:    str("extensao"),
		Tamanho:     tamanho,
		Data:        time.Now().UTC(),
	}
	if err := s.materiais.CreateArquivo(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AdminService) ListPastasAdmin(ctx context.Context, clienteUUID string, page PageParams) (dto.Page[any], error) {
	items, total, err := s.materiais.ListPastasByCliente(ctx, clienteUUID, page)
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

func (s *AdminService) ListArquivosAdmin(ctx context.Context, clienteUUID string, page PageParams) (dto.Page[any], error) {
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

// Reuniões

func (s *AdminService) CreateReuniao(ctx context.Context, input map[string]any) (any, error) {
	str := func(k string) string {
		if v, ok := input[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	duracaoMin := 0
	if v, ok := input["duracao_min"]; ok && v != nil {
		switch n := v.(type) {
		case float64:
			duracaoMin = int(n)
		case int:
			duracaoMin = n
		}
	}
	var pauta []string
	if v, ok := input["pauta"]; ok && v != nil {
		switch arr := v.(type) {
		case []interface{}:
			for _, x := range arr {
				if s, ok := x.(string); ok {
					pauta = append(pauta, s)
				}
			}
		case []string:
			pauta = arr
		case string:
			for _, line := range splitLines(arr) {
				if line != "" {
					pauta = append(pauta, line)
				}
			}
		}
	}
	dataHora := time.Now().UTC()
	if s := str("data_hora"); s != "" {
		// ISO ou "2006-01-02T15:04:05"
		for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05", "2006-01-02 15:04:05", "2006-01-02"} {
			if t, err := time.Parse(layout, s); err == nil {
				dataHora = t.UTC()
				break
			}
		}
	}
	status := "futura"
	if dataHora.Before(time.Now().UTC()) {
		status = "historico"
	}
	reun := &domain.Reuniao{
		UUID:        uuid.New().String(),
		ClienteUUID: str("cliente_uuid"),
		Titulo:      str("titulo"),
		DataHora:    dataHora,
		Via:         str("via"),
		OwnerUUID:   str("owner_uuid"),
		Pauta:       pauta,
		Status:      status,
		DuracaoMin:  duracaoMin,
	}
	if err := s.reunioes.Create(ctx, reun); err != nil {
		return nil, err
	}
	return reun, nil
}

func (s *AdminService) ListReunioesAdmin(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.reunioes.ListAdmin(ctx, page)
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

// Chamados

func (s *AdminService) CreateChamado(ctx context.Context, input map[string]any) (any, error) {
	str := func(k string) string {
		if v, ok := input[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	c := &domain.Chamado{
		UUID:        uuid.New().String(),
		ClienteUUID: str("cliente_uuid"),
		Titulo:      str("titulo"),
		Descricao:   str("descricao"),
		Categoria:   str("categoria"),
		Status:      "aberto",
	}
	if c.Categoria == "" {
		c.Categoria = "Suporte"
	}
	if err := s.chamados.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AdminService) ListChamadosAdmin(ctx context.Context, page PageParams) (dto.Page[any], error) {
	items, total, err := s.chamados.ListAdmin(ctx, page)
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

// Produção (Kanban) — admin vê quadro por cliente, cria e move cards

func (s *AdminService) GetProducaoAdmin(ctx context.Context, clienteUUID string) (any, error) {
	if clienteUUID == "" {
		return nil, errors.New("cliente_uuid é obrigatório")
	}
	cards, _, err := s.kanban.ListCardsByCliente(ctx, clienteUUID, repository.PageParams{Limit: 500, Offset: 0})
	if err != nil {
		return nil, err
	}
	columnIDs := []string{"backlog", "doing", "review", "done"}
	byColumn := make(map[string][]map[string]any)
	for _, id := range columnIDs {
		byColumn[id] = nil
	}
	for i := range cards {
		c := &cards[i]
		colID := c.ColumnID
		if colID == "" {
			colID = "backlog"
		}
		if byColumn[colID] == nil {
			byColumn[colID] = make([]map[string]any, 0)
		}
		byColumn[colID] = append(byColumn[colID], cardToProducaoItem(c))
	}
	columns := make([]map[string]any, 0, len(producaoColumns))
	for _, col := range producaoColumns {
		id := col["id"].(string)
		columns = append(columns, map[string]any{
			"id":    id,
			"label": col["label"],
			"dot":   col["dot"],
			"cards": byColumn[id],
		})
	}
	return map[string]any{"columns": columns}, nil
}

func (s *AdminService) CreateProducaoCard(ctx context.Context, input map[string]any) (any, error) {
	str := func(k string) string {
		if v, ok := input[k]; ok && v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
		return ""
	}
	clienteUUID := str("cliente_uuid")
	if clienteUUID == "" {
		return nil, errors.New("cliente_uuid é obrigatório")
	}
	columnID := str("column_id")
	if columnID == "" {
		columnID = "backlog"
	}
	c := &domain.KanbanCard{
		UUID:        uuid.New().String(),
		ClienteUUID:  clienteUUID,
		ColumnID:    columnID,
		Titulo:      str("title"),
		Tipo:        str("type"),
		Prioridade:  str("priority"),
		OwnerNome:   str("owner"),
		Comentarios: 0,
		Arquivos:    0,
	}
	if c.Titulo == "" {
		c.Titulo = str("titulo")
	}
	if c.Tipo == "" {
		c.Tipo = "Campanha"
	}
	if c.Prioridade == "" {
		c.Prioridade = "Média"
	}
	if err := s.kanban.CreateCard(ctx, c); err != nil {
		return nil, err
	}
	return cardToProducaoItem(c), nil
}

func (s *AdminService) MoveProducaoCard(ctx context.Context, cardID string, columnID string) (any, error) {
	if cardID == "" || columnID == "" {
		return nil, errors.New("id do card e column_id são obrigatórios")
	}
	card, err := s.kanban.GetCardByUUID(ctx, cardID)
	if err != nil {
		return nil, err
	}
	if card == nil {
		return nil, errors.New("card não encontrado")
	}
	card.ColumnID = columnID
	if err := s.kanban.UpdateCard(ctx, card); err != nil {
		return nil, err
	}
	return cardToProducaoItem(card), nil
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

