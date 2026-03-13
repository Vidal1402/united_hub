package domain

import "time"

type Reuniao struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	Titulo      string    `bson:"titulo"`
	DataHora    time.Time `bson:"data_hora"`
	Via         string    `bson:"via"`
	OwnerUUID   string    `bson:"owner_uuid"`
	Pauta       []string  `bson:"pauta"`
	Status      string    `bson:"status"` // futura/historico etc
	DuracaoMin  int       `bson:"duracao_min"`
	TemGravacao bool      `bson:"tem_gravacao"`
	TemAta      bool      `bson:"tem_ata"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

