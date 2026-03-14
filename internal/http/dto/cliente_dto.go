package dto

import "time"

// ClienteInput representa dados de entrada para criação/atualização de cliente.
type ClienteInput struct {
	Nome      string `json:"nome" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Segmento  string `json:"segmento" validate:"required"`
	Plano     string `json:"plano" validate:"required"`
	Status    string `json:"status" validate:"required"`
	Cidade    string `json:"cidade" validate:"required"`
	OwnerUUID string `json:"owner_uuid" validate:"omitempty,uuid4"` // opcional; front pode enviar vazio
}

// ClienteOutput representa o cliente retornado na API.
type ClienteOutput struct {
	UUID      string    `json:"uuid"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Segmento  string    `json:"segmento"`
	Plano     string    `json:"plano"`
	Status    string    `json:"status"`
	Cidade    string    `json:"cidade"`
	OwnerUUID string    `json:"owner_uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

