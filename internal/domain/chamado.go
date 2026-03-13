package domain

import "time"

type Chamado struct {
	UUID         string    `bson:"uuid"`
	ClienteUUID  string    `bson:"cliente_uuid"`
	Categoria    string    `bson:"categoria"`
	Titulo       string    `bson:"titulo"`
	Status       string    `bson:"status"`
	CriadoEm     time.Time `bson:"criado_em"`
	AtualizadoEm time.Time `bson:"atualizado_em"`
	Descricao    string    `bson:"descricao"`
}

