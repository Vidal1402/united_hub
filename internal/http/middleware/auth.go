package middleware

import (
  "context"
  "net/http"
  "strings"

  "backend_united_hub/internal/auth"
  "backend_united_hub/internal/http/response"
  "github.com/go-chi/chi/v5"
)

type ctxKey string

const (
  ctxClaims ctxKey = "claims"
)

func GetClaims(r *http.Request) (auth.Claims, bool) {
  v := r.Context().Value(ctxClaims)
  c, ok := v.(auth.Claims)
  return c, ok
}

// SkipAuthForOPTIONS responde 204 ao preflight OPTIONS sem validar token.
// Deve ser usado antes de RequireJWT/RequireRole para que o browser não receba 401 no preflight.
func SkipAuthForOPTIONS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusNoContent)
      return
    }
    next.ServeHTTP(w, r)
  })
}

func RequireJWT(secret string) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      if r.Method == http.MethodOptions {
        next.ServeHTTP(w, r)
        return
      }
      h := r.Header.Get("Authorization")
      parts := strings.SplitN(h, " ", 2)
      if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
        response.Error(w, http.StatusUnauthorized, "missing token", nil)
        return
      }

      claims, err := auth.ParseToken(parts[1], secret)
      if err != nil {
        response.Error(w, http.StatusUnauthorized, "invalid token", nil)
        return
      }

      ctx := r.Context()
      ctx = contextWithClaims(ctx, claims)
      next.ServeHTTP(w, r.WithContext(ctx))
    })
  }
}

func RequireRole(role auth.Role) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      if r.Method == http.MethodOptions {
        next.ServeHTTP(w, r)
        return
      }
      c, ok := GetClaims(r)
      if !ok {
        response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
        return
      }
      if c.Role != role {
        response.Error(w, http.StatusForbidden, "forbidden", nil)
        return
      }
      next.ServeHTTP(w, r)
    })
  }
}

func contextWithClaims(ctx context.Context, c auth.Claims) context.Context {
  return context.WithValue(ctx, ctxClaims, c)
}

// Attach chi import to avoid unused in some builds
var _ = chi.RouteContext