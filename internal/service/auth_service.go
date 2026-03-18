package service

import (
	"context"
	"errors"
	"strings"
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

	hash := strings.TrimSpace(u.SenhaHash)
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(in.Password)); err != nil {
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
// Para role client, busca o documento do cliente e preenche performance_channels (aba Performance).
func (s *AuthService) Me(ctx context.Context, claims auth.Claims) (dto.UserInfo, error) {
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

	if u != nil {
		info.CanProducao = u.CanProducao
		info.CanPerformance = u.CanPerformance
	}

	// Cliente: buscar documento do cliente para retornar performance_channels e dados da aba Performance.
	// Sempre devolver performance_channels para role client (objeto vazio se não houver dados).
	if claims.Role == auth.RoleClient && claims.ClienteID != uuid.Nil {
		info.PerformanceChannels = map[string]interface{}{}
		cliente, err := s.clientes.GetByUUID(ctx, claims.ClienteID.String())
		if err == nil && cliente != nil && cliente.PerformanceChannels != nil {
			info.PerformanceChannels = cliente.PerformanceChannels
			// Expor na raiz para o front preencher Leads por Período, Funil, Conversões
			if v, ok := cliente.PerformanceChannels["leads_por_periodo"]; ok && v != nil {
				info.LeadsPorPeriodo = v
			}
			if v, ok := cliente.PerformanceChannels["funil"]; ok && v != nil {
				info.Funil = v
			}
			if v, ok := cliente.PerformanceChannels["conversoes_totais"]; ok && v != nil {
				info.ConversoesTotais = v
			}
		}
	}

	return info, nil
}

