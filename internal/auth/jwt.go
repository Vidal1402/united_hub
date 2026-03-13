package auth

import (
	"errors"
	"time"

	"backend_united_hub/internal/domain"
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
	Role           Role      `json:"role"`
	ClienteID      uuid.UUID `json:"cliente_id"`
	UsuarioID      uuid.UUID `json:"usuario_id"`
	Email          string    `json:"email"`
	CanProducao    bool      `json:"can_producao"`
	CanPerformance bool      `json:"can_performance"`
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

// SignToken cria um JWT para o usuário informado.
func SignToken(u domain.Usuario, clienteID uuid.UUID, secret string, ttl time.Duration) (string, error) {
	now := time.Now()
	uid, err := uuid.Parse(u.UUID)
	if err != nil {
		return "", err
	}

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.UUID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
		Role:           Role(u.Role),
		ClienteID:      clienteID,
		UsuarioID:      uid,
		Email:          u.Email,
		CanProducao:    u.CanProducao,
		CanPerformance: u.CanPerformance,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}
