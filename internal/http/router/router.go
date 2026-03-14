package router

import (
	"net/http"
	"os"
	"strings"
	"time"

	"backend_united_hub/internal/auth"
	adminhandler "backend_united_hub/internal/http/handlers/admin"
	authhandler "backend_united_hub/internal/http/handlers/auth"
	clientehandler "backend_united_hub/internal/http/handlers/cliente"
	"backend_united_hub/internal/http/handlers/health"
	"backend_united_hub/internal/http/middleware"
	"backend_united_hub/internal/repository"
	"backend_united_hub/internal/service"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// origens CORS permitidas: dev (localhost) + produção (env CORS_ORIGINS ou vazio)
func corsOrigins() []string {
	base := []string{
		"http://localhost:5173",
		"http://localhost:5174",
		"http://127.0.0.1:5173",
		"http://127.0.0.1:5174",
	}
	if s := os.Getenv("CORS_ORIGINS"); s != "" {
		// CORS_ORIGINS pode ser "https://meu-front.onrender.com" ou várias separadas por vírgula
		for _, o := range strings.Split(s, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				base = append(base, o)
			}
		}
	}
	return base
}

type Deps struct {
	JWTSecret      string
	RequestTimeout time.Duration
	DB             *mongo.Database
	Redis          *redis.Client
	Validator      *validator.Validate
}

func New(d Deps) http.Handler {
	r := chi.NewRouter()

	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(middleware.Recoverer())
	r.Use(middleware.WithTimeout(d.RequestTimeout))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           86400,
	}))
	// OPTIONS (preflight) responde 204 na raiz, sem exigir token — antes de qualquer rota protegida
	r.Use(middleware.PreflightNoAuth)

	h := health.New()
	r.Get("/healthz", h.Get)

	// Repositórios
	cliRepo := repository.NewMongoClienteRepository(d.DB)
	colabRepo := repository.NewMongoColaboradorRepository(d.DB)
	kanbanRepo := repository.NewMongoKanbanRepository(d.DB)
	relatorioRepo := repository.NewMongoRelatorioRepository(d.DB)
	materialRepo := repository.NewMongoMaterialRepository(d.DB)
	reuniaoRepo := repository.NewMongoReuniaoRepository(d.DB)
	financeiroRepo := repository.NewMongoFinanceiroRepository(d.DB)
	cursoRepo := repository.NewMongoCursoRepository(d.DB)
	chamadoRepo := repository.NewMongoChamadoRepository(d.DB)
	produtoRepo := repository.NewMongoProdutoRepository(d.DB)
	alertaRepo := repository.NewMongoAlertaRepository(d.DB)
	notificacaoRepo := repository.NewMongoNotificacaoRepository(d.DB)
	usuarioRepo := repository.NewMongoUsuarioRepository(d.DB)

	// Services
	clienteSvc := service.NewClienteService(
		cliRepo,
		kanbanRepo,
		relatorioRepo,
		materialRepo,
		reuniaoRepo,
		financeiroRepo,
		cursoRepo,
		chamadoRepo,
		alertaRepo,
		notificacaoRepo,
		d.Redis,
	)
	adminSvc := service.NewAdminService(
		cliRepo,
		colabRepo,
		financeiroRepo,
		produtoRepo,
		alertaRepo,
		notificacaoRepo,
		relatorioRepo,
		materialRepo,
		reuniaoRepo,
		chamadoRepo,
		kanbanRepo,
		usuarioRepo,
	)

	authSvc := service.NewAuthService(usuarioRepo, cliRepo, d.JWTSecret)

	// Handlers
	clienteH := clientehandler.New(clienteSvc, d.Validator)
	adminH := adminhandler.New(adminSvc, d.Validator)
	authH := authhandler.New(authSvc, d.Validator)

	r.Route("/api", func(api chi.Router) {
		// auth
		api.Post("/auth/login", authH.Login)
		api.With(middleware.RequireJWT(d.JWTSecret)).Get("/auth/me", authH.Me)

		api.Route("/cliente", func(cr chi.Router) {
			cr.Use(middleware.SkipAuthForOPTIONS)
			cr.Use(middleware.RequireJWT(d.JWTSecret))
			cr.Use(middleware.RequireRole(auth.RoleClient))

			cr.Get("/producao", clienteH.GetProducao)
			cr.Post("/producao/solicitacoes", clienteH.CreateSolicitacao)
			cr.Get("/dashboard/chart", clienteH.GetDashboardChart)
			cr.Get("/dashboard/funnel", clienteH.GetDashboardFunnel)
			cr.Get("/dashboard/kpis", clienteH.GetDashboardKPIs)
			cr.Get("/dashboard/score", clienteH.GetDashboardScore)

			cr.Get("/relatorios", clienteH.ListRelatorios)

			cr.Get("/materiais/pastas", clienteH.ListMateriaisPastas)
			cr.Get("/materiais/arquivos", clienteH.ListMateriaisArquivos)
			cr.Post("/materiais/upload", clienteH.UploadMaterial)

			cr.Get("/reunioes/proximas", clienteH.ListReunioesProximas)
			cr.Get("/reunioes/historico", clienteH.ListReunioesHistorico)

			cr.Get("/financeiro/faturas", clienteH.ListFaturas)
			cr.Get("/financeiro/plano", clienteH.GetPlanoFinanceiro)

			cr.Get("/academy/cursos", clienteH.ListCursos)

			cr.Get("/suporte/chamados", clienteH.ListChamados)
			cr.Post("/suporte/chamados", clienteH.CreateChamado)
			cr.Get("/suporte/faq", clienteH.ListFAQ)

			cr.Get("/config/perfil", clienteH.GetPerfil)
			cr.Put("/config/perfil", clienteH.UpdatePerfil)

			cr.Get("/config/usuarios", clienteH.ListUsuarios)

			cr.Get("/config/notificacoes", clienteH.GetNotificacoesConfig)
			cr.Put("/config/notificacoes", clienteH.UpdateNotificacoesConfig)

			cr.Get("/config/integracoes", clienteH.ListIntegracoes)
			cr.Post("/config/integracoes/{id}/conectar", clienteH.ConnectIntegracao)
		})

		api.Route("/admin", func(ar chi.Router) {
			ar.Use(middleware.SkipAuthForOPTIONS)
			ar.Use(middleware.RequireJWT(d.JWTSecret))
			ar.Use(middleware.RequireRole(auth.RoleAdmin))

			ar.Get("/overview", adminH.GetOverview)
			ar.Get("/overview/mrr-mensal", adminH.GetMRRMensal)

			ar.Get("/clientes", adminH.ListClientes)
			ar.Post("/clientes", adminH.CreateCliente)
			ar.Get("/clientes/{id}", adminH.GetCliente)
			ar.Put("/clientes/{id}", adminH.UpdateCliente)
			ar.Put("/clientes/{id}/desativar", adminH.DesativarCliente)

			ar.Get("/colaboradores", adminH.ListColaboradores)
			ar.Post("/colaboradores", adminH.CreateColaborador)
			ar.Get("/colaboradores/{id}", adminH.GetColaborador)
			ar.Put("/colaboradores/{id}", adminH.UpdateColaborador)

			ar.Get("/financeiro/receber", adminH.ListReceber)
			ar.Get("/financeiro/pagar", adminH.ListPagar)
			ar.Post("/financeiro/pagar", adminH.CreatePagar)
			ar.Put("/financeiro/pagar/{id}/marcar-pago", adminH.MarcarPagamentoPago)
			ar.Post("/financeiro/lancamento", adminH.CreateLancamento)
			ar.Put("/financeiro/receber/{id}/marcar-pago", adminH.MarcarRecebivelPago)

			ar.Get("/produtos/{familia}", adminH.ListProdutosPorFamilia)
			ar.Post("/produtos/{familia}", adminH.CreateProduto)
			ar.Put("/produtos/{familia}/{id}", adminH.UpdateProduto)
			ar.Delete("/produtos/{familia}/{id}", adminH.DeleteProduto)

			ar.Get("/alertas", adminH.ListAlertas)
			ar.Put("/alertas/{id}/resolver", adminH.ResolverAlerta)

			ar.Get("/notificacoes/enviadas", adminH.ListNotificacoesEnviadas)
			ar.Post("/notificacoes/enviar", adminH.EnviarNotificacao)

			ar.Get("/relatorios", adminH.ListRelatoriosAdmin)
			ar.Post("/relatorios", adminH.CreateRelatorio)

			ar.Post("/materiais/pastas", adminH.CreateMaterialPasta)
			ar.Options("/materiais/pastas/{id}", adminH.OptionsNoContent)
			ar.Patch("/materiais/pastas/{id}", adminH.UpdateMaterialPasta)
			ar.Post("/materiais/arquivos", adminH.CreateMaterialArquivo)
			ar.Options("/materiais/arquivos/{id}", adminH.OptionsNoContent)
			ar.Patch("/materiais/arquivos/{id}", adminH.UpdateMaterialArquivo)
			ar.Get("/materiais/pastas", adminH.ListPastasAdmin)
			ar.Get("/materiais/arquivos", adminH.ListArquivosAdmin)

			ar.Get("/reunioes", adminH.ListReunioesAdmin)
			ar.Post("/reunioes", adminH.CreateReuniao)

			ar.Get("/chamados", adminH.ListChamadosAdmin)
			ar.Post("/chamados", adminH.CreateChamado)

			ar.Get("/producao", adminH.GetProducaoAdmin)
			ar.Post("/producao/cards", adminH.CreateProducaoCard)
			ar.Patch("/producao/cards/{id}", adminH.UpdateProducaoCard)
			ar.Post("/producao/cards/{id}/comments", adminH.AddProducaoCardComment)

			ar.Get("/comercial", adminH.GetComercial)

			ar.Post("/usuarios", adminH.CreateUsuario)
			ar.Get("/usuarios", adminH.ListUsuarios)
			ar.Put("/usuarios/{id}", adminH.UpdateUsuario)
		})
	})

	return r
}
