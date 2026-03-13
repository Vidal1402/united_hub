package domain

import "time"

type Curso struct {
	Slug      string    `bson:"slug"`
	Titulo    string    `bson:"titulo"`
	Categoria string    `bson:"categoria"`
	Formato   string    `bson:"formato"`
	Duracao   string    `bson:"duracao"`
	Nivel     string    `bson:"nivel"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type CursoProgresso struct {
	ClienteUUID string    `bson:"cliente_uuid"`
	CursoSlug   string    `bson:"curso_slug"`
	Concluido   bool      `bson:"concluido"`
	Progresso   int       `bson:"progresso"` // 0-100
	UpdatedAt   time.Time `bson:"updated_at"`
}

