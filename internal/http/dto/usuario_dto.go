package dto

type UsuarioOutput struct {
	UUID           string `json:"uuid"`
	ClienteUUID    string `json:"cliente_uuid,omitempty"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	CanProducao    bool   `json:"can_producao"`
	CanPerformance bool   `json:"can_performance"`
}

type UsuarioCreateInput struct {
	ClienteUUID    string `json:"cliente_uuid" validate:"omitempty,uuid4"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=6"`
	Role           string `json:"role" validate:"required,oneof=client admin"`
	CanProducao    bool   `json:"can_producao"`
	CanPerformance bool   `json:"can_performance"`
}

type UsuarioUpdateInput struct {
	CanProducao    *bool `json:"can_producao"`
	CanPerformance *bool `json:"can_performance"`
}

