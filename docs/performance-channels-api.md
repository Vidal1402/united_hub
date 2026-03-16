# performance_channels no backend — O que conferir

O backend expõe `performance_channels` em **duas** respostas. O front pode usar qualquer uma (ou as duas).

---

## 1. GET `/api/auth/me`

- **Autenticação:** Bearer JWT (cliente ou admin).
- **Resposta:** objeto do usuário. Para **role client**, o backend preenche `performance_channels` a partir do documento do cliente no MongoDB.

**Formato da resposta (exemplo para cliente):**

```json
{
  "name": "cliente@email.com",
  "email": "cliente@email.com",
  "role": "client",
  "cliente_uuid": "uuid-do-cliente",
  "can_producao": true,
  "can_performance": true,
  "performance_channels": {
    "meta_ads": { "gasto": 1000, "leads": 50, "conversoes": 10 },
    "google_ads": { "gasto": 500, "leads": 20, "conversoes": 5 }
  }
}
```

- **Onde ler no front:** na **raiz** da resposta: `response.performance_channels` (ou `user.performance_channels` se guardar o objeto como "user").

---

## 2. GET `/api/cliente/config/perfil` e GET `/api/cliente/me`

- **Rotas:** ambas usam o mesmo handler e retornam o **cliente completo**.
  - `GET /api/cliente/config/perfil`
  - `GET /api/cliente/me`
- **Autenticação:** Bearer JWT de **cliente** (o `cliente_uuid` das claims identifica o cliente).

**Formato da resposta:**

O corpo é o **objeto do cliente** na raiz (não dentro de `data`/`client`/`cliente`):

```json
{
  "uuid": "uuid-do-cliente",
  "nome": "Nome do Cliente",
  "email": "cliente@email.com",
  "segmento": "",
  "plano": "",
  "status": "ativo",
  "cidade": "",
  "owner_uuid": "",
  "performance_channels": {
    "meta_ads": { "gasto": 1000, "leads": 50, "conversoes": 10 },
    "google_ads": { "gasto": 500, "leads": 20, "conversoes": 5 }
  },
  "created_at": "...",
  "updated_at": "..."
}
```

- **Onde ler no front:** na **raiz** da resposta: `response.performance_channels`. Não está dentro de `data`, `client` ou `cliente`; a resposta **é** o cliente.

---

## Resumo para o front

| Endpoint | Onde está `performance_channels` |
|----------|-----------------------------------|
| `GET /api/auth/me` | Na raiz do objeto usuário: `user.performance_channels` |
| `GET /api/cliente/config/perfil` | Na raiz do objeto cliente (o body é o cliente): `cliente.performance_channels` |
| `GET /api/cliente/me` | Idem: na raiz do objeto cliente: `cliente.performance_channels` |

Se os números não aparecem na aba Performance:

1. Confirmar que o front chama um desses endpoints após o login (ou ao abrir a aba).
2. Confirmar que lê o campo na raiz (não em `data.client` ou similar, a menos que o front encapsule a resposta).
3. Se o backend retornar `performance_channels: {}` ou omitir o campo, o admin ainda não salvou dados de canais para esse cliente no painel (PUT `/api/admin/clientes/:id` com `performance_channels` no body).

---

## Conferência no backend (já implementado)

- **AuthService.Me()** (`internal/service/auth_service.go`): para `role == client` e `cliente_uuid` preenchido, busca o documento do cliente e atribui `info.PerformanceChannels = cliente.PerformanceChannels` (e inicializa como `{}` se nil).
- **ClienteService.GetPerfil()** (`internal/service/cliente_service.go`): retorna o `*domain.Cliente` completo; o `domain.Cliente` tem o campo `PerformanceChannels` com tag `json:"performance_channels,omitempty"`.
- **Handler GetPerfil** (`internal/http/handlers/cliente/handler.go`): faz `response.JSON(w, http.StatusOK, result)` com o cliente; portanto a resposta já inclui `performance_channels` na raiz.

Nenhuma alteração adicional é necessária no backend para que os dois pontos de consumo acima funcionem.

---

## Frontend — aba Performance (referência)

**Arquivo:** `src/pages/united-dashboard.jsx` (componente `PerformancePage`).

**Endpoints usados:**

1. **GET /api/cliente/config/perfil** — corpo = objeto do cliente; leitura na raiz: `response.performance_channels`.
2. **GET /api/auth/me** — corpo = objeto do usuário; leitura na raiz: `response.performance_channels`.

As duas respostas são obtidas em paralelo (`Promise.all`). O valor usado é o primeiro não vazio entre: `extractPc(perfilRes)`, `extractPc(authMeRes)` e `extractPc(user)` (do AuthContext). A função `extractPc(obj)` lê `obj.performance_channels` (e variantes em camelCase/PascalCase e caminhos aninhados como fallback).

Os números exibidos (KPIs e “Dados por canal”) vêm desse `performance_channels` normalizado, com chaves em snake_case (`meta_ads`, `google_ads`, `organico`, `outros`) e campos por canal: `gasto`, `leads`, `conversoes`.
