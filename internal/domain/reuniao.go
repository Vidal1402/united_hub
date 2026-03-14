package domain

import "time"

type Reuniao struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid" json:"cliente_uuid"`
	Titulo      string    `bson:"titulo" json:"titulo"`
	DataHora    time.Time `bson:"data_hora" json:"data_hora"`
	Via         string    `bson:"via" json:"via"`
	OwnerUUID   string    `bson:"owner_uuid" json:"owner_uuid"`
	Pauta       []string  `bson:"pauta" json:"pauta"`
	Status      string    `bson:"status" json:"status"` // futura/historico etc
	DuracaoMin  int       `bson:"duracao_min" json:"duracao_min"`
	TemGravacao bool      `bson:"tem_gravacao" json:"tem_gravacao"`
	TemAta      bool      `bson:"tem_ata" json:"tem_ata"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

