package dto

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserInfo struct {
	Name                string                 `json:"name"`
	Email               string                 `json:"email"`
	Role                string                 `json:"role"`
	ClienteUUID         string                 `json:"cliente_uuid,omitempty"`
	CanProducao         bool                   `json:"can_producao"`
	CanPerformance      bool                   `json:"can_performance"`
	PerformanceChannels map[string]interface{} `json:"performance_channels,omitempty"` // preenchido para role client (aba Performance)
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

