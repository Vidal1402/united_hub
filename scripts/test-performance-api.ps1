# Testa se a API retorna performance_channels para o cliente logado.
# Uso: .\scripts\test-performance-api.ps1
# Ou com parâmetros: .\scripts\test-performance-api.ps1 -BaseUrl "http://localhost:8080" -Email "cliente@email.com" -Password "senha123"

param(
    [string]$BaseUrl = "http://localhost:8080",
    [string]$Email = "",
    [string]$Password = ""
)

if (-not $Email) { $Email = Read-Host "Email (usuário cliente)" }
if (-not $Password) { $Secure = Read-Host "Senha" -AsSecureString; $Password = [Runtime.InteropServices.Marshal]::PtrToStringAuto([Runtime.InteropServices.Marshal]::SecureStringToBSTR($Secure)) }

Write-Host "Login em $BaseUrl ..." -ForegroundColor Cyan
try {
    $loginBody = @{ email = $Email; password = $Password } | ConvertTo-Json
    $login = Invoke-RestMethod -Uri "$BaseUrl/api/auth/login" -Method Post -Body $loginBody -ContentType "application/json; charset=utf-8"
} catch {
    Write-Host "Erro no login: $_" -ForegroundColor Red
    exit 1
}

$token = $login.token
if (-not $token) {
    Write-Host "Resposta sem token. Resposta: $($login | ConvertTo-Json)" -ForegroundColor Red
    exit 1
}

Write-Host "Token obtido. Chamando GET /api/auth/me ..." -ForegroundColor Cyan
$headers = @{ Authorization = "Bearer $token" }
try {
    $me = Invoke-RestMethod -Uri "$BaseUrl/api/auth/me" -Method Get -Headers $headers
} catch {
    Write-Host "Erro em /api/auth/me: $_" -ForegroundColor Red
    exit 1
}

Write-Host "`n--- Resposta completa de /api/auth/me ---" -ForegroundColor Yellow
$me | ConvertTo-Json -Depth 6

Write-Host "`n--- performance_channels (números para a aba Performance) ---" -ForegroundColor Yellow
if ($me.performance_channels) {
    $me.performance_channels | ConvertTo-Json -Depth 6
    $count = ($me.performance_channels | Get-Member -MemberType NoteProperty).Count
    Write-Host "`nCanais encontrados: $count" -ForegroundColor Green
} else {
    Write-Host "{} ou null - nenhum dado de canais ainda. O admin precisa salvar em Salvar informações dos canais para esse cliente." -ForegroundColor Magenta
}

Write-Host "`n--- Teste GET /api/cliente/me (cliente completo) ---" -ForegroundColor Cyan
try {
    $cliente = Invoke-RestMethod -Uri "$BaseUrl/api/cliente/me" -Method Get -Headers $headers
    if ($cliente.performance_channels) {
        $cliente.performance_channels | ConvertTo-Json -Depth 6
    } else {
        Write-Host "performance_channels vazio ou ausente neste endpoint também." -ForegroundColor Magenta
    }
} catch {
    Write-Host "Erro (ex.: 403 se não for role client): $_" -ForegroundColor Red
}
