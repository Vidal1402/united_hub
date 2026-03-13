# backend_united_hub

Backend Go (Clean Architecture) para **United — Growth Hub** (área do cliente) e **Admin Panel**. API REST com MongoDB, Redis (cache opcional), JWT e controle de permissões por usuário.

---

## Índice

- [Stack e dependências](#stack-e-dependências)
- [Variáveis de ambiente](#variáveis-de-ambiente)
- [Como rodar](#como-rodar)
- [Rotas da API](#rotas-da-api)
- [Autenticação e permissões](#autenticação-e-permissões)
- [Paginação](#paginação)
- [DTOs e campos](#dtos-e-campos)
- [Collections (MongoDB)](#collections-mongodb)
- [Deploy (Render)](#deploy-render)

---

## Stack e dependências

| Pacote | Uso |
|--------|-----|
| `github.com/go-chi/chi/v5` | Router HTTP |
| `github.com/go-chi/cors` | CORS |
| `github.com/go-playground/validator/v10` | Validação de body/query |
| `github.com/golang-jwt/jwt/v5` | JWT (login + claims) |
| `github.com/google/uuid` | UUIDs |
| `github.com/joho/godotenv` | `.env` local |
| `github.com/redis/go-redis/v9` | Cache (opcional) |
| `go.mongodb.org/mongo-driver` | MongoDB |
| `golang.org/x/crypto` | bcrypt (senha) |

- **Go**: 1.26+

---

## Variáveis de ambiente

| Variável | Obrigatório | Default | Descrição |
|----------|-------------|---------|-----------|
| `PORT` | Não | `8080` | Porta HTTP (Render define automaticamente) |
| `MONGODB_URI` | Sim* | — | URI do MongoDB (Atlas ou local). Fallback: `DATABASE_URL` |
| `MONGODB_DB` | Não | `united_hub` | Nome do banco |
| `REDIS_ADDR` | Não | `localhost:6379` | Endereço Redis (se falhar, app sobe sem cache) |
| `REDIS_PASSWORD` | Não | — | Senha Redis |
| `REDIS_DB` | Não | `0` | Índice do DB Redis |
| `JWT_SECRET` | Sim | — | Chave para assinar/validar JWT |
| `REQUEST_TIMEOUT_MS` | Não | `8000` | Timeout global dos requests (ms) |
| `UPLOAD_DIR` | Não | `./storage` | Diretório de uploads |

\* Sem `MONGODB_URI` (ou `DATABASE_URL`) e `JWT_SECRET` a aplicação não inicia.

---

## Como rodar

### Local

```powershell
cd c:\Users\luisg\backend_united_hub
Copy-Item .env.example .env
# Edite .env com MONGODB_URI e JWT_SECRET (Redis opcional)
go mod tidy
go test ./...
go run .\cmd\api\main.go
```

- Health: **GET** `http://localhost:8080/healthz` → `{"status":"ok"}`

### Docker (Redis local, opcional)

```powershell
docker run --name redis-united-hub -p 6379:6379 redis:7-alpine
```

---

## Rotas da API

Base URL: `http://localhost:8080` (ou a URL do deploy).

Todas as rotas `/api/*` (exceto `/api/auth/login`) que exigem autenticação usam:

- **Header**: `Authorization: Bearer <token>`

Respostas de erro comuns:

- `401 Unauthorized`: token ausente ou inválido
- `403 Forbidden`: role incorreta ou permissão insuficiente (ex.: `can_producao` / `can_performance`)
- `400 Bad Request`: body inválido ou falha de validação

---

### Públicas (sem auth)

| Método | Path | Descrição |
|--------|------|-----------|
| GET | `/healthz` | Health check. Resposta: `{"status":"ok"}` |

---

### Auth

| Método | Path | Auth | Body | Resposta |
|--------|------|------|------|----------|
| POST | `/api/auth/login` | Não | `LoginInput` (email, password) | `LoginResponse` (token + user) |
| GET | `/api/auth/me` | JWT (qualquer role) | — | `UserInfo` |

**LoginInput**

```json
{
  "email": "string (required, email)",
  "password": "string (required)"
}
```

**LoginResponse**

```json
{
  "token": "string (JWT)",
  "user": {
    "name": "string",
    "email": "string",
    "role": "client | admin",
    "cliente_uuid": "string (omitempty para admin)",
    "can_producao": true,
    "can_performance": false
  }
}
```

**UserInfo** (usado em `login.user` e `GET /api/auth/me`)

```json
{
  "name": "string",
  "email": "string",
  "role": "string",
  "cliente_uuid": "string (omitempty)",
  "can_producao": true,
  "can_performance": false
}
```

---

### Cliente (`/api/cliente/*`)

Requer **JWT** com **role = client**. Algumas rotas checam permissões:

- **Produção**: exige `can_producao == true`
- **Dashboard** (chart, funnel, kpis, score): exige `can_performance == true`

| Método | Path | Permissão | Query | Resposta |
|--------|------|------------|-------|----------|
| GET | `/api/cliente/producao` | can_producao | — | Objeto com `columns` e `cards` (kanban) |
| GET | `/api/cliente/dashboard/chart` | can_performance | `period` (opcional) | Objeto (ex.: points) |
| GET | `/api/cliente/dashboard/funnel` | can_performance | — | Objeto (ex.: stages) |
| GET | `/api/cliente/dashboard/kpis` | can_performance | — | Objeto (ex.: kpis) |
| GET | `/api/cliente/dashboard/score` | can_performance | — | Objeto (ex.: score) |
| GET | `/api/cliente/relatorios` | — | limit, offset | `Page` (items + total + limit + offset) |
| GET | `/api/cliente/materiais/pastas` | — | — | Lista de pastas |
| GET | `/api/cliente/materiais/arquivos` | — | limit, offset | `Page` de arquivos |
| POST | `/api/cliente/materiais/upload` | — | — | Body: `UploadMaterialInput` → 201 |
| GET | `/api/cliente/reunioes/proximas` | — | — | Lista de reuniões |
| GET | `/api/cliente/reunioes/historico` | — | limit, offset | `Page` de reuniões |
| GET | `/api/cliente/financeiro/faturas` | — | limit, offset | `Page` de faturas |
| GET | `/api/cliente/financeiro/plano` | — | — | Objeto plano |
| GET | `/api/cliente/academy/cursos` | — | — | Lista de cursos |
| GET | `/api/cliente/suporte/chamados` | — | limit, offset | `Page` de chamados |
| POST | `/api/cliente/suporte/chamados` | — | — | Body: `CreateChamadoInput` → 201 |
| GET | `/api/cliente/suporte/faq` | — | — | Lista FAQ |
| GET | `/api/cliente/config/perfil` | — | — | Perfil (cliente) |
| PUT | `/api/cliente/config/perfil` | — | — | Body: `UpdatePerfilInput` → 200 |
| GET | `/api/cliente/config/usuarios` | — | — | Lista usuários do tenant |
| GET | `/api/cliente/config/notificacoes` | — | — | Config de notificações |
| PUT | `/api/cliente/config/notificacoes` | — | — | Body: `UpdateNotificacoesConfigInput` → 200 |
| GET | `/api/cliente/config/integracoes` | — | — | Lista integrações |
| POST | `/api/cliente/config/integracoes/{id}/conectar` | — | — | id no path → 200 |

---

### Admin (`/api/admin/*`)

Requer **JWT** com **role = admin**.

| Método | Path | Query / Body | Resposta |
|--------|------|--------------|----------|
| GET | `/api/admin/overview` | — | Objeto overview |
| GET | `/api/admin/overview/mrr-mensal` | — | Objeto MRR |
| GET | `/api/admin/clientes` | limit, offset | `Page` de clientes |
| POST | `/api/admin/clientes` | Body: `ClienteInput` | 201 objeto cliente |
| GET | `/api/admin/clientes/{id}` | — | Cliente |
| PUT | `/api/admin/clientes/{id}` | Body: `ClienteInput` | Objeto cliente |
| PUT | `/api/admin/clientes/{id}/desativar` | — | 200 |
| GET | `/api/admin/colaboradores` | limit, offset | `Page` de colaboradores |
| POST | `/api/admin/colaboradores` | Body (nome, email, cargo, role, status) | 201 |
| GET | `/api/admin/colaboradores/{id}` | — | Colaborador |
| PUT | `/api/admin/colaboradores/{id}` | Body | Objeto colaborador |
| GET | `/api/admin/financeiro/receber` | limit, offset | `Page` recebíveis |
| GET | `/api/admin/financeiro/pagar` | limit, offset | `Page` pagamentos |
| POST | `/api/admin/financeiro/lancamento` | Body | 201 |
| PUT | `/api/admin/financeiro/receber/{id}/marcar-pago` | — | 200 |
| GET | `/api/admin/produtos/{familia}` | limit, offset | `Page` produtos |
| POST | `/api/admin/produtos/{familia}` | Body | 201 |
| PUT | `/api/admin/produtos/{familia}/{id}` | Body | Objeto produto |
| DELETE | `/api/admin/produtos/{familia}/{id}` | — | 204 |
| GET | `/api/admin/alertas` | limit, offset | `Page` alertas |
| PUT | `/api/admin/alertas/{id}/resolver` | — | 200 |
| GET | `/api/admin/notificacoes/enviadas` | limit, offset | `Page` notificações |
| POST | `/api/admin/notificacoes/enviar` | Body | 201 |
| GET | `/api/admin/relatorios` | limit, offset | `Page` relatórios |
| GET | `/api/admin/comercial` | — | Objeto comercial |
| POST | `/api/admin/usuarios` | Body: `UsuarioCreateInput` | 201 `UsuarioOutput` |
| GET | `/api/admin/usuarios` | limit, offset | `Page<UsuarioOutput>` |
| PUT | `/api/admin/usuarios/{id}` | Body: `UsuarioUpdateInput` | `UsuarioOutput` |

---

## Autenticação e permissões

- **Login**: `POST /api/auth/login` com `email` e `password`. O backend busca o usuário na collection `usuarios`, compara senha com bcrypt e devolve um JWT com:
  - `role`: `client` ou `admin`
  - `cliente_id`: UUID do cliente (para client); admin pode ter vazio
  - `usuario_id`, `email`, `can_producao`, `can_performance`

- **Rotas cliente**:
  - `/api/cliente/producao` → exige `can_producao == true` (senão 403).
  - `/api/cliente/dashboard/*` → exige `can_performance == true` (senão 403).
  - Demais rotas cliente só exigem role `client`.

- **Rotas admin**: só exige role `admin`.

---

## Paginação

Rotas que listam recursos aceitam query params:

| Param | Tipo | Default (ex.) | Descrição |
|-------|------|----------------|-----------|
| `limit` | int | 20 | Máximo de itens (máx. 100 em vários handlers) |
| `offset` | int | 0 | Quantos itens pular |

Resposta paginada (DTO `Page<T>`):

```json
{
  "items": [ ... ],
  "total": 100,
  "limit": 20,
  "offset": 0
}
```

---

## DTOs e campos

### Auth

- **LoginInput**: `email` (required, email), `password` (required)
- **UserInfo**: `name`, `email`, `role`, `cliente_uuid`, `can_producao`, `can_performance`
- **LoginResponse**: `token`, `user` (UserInfo)

### Usuário (admin)

- **UsuarioOutput**: `uuid`, `cliente_uuid`, `email`, `role`, `can_producao`, `can_performance`
- **UsuarioCreateInput**: `cliente_uuid` (opcional, uuid4), `email` (required, email), `password` (required, min 6), `role` (required, oneof=client admin), `can_producao`, `can_performance`
- **UsuarioUpdateInput**: `can_producao` (optional bool), `can_performance` (optional bool)

### Cliente

- **ClienteInput** (create/update): `nome`, `email`, `segmento`, `plano`, `status`, `cidade`, `owner_uuid` (uuid4)
- **ClienteOutput**: idem + `uuid`, `created_at`, `updated_at`

### Outros inputs (resumo)

- **UploadMaterialInput**: `pasta_uuid`, `nome`, `extensao`, `tamanho`, `url`, `data` (opcional)
- **CreateChamadoInput**: `categoria`, `titulo`, `descricao`
- **UpdatePerfilInput**: `nome`, `email`, `cidade`
- **UpdateNotificacoesConfigInput**: `canal_email`, `canal_plataforma`, `canal_whatsapp`

### Paginação

- **Page&lt;T&gt;** (resposta): `items` (array), `total` (int64), `limit` (int), `offset` (int)

---

## Collections (MongoDB)

O backend usa **MongoDB**. As collections são criadas/ajustadas no startup via `migrations.UpMongo` (índices). Nenhum dado é inserido automaticamente; usuários/clientes iniciais devem ser criados via API ou manualmente.

| Collection | Chave lógica | Campos principais |
|------------|--------------|-------------------|
| **usuarios** | uuid (unique), email (unique) | uuid, cliente_uuid, email, senha_hash, role (client\|admin), can_producao, can_performance, created_at, updated_at |
| **clientes** | uuid, email (unique) | uuid, nome, email, segmento, plano, status, cidade, owner_uuid, created_at, updated_at |
| **colaboradores** | uuid, email (unique) | uuid, nome, email, cargo, role, status, created_at, updated_at |
| **kanban_columns** | uuid | uuid, cliente_uuid, key, label, ordenacao, created_at, updated_at |
| **kanban_cards** | uuid | uuid, cliente_uuid, column_uuid, titulo, tipo, owner_uuid, prazo, prioridade, comentarios, arquivos, created_at, updated_at |
| **relatorios** | uuid | uuid, cliente_uuid, titulo, tipo, periodo, owner_uuid, data, paginas, file_url, created_at, updated_at |
| **materiais_pastas** | uuid | uuid, cliente_uuid, nome, icone, created_at, updated_at |
| **materiais_arquivos** | uuid | uuid, cliente_uuid, pasta_uuid, nome, extensao, tamanho, data, url, created_at, updated_at |
| **reunioes** | uuid | uuid, cliente_uuid, titulo, data_hora, via, owner_uuid, pauta[], status, duracao_min, tem_gravacao, tem_ata, created_at, updated_at |
| **faturas** | uuid | uuid, cliente_uuid, periodo, valor_centavos, vencimento, status, data_pagamento?, created_at, updated_at |
| **recebiveis** | uuid | uuid, cliente_uuid, descricao, valor_centavos, vencimento, status, plano, created_at, updated_at |
| **pagamentos** | uuid | uuid, descricao, valor_centavos, vencimento, status, categoria, created_at, updated_at |
| **cursos** | slug (unique) | slug, titulo, categoria, formato, duracao, nivel, created_at, updated_at |
| **cursos_progresso** | (cliente_uuid, curso_slug) unique | cliente_uuid, curso_slug, concluido, progresso, updated_at |
| **chamados** | uuid | uuid, cliente_uuid, categoria, titulo, status, criado_em, atualizado_em, descricao |
| **produtos** | (familia, slug) unique | uuid, familia, slug, nome, preco_centavos, descricao, features[], created_at, updated_at |
| **alertas** | uuid | uuid, cliente_uuid?, titulo, tipo, prioridade, target, status, created_at, resolved_at? |
| **notificacoes** | uuid | uuid, titulo, target, canal, tipo, lidas, criado_em, conteudo |

---

## Deploy (Render)

- **Build Command**: `go build -o main ./cmd/api`
- **Start Command**: `./main`
- **Environment**: definir no painel do Render pelo menos `MONGODB_URI` e `JWT_SECRET`. `PORT` é definido pelo Render. Redis é opcional (se não configurar, a API sobe sem cache).

---

## Estrutura do projeto (resumo)

```
cmd/api/main.go          # Entrypoint: config, Mongo, Redis, migrations, router
internal/
  auth/                   # JWT: Claims, ParseToken, SignToken
  config/                 # Load env (MONGODB_URI, JWT_SECRET, etc.)
  cache/                  # Redis client
  db/                     # Mongo client
  domain/                 # Entidades (Usuario, Cliente, Kanban, etc.)
  http/
    dto/                  # DTOs de request/response (auth, usuario, cliente, paginação)
    handlers/
      auth/               # POST login, GET me
      cliente/            # Rotas /api/cliente/*
      admin/              # Rotas /api/admin/*
      health/             # GET /healthz
    middleware/           # RequireJWT, RequireRole, Recoverer, WithTimeout
    pagination/           # Parse limit/offset da query
    response/             # JSON, Error
    router/               # Montagem de rotas e injeção de deps
  migrations/             # UpMongo (índices em todas as collections)
  repository/             # Interfaces + implementações Mongo (por entidade)
  service/                # AuthService, ClienteService, AdminService
```
