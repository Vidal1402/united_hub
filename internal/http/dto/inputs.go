package dto

import "time"

// DTOs de entrada usados pelos services/handlers.

type UploadMaterialInput struct {
	PastaUUID string `json:"pasta_uuid" validate:"required"`
	Nome      string `json:"nome" validate:"required"`
	Extensao  string `json:"extensao" validate:"required"`
	Tamanho   int64  `json:"tamanho" validate:"required,gte=0"`
	URL       string `json:"url" validate:"required,url"`
	Data      time.Time `json:"data"` // opcional, pode vir do client
}

type CreateChamadoInput struct {
	Categoria string `json:"categoria" validate:"required"`
	Titulo    string `json:"titulo" validate:"required"`
	Descricao string `json:"descricao" validate:"required"`
}

type UpdatePerfilInput struct {
	Nome   string `json:"nome" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Cidade string `json:"cidade" validate:"required"`
}

type UpdateNotificacoesConfigInput struct {
	CanalEmail     bool `json:"canal_email"`
	CanalPlataforma bool `json:"canal_plataforma"`
	CanalWhatsApp  bool `json:"canal_whatsapp"`
}

type CreateClienteInput = ClienteInput

type UpdateClienteInput = ClienteInput

