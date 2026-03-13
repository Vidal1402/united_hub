package dto

import "time"

type CursoOutput struct {
	Slug      string    `json:"slug"`
	Titulo    string    `json:"titulo"`
	Categoria string    `json:"categoria"`
	Formato   string    `json:"formato"`
	Duracao   string    `json:"duracao"`
	Nivel     string    `json:"nivel"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CursoProgressoOutput struct {
	ClienteUUID string    `json:"cliente_uuid"`
	CursoSlug   string    `json:"curso_slug"`
	Concluido   bool      `json:"concluido"`
	Progresso   int       `json:"progresso"`
	UpdatedAt   time.Time `json:"updated_at"`
}

