# Postman — testar performance_channels

## Importar a collection

1. Abra o Postman.
2. **Import** → **Upload Files** (ou arraste o arquivo).
3. Selecione `performance-channels.postman_collection.json`.

## Variáveis

- **base_url**: URL do backend (ex.: `http://localhost:8080` ou `https://united-hub-3a6p.onrender.com`). Edite na collection (⋯ → Edit → Variables).
- **token**: preenchido automaticamente após rodar **1. Login (cliente)**.

## Passos

1. Ajuste **base_url** se precisar (ex.: produção).
2. Abra **1. Login (cliente)** e no **Body** troque `SEU_EMAIL_CLIENTE` e `SUA_SENHA` por um usuário **cliente** real.
3. **Send** em **1. Login (cliente)**. O token é salvo sozinho.
4. **Send** em **2. GET /api/auth/me** — na resposta, veja se existe **performance_channels** na raiz (ex.: `"performance_channels": { "meta_ads": { ... } }`).
5. **Send** em **3. GET /api/cliente/config/perfil** — idem: **performance_channels** na raiz.
6. (Opcional) **4. GET /api/cliente/me** — mesmo formato que o item 3.

Se em **todas** as respostas `performance_channels` vier vazio `{}` ou não existir, o backend está certo e o dado ainda não foi salvo para esse cliente (admin deve usar PUT `/api/admin/clientes/:id` com `performance_channels` no body). Se **alguma** rota vier com números e o front não mostrar, o problema é no front (rota ou onde lê o campo).
