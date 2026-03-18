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
	PerformanceChannels map[string]interface{} `json:"performance_channels"` // sempre presente para role client (vazio {} ou com dados)
	// Aba Performance: preenchidos a partir de performance_channels quando o admin salva
	LeadsPorPeriodo  interface{} `json:"leads_por_periodo,omitempty"`  // ex.: [ {"periodo":"2025-01","leads":10}, ... ]
	Funil            interface{} `json:"funil,omitempty"`              // ex.: {"impressoes":1000,"cliques":100,"leads":10,"conversoes":5}
	ConversoesTotais interface{} `json:"conversoes_totais,omitempty"`  // número total
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

