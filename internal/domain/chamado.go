package domain

import "time"

type Chamado struct {
	UUID         string    `bson:"uuid" json:"uuid"`
	ClienteUUID  string    `bson:"cliente_uuid" json:"cliente_uuid"`
	Categoria    string    `bson:"categoria" json:"categoria"`
	Titulo       string    `bson:"titulo" json:"titulo"`
	Status       string    `bson:"status" json:"status"`
	CriadoEm     time.Time `bson:"criado_em" json:"criado_em"`
	AtualizadoEm time.Time `bson:"atualizado_em" json:"atualizado_em"`
	Descricao    string    `bson:"descricao" json:"descricao"`
}

