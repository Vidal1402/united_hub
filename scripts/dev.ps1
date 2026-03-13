$ErrorActionPreference = "Stop"

Set-Location "$PSScriptRoot\.."

docker compose up -d
Copy-Item .env.example .env -Force

go run .\cmd\api\main.go