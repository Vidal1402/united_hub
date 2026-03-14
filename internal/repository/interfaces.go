package repository

import (
	"context"

	"backend_united_hub/internal/domain"
)

// PageParams representa os parâmetros de paginação básicos.
type PageParams struct {
	Limit  int
	Offset int
}

type ClienteRepository interface {
	GetByUUID(ctx context.Context, uuid string) (*domain.Cliente, error)
	List(ctx context.Context, pag PageParams) ([]domain.Cliente, int64, error)
	ListByOwner(ctx context.Context, ownerUUID string, pag PageParams) ([]domain.Cliente, int64, error)
	Create(ctx context.Context, c *domain.Cliente) error
	Update(ctx context.Context, c *domain.Cliente) error
	Desativar(ctx context.Context, uuid string) error
}

type ColaboradorRepository interface {
	GetByUUID(ctx context.Context, uuid string) (*domain.Colaborador, error)
	List(ctx context.Context, pag PageParams) ([]domain.Colaborador, int64, error)
	Create(ctx context.Context, c *domain.Colaborador) error
	Update(ctx context.Context, c *domain.Colaborador) error
}

type KanbanRepository interface {
	ListColumnsByCliente(ctx context.Context, clienteUUID string) ([]domain.KanbanColumn, error)
	ListCardsByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.KanbanCard, int64, error)
	CreateCard(ctx context.Context, c *domain.KanbanCard) error
	GetCardByUUID(ctx context.Context, uuid string) (*domain.KanbanCard, error)
	UpdateCard(ctx context.Context, c *domain.KanbanCard) error
}

type RelatorioRepository interface {
	ListByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Relatorio, int64, error)
	ListAdmin(ctx context.Context, pag PageParams) ([]domain.Relatorio, int64, error)
	Create(ctx context.Context, r *domain.Relatorio) error
}

type MaterialRepository interface {
	ListPastasByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.MaterialPasta, int64, error)
	ListArquivosByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.MaterialArquivo, int64, error)
	CreatePasta(ctx context.Context, p *domain.MaterialPasta) error
	CreateArquivo(ctx context.Context, a *domain.MaterialArquivo) error
	GetPastaByUUID(ctx context.Context, uuid string) (*domain.MaterialPasta, error)
	UpdatePasta(ctx context.Context, p *domain.MaterialPasta) error
	GetArquivoByUUID(ctx context.Context, uuid string) (*domain.MaterialArquivo, error)
	UpdateArquivo(ctx context.Context, a *domain.MaterialArquivo) error
}

type ReuniaoRepository interface {
	ListProximasByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Reuniao, int64, error)
	ListHistoricoByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Reuniao, int64, error)
	ListAdmin(ctx context.Context, pag PageParams) ([]domain.Reuniao, int64, error)
	Create(ctx context.Context, r *domain.Reuniao) error
}

type FinanceiroRepository interface {
	ListFaturasByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Fatura, int64, error)
	ListRecebiveis(ctx context.Context, pag PageParams) ([]domain.Recebivel, int64, error)
	ListPagamentos(ctx context.Context, pag PageParams) ([]domain.Pagamento, int64, error)
	CreateRecebivel(ctx context.Context, r *domain.Recebivel) error
	CreatePagamento(ctx context.Context, p *domain.Pagamento) error
	MarkRecebivelPago(ctx context.Context, uuid string) error
	MarkPagamentoPago(ctx context.Context, uuid string) error
}

type CursoRepository interface {
	ListCursos(ctx context.Context, pag PageParams) ([]domain.Curso, int64, error)
	ListCursosComProgresso(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Curso, []domain.CursoProgresso, int64, error)
}

type ChamadoRepository interface {
	ListByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Chamado, int64, error)
	ListAdmin(ctx context.Context, pag PageParams) ([]domain.Chamado, int64, error)
	Create(ctx context.Context, c *domain.Chamado) error
}

type ProdutoRepository interface {
	GetByUUID(ctx context.Context, uuid string) (*domain.Produto, error)
	ListByFamilia(ctx context.Context, familia string, pag PageParams) ([]domain.Produto, int64, error)
	Create(ctx context.Context, p *domain.Produto) error
	Update(ctx context.Context, p *domain.Produto) error
	Delete(ctx context.Context, uuid string) error
}

type AlertaRepository interface {
	List(ctx context.Context, pag PageParams) ([]domain.Alerta, int64, error)
	Resolver(ctx context.Context, uuid string) error
}

type NotificacaoRepository interface {
	ListEnviadas(ctx context.Context, pag PageParams) ([]domain.Notificacao, int64, error)
	Enviar(ctx context.Context, n *domain.Notificacao) error
}

type UsuarioRepository interface {
	GetByUUID(ctx context.Context, uuid string) (*domain.Usuario, error)
	GetByEmail(ctx context.Context, email string) (*domain.Usuario, error)
	List(ctx context.Context, pag PageParams) ([]domain.Usuario, int64, error)
	Create(ctx context.Context, u *domain.Usuario) error
	Update(ctx context.Context, u *domain.Usuario) error
}


