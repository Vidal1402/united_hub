# backend_united_hub

Backend Go (Clean Architecture) para United — Growth Hub + Admin Panel.

## Rodar local

1) Suba Postgres e Redis:

```powershell
cd c:\Users\luisg\backend_united_hub
docker compose up -d
```

2) Copie variáveis:

```powershell
Copy-Item .env.example .env
```

3) Rode a API:

```powershell
go run .\cmd\api\main.go
```

Health:
- GET `/healthz`
