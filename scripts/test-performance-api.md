# Como testar se os números (performance_channels) estão vindo da API

Use estes passos para ver se o backend está retornando `performance_channels` antes de debugar o front.

## 1. Base URL

- Local: `http://localhost:8080`
- Produção: `https://united-hub-3a6p.onrender.com` (ou a URL do seu backend)

Defina a variável (PowerShell):

```powershell
$BASE = "http://localhost:8080"
# ou: $BASE = "https://united-hub-3a6p.onrender.com"
```

## 2. Login com um usuário cliente

O usuário precisa ter `role: "client"` e `cliente_uuid` preenchido (UUID de um cliente que exista na collection `clientes`).

**PowerShell:**

```powershell
$body = '{"email":"SEU_EMAIL_CLIENTE","password":"SUA_SENHA"}' | ConvertTo-Json
$login = Invoke-RestMethod -Uri "$BASE/api/auth/login" -Method Post -Body $body -ContentType "application/json"
$token = $login.token
Write-Host "Token obtido (primeiros 50 chars): $($token.Substring(0, [Math]::Min(50, $token.Length)))..."
```

**curl (Git Bash / WSL):**

```bash
BASE="http://localhost:8080"
LOGIN=$(curl -s -X POST "$BASE/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"SEU_EMAIL_CLIENTE","password":"SUA_SENHA"}')
echo "$LOGIN"
TOKEN=$(echo "$LOGIN" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
echo "Token: ${TOKEN:0:50}..."
```

## 3. Chamar GET /api/auth/me (com token)

**PowerShell:**

```powershell
$headers = @{ Authorization = "Bearer $token" }
$me = Invoke-RestMethod -Uri "$BASE/api/auth/me" -Method Get -Headers $headers
$me | ConvertTo-Json -Depth 5
# Ver só performance_channels:
$me.performance_channels | ConvertTo-Json -Depth 5
```

**curl:**

```bash
curl -s -X GET "$BASE/api/auth/me" \
  -H "Authorization: Bearer $TOKEN" | jq .
# Ver só performance_channels:
curl -s -X GET "$BASE/api/auth/me" \
  -H "Authorization: Bearer $TOKEN" | jq '.performance_channels'
```

## 4. Chamar GET /api/cliente/me (alternativa)

Retorna o cliente completo (incluindo `performance_channels`).

**PowerShell:**

```powershell
$meCliente = Invoke-RestMethod -Uri "$BASE/api/cliente/me" -Method Get -Headers $headers
$meCliente.performance_channels | ConvertTo-Json -Depth 5
```

**curl:**

```bash
curl -s -X GET "$BASE/api/cliente/me" \
  -H "Authorization: Bearer $TOKEN" | jq '.performance_channels'
```

## 5. O que esperar

- Se **não** houver dados salvos no cliente: `performance_channels` pode vir como `{}` ou `null`.
- Se o **admin já salvou** canais para esse cliente: deve vir algo como:

```json
{
  "meta_ads": { "gasto": 1000, "leads": 50, "conversoes": 10 },
  "google_ads": { "gasto": 500, "leads": 20, "conversoes": 5 }
}
```

Se vier `{}` ou não vier o campo, confira:

1. O `cliente_uuid` do usuário (no JWT / na tabela de usuários) é o mesmo do cliente que o admin editou?
2. No MongoDB, o documento em `clientes` com esse `uuid` tem o campo `performance_channels` preenchido? (ex.: no Mongo Shell ou Compass: `db.clientes.findOne({ uuid: "UUID_DO_CLIENTE" })`).

## 6. Teste rápido em uma linha (PowerShell, após $BASE e $token)

```powershell
(Invoke-RestMethod -Uri "$BASE/api/auth/me" -Headers @{Authorization="Bearer $token"}).performance_channels
```

Se aparecer um objeto com `meta_ads`, `google_ads`, etc., a API está devolvendo os números; aí o problema é só no front (onde lê ou onde imprime).
