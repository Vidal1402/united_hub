package middleware

import (
  "log/slog"
  "net/http"

  "github.com/go-chi/httplog/v2"
)

func Logger(l *slog.Logger) func(http.Handler) http.Handler {
  return httplog.RequestLogger(httplog.NewLogger("api", httplog.Options{JSON: true, Concise: true, LogLevel: slog.LevelInfo}))
}