package service

import (
	"context"
	"errors"
	"time"

	"backend_united_hub/internal/auth"
	"backend_united_hub/internal/domain"
	"backend_united_hub/internal/http/dto"
	"backend_united_hub/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	usuarios repository.UsuarioRepository
	clientes repository.ClienteRepository
	jwtSecret string
}

func NewAuthService(
	usuarios repository.UsuarioRepository,
	clientes repository.ClienteRepository,
	jwtSecret string,
) *AuthService {
	return &AuthService{
		usuarios: usuarios,
		clientes: clientes,
		jwtSecret: jwtSecret,
	}
}

// Login executa o fluxo de autenticação básico com email/senha.
func (s *AuthService) Login(ctx context.Context, in dto.LoginInput) (dto.LoginResponse, error) {
	u, err := s.usuarios.GetByEmail(ctx, in.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	if u == nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.SenhaHash), []byte(in.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	var clienteID uuid.UUID
	if u.Role == "client" && u.ClienteUUID != "" {
		clienteID, err = uuid.Parse(u.ClienteUUID)
		if err != nil {
			return dto.LoginResponse{}, err
		}
	} else {
		clienteID = uuid.Nil
	}

	token, err := auth.SignToken(*u, clienteID, s.jwtSecret, 24*time.Hour)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	info := dto.UserInfo{
		Name:           u.Email,
		Email:          u.Email,
		Role:           u.Role,
		ClienteUUID:    u.ClienteUUID,
		CanProducao:    u.CanProducao,
		CanPerformance: u.CanPerformance,
	}

	return dto.LoginResponse{
		Token: token,
		User:  info,
	}, nil
}

// Me devolve as informações básicas de usuário a partir das claims.
func (s *AuthService) Me(ctx context.Context, claims auth.Claims) (dto.UserInfo, error) {
	// Opcionalmente buscar usuário para garantir que ainda existe.
	var u *domain.Usuario
	if claims.UsuarioID != uuid.Nil {
		user, err := s.usuarios.GetByUUID(ctx, claims.UsuarioID.String())
		if err != nil {
			return dto.UserInfo{}, err
		}
		u = user
	}

	info := dto.UserInfo{
		Name:           claims.Email,
		Email:          claims.Email,
		Role:           string(claims.Role),
		ClienteUUID:    "",
		CanProducao:    claims.CanProducao,
		CanPerformance: claims.CanPerformance,
	}
	if claims.ClienteID != uuid.Nil {
		info.ClienteUUID = claims.ClienteID.String()
	}

	// Se carregamos o usuário do banco, podemos refinar info.
	if u != nil {
		info.CanProducao = u.CanProducao
		info.CanPerformance = u.CanPerformance
	}

	return info, nil
}

