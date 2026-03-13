## Esquema de Dados – MongoDB (United Hub)

Este documento descreve **todas as collections** que o backend utiliza, com:

- Campos (sem faltar nenhum)
- Tipos lógicos esperados
- Chaves primárias lógicas (PK)
- Chaves estrangeiras lógicas (FK)
- Índices criados em `migrations.UpMongo`
- Dependências entre entidades

> Observação: o MongoDB **não aplica chaves estrangeiras nativamente**. As FKs abaixo são relações lógicas, garantidas pela aplicação e por índices quando faz sentido.

---

### Convenções gerais

- **UUID**: sempre string no formato UUID v4.
- **Datas**: armazenadas como `datetime` (`time.Time` no Go, `Date` no Mongo).
- **Booleans**: `true`/`false`.
- **Inteiros**:
  - `int`: contagens, ordenações, progresso (0–100).
  - `int64`: valores monetários em centavos (`*_centavos`, `tamanho` de arquivo).
- **Chaves**:
  - PK lógica quase sempre é o campo `uuid` (ou `slug` / par composto como indicado).
  - Índices únicos garantem unicidade onde necessário.

---

## 1. Collection `usuarios`

Entidade de **autenticação** e **permissões**. Usada por `/api/auth/*` e `/api/admin/usuarios*`.

**Campos**

- `uuid` (string, UUID v4) – identificador do usuário.
- `cliente_uuid` (string, UUID v4, opcional) – tenant associado (para `role = "client"`). Pode ser vazio para admin global.
- `email` (string) – login do usuário.
- `senha_hash` (string) – hash bcrypt da senha.
- `role` (string) – `"client"` ou `"admin"`.
- `can_producao` (bool) – habilita rota `/api/cliente/producao`.
- `can_performance` (bool) – habilita rotas `/api/cliente/dashboard/*`.
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique index).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid` (quando usuário é client).

**Índices**

- `uuid` (unique).
- `email` (unique).

---

## 2. Collection `clientes`

Representa o **cliente/tenant** do hub.

**Campos**

- `uuid` (string, UUID v4).
- `nome` (string).
- `email` (string).
- `segmento` (string) – ex.: nicho/vertical.
- `plano` (string) – ex.: `"Starter"`, `"Growth"`, `"Pro"`, `"Scale"`.
- `status` (string) – ex.: `"ativo"`, `"inativo"`.
- `cidade` (string).
- `owner_uuid` (string, UUID v4) – responsável (FK para `colaboradores.uuid`).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique index).

**FKs lógicas**

- `owner_uuid` → `colaboradores.uuid`.

**Índices**

- `uuid` (unique).
- `email` (unique).

---

## 3. Collection `colaboradores`

Colaboradores internos (ex.: time da agência) que podem ser donos de clientes, reuniões, etc.

**Campos**

- `uuid` (string, UUID v4).
- `nome` (string).
- `email` (string).
- `cargo` (string).
- `role` (string) – ex.: `"admin"`, `"user"` (interno, não é o mesmo `role` de `usuarios`).
- `status` (string).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**Índices**

- `uuid` (unique).
- `email` (unique).

**Usado por**

- `clientes.owner_uuid`.
- `reunioes.owner_uuid`.
- `relatorios.owner_uuid`.

---

## 4. Collection `kanban_columns`

Colunas do quadro de produção por cliente.

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `key` (string) – ex.: `"sol"`, `"pend"`, `"prod"`.
- `label` (string) – ex.: `"Solicitações"`, `"Pendente"`.
- `ordenacao` (int) – ordem da coluna.
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.

**Índices**

- `cliente_uuid`.
- `uuid` (unique).

**Usado por**

- `kanban_cards.column_uuid`.

---

## 5. Collection `kanban_cards`

Cards do Kanban de produção do cliente.

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `column_uuid` (string, UUID v4).
- `titulo` (string).
- `tipo` (string) – categoria/label do card.
- `owner_uuid` (string, UUID v4) – responsável (colaborador).
- `prazo` (datetime).
- `prioridade` (string) – `"Alta"`, `"Média"`, `"Baixa"`.
- `comentarios` (int) – contagem de comentários.
- `arquivos` (int) – contagem de anexos.
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.
- `column_uuid` → `kanban_columns.uuid`.
- `owner_uuid` → `colaboradores.uuid`.

**Índices**

- Composto `{cliente_uuid, column_uuid}`.
- `uuid` (unique).

---

## 6. Collection `relatorios`

Relatórios entregues ao cliente (mensais, de campanha, etc.).

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `titulo` (string).
- `tipo` (string) – ex.: `"Mensal"`, `"Campanha"`.
- `periodo` (string) – ex.: `"2026-01"`.
- `owner_uuid` (string, UUID v4) – responsável (colaborador).
- `data` (datetime) – data de emissão.
- `paginas` (int).
- `file_url` (string) – URL do PDF/arquivo.
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.
- `owner_uuid` → `colaboradores.uuid`.

**Índices**

- Composto `{cliente_uuid, periodo}`.
- `uuid` (unique).

---

## 7. Collection `materiais_pastas`

Pastas de materiais (drive do cliente).

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `nome` (string).
- `icone` (string) – nome do ícone/emoji.
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.

**Índices**

- `cliente_uuid`.
- `uuid` (unique).

---

## 8. Collection `materiais_arquivos`

Arquivos de materiais, dentro das pastas.

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `pasta_uuid` (string, UUID v4).
- `nome` (string).
- `extensao` (string) – ex.: `"pdf"`, `"png"`.
- `tamanho` (int64) – em bytes.
- `data` (datetime) – data do material (upload ou referência).
- `url` (string) – URL do arquivo.
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.
- `pasta_uuid` → `materiais_pastas.uuid`.

**Índices**

- Composto `{cliente_uuid, pasta_uuid}`.
- `uuid` (unique).

---

## 9. Collection `reunioes`

Reuniões com o cliente (passadas e futuras).

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `titulo` (string).
- `data_hora` (datetime).
- `via` (string) – ex.: `"Zoom"`, `"Meet"`, `"Presencial"`.
- `owner_uuid` (string, UUID v4) – responsável (colaborador).
- `pauta` (array de string).
- `status` (string) – ex.: `"futura"`, `"historico"`.
- `duracao_min` (int).
- `tem_gravacao` (bool).
- `tem_ata` (bool).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.
- `owner_uuid` → `colaboradores.uuid`.

**Índices**

- Composto `{cliente_uuid, data_hora}`.
- `uuid` (unique).

---

## 10. Collection `faturas`

Faturas do cliente (financeiro).

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `periodo` (string).
- `valor_centavos` (int64) – valor total em centavos.
- `vencimento` (datetime).
- `status` (string) – `"Pago"`, `"Pendente"`, `"Vencido"`.
- `data_pagamento` (datetime, opcional).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.

**Índices**

- Composto `{cliente_uuid, status}`.
- `vencimento`.
- `uuid` (unique).

---

## 11. Collection `recebiveis`

Recebíveis no ADM (contas a receber).

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `descricao` (string).
- `valor_centavos` (int64).
- `vencimento` (datetime).
- `status` (string).
- `plano` (string).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.

**Índices**

- Composto `{cliente_uuid, status}`.
- `uuid` (unique).

---

## 12. Collection `pagamentos`

Pagamentos no ADM (contas a pagar).

**Campos**

- `uuid` (string, UUID v4).
- `descricao` (string).
- `valor_centavos` (int64).
- `vencimento` (datetime).
- `status` (string) – `"Pago"`, `"Pendente"`.
- `categoria` (string).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**Índices**

- `status`.
- `uuid` (unique).

---

## 13. Collection `cursos`

Cursos da academy.

**Campos**

- `slug` (string).
- `titulo` (string).
- `categoria` (string).
- `formato` (string).
- `duracao` (string).
- `nivel` (string).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `slug` (unique).

**Índices**

- `slug` (unique).

---

## 14. Collection `cursos_progresso`

Progresso do cliente nos cursos.

**Campos**

- `cliente_uuid` (string, UUID v4).
- `curso_slug` (string) – FK para `cursos.slug`.
- `concluido` (bool).
- `progresso` (int) – 0–100.
- `updated_at` (datetime).

**PK lógica (composta)**

- `(cliente_uuid, curso_slug)` – índice único.

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.
- `curso_slug` → `cursos.slug`.

**Índices**

- Composto `{cliente_uuid, curso_slug}` (unique).

---

## 15. Collection `chamados`

Chamados de suporte do cliente.

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4).
- `categoria` (string).
- `titulo` (string).
- `status` (string).
- `criado_em` (datetime).
- `atualizado_em` (datetime).
- `descricao` (string).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid`.

**Índices**

- Composto `{cliente_uuid, status}`.
- `uuid` (unique).

---

## 16. Collection `produtos`

Produtos/planos configuráveis no ADM.

**Campos**

- `uuid` (string, UUID v4).
- `familia` (string) – ex.: `"marketing"`, `"food"`, `"ia"`, `"crm"`.
- `slug` (string).
- `nome` (string).
- `preco_centavos` (int64).
- `descricao` (string).
- `features` (array de string).
- `created_at` (datetime).
- `updated_at` (datetime).

**PK lógica**

- `uuid` (unique).

**Índice de unicidade funcional**

- `(familia, slug)` (unique).

**Índices**

- Composto `{familia, slug}` (unique).

---

## 17. Collection `alertas`

Alertas gerados para clientes ou globais.

**Campos**

- `uuid` (string, UUID v4).
- `cliente_uuid` (string, UUID v4, opcional).
- `titulo` (string).
- `tipo` (string).
- `prioridade` (string).
- `target` (string) – ex.: `"cliente"`, `"admin"`, `"todos"`.
- `status` (string) – `"Ativo"`, `"Resolvido"`.
- `created_at` (datetime).
- `resolved_at` (datetime, opcional).

**PK lógica**

- `uuid` (unique).

**FKs lógicas**

- `cliente_uuid` → `clientes.uuid` (quando preenchido).

**Índices**

- Composto `{cliente_uuid, status}`.
- `created_at` (desc).

---

## 18. Collection `notificacoes`

Notificações enviadas pelo sistema.

**Campos**

- `uuid` (string, UUID v4).
- `titulo` (string).
- `target` (string) – identificador do alvo (cliente específico ou `"todos"`).\n+- `canal` (string) – `"Email"`, `"Plataforma"`, `"WhatsApp"`, etc.\n+- `tipo` (string) – ex.: `"sistema"`, `"financeiro"`, etc.\n+- `lidas` (int) – contagem de leituras.\n+- `criado_em` (datetime).\n+- `conteudo` (string) – texto da notificação.\n\n**PK lógica**\n\n- `uuid` (unique).\n\n**Índices**\n\n- Composto `{tipo, created_at}` (na coleção, o campo temporal indexado é `created_at`; o campo lógico é `criado_em` no domínio).\n\n---\n\n## Relações principais (resumo)\n\n- **Usuario → Cliente**\n  - `usuarios.cliente_uuid` → `clientes.uuid` (para usuários de role `client`).\n\n- **Cliente ↔ Colaborador**\n  - `clientes.owner_uuid` → `colaboradores.uuid`.\n\n- **Kanban**\n  - `kanban_columns.cliente_uuid` → `clientes.uuid`.\n  - `kanban_cards.cliente_uuid` → `clientes.uuid`.\n  - `kanban_cards.column_uuid` → `kanban_columns.uuid`.\n  - `kanban_cards.owner_uuid` → `colaboradores.uuid`.\n\n- **Reuniões/Relatórios**\n  - `reunioes.cliente_uuid` → `clientes.uuid`.\n  - `reunioes.owner_uuid` → `colaboradores.uuid`.\n  - `relatorios.cliente_uuid` → `clientes.uuid`.\n  - `relatorios.owner_uuid` → `colaboradores.uuid`.\n\n- **Materiais**\n  - `materiais_pastas.cliente_uuid` → `clientes.uuid`.\n  - `materiais_arquivos.cliente_uuid` → `clientes.uuid`.\n  - `materiais_arquivos.pasta_uuid` → `materiais_pastas.uuid`.\n\n- **Financeiro**\n  - `faturas.cliente_uuid` → `clientes.uuid`.\n  - `recebiveis.cliente_uuid` → `clientes.uuid`.\n\n- **Cursos**\n  - `cursos_progresso.cliente_uuid` → `clientes.uuid`.\n  - `cursos_progresso.curso_slug` → `cursos.slug`.\n\n- **Suporte/Alertas/Notificações**\n  - `chamados.cliente_uuid` → `clientes.uuid`.\n  - `alertas.cliente_uuid` → `clientes.uuid` (quando presente).\n  - `notificacoes.target` pode apontar para cliente ou ser `"todos"`.\n\nCom essas collections, campos e índices, o backend consegue operar **100% funcional** conforme implementado: autenticação JWT, controle de permissões (`role`, `can_producao`, `can_performance`), rotas de cliente/admin, dashboard, produção, materiais, reuniões, financeiro, academy, suporte, alertas e notificações. \n*** End Patch】} ***!
