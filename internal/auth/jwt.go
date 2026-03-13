package auth

import (
  "errors"
  "time"

  "github.com/golang-jwt/jwt/v5"
  "github.com/google/uuid"
)

type Role string

const (
  RoleClient Role = "client"
  RoleAdmin  Role = "admin"
)

type Claims struct {
  jwt.RegisteredClaims
  Role     Role      `json:"role"`
  ClienteID uuid.UUID `json:"cliente_id"`
}

func ParseToken(tokenString string, secret string) (Claims, error) {
  token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
    if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("unexpected signing method")
    }
    return []byte(secret), nil
  }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
  if err != nil {
    return Claims{}, err
  }

  c, ok := token.Claims.(*Claims)
  if !ok || !token.Valid {
    return Claims{}, errors.New("invalid token")
  }

  // extra sanity: exp if provided
  if c.ExpiresAt != nil && c.ExpiresAt.Time.Before(time.Now()) {
    return Claims{}, errors.New("token expired")
  }

  return *c, nil
}