package domain

import "time"

type KanbanColumn struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	Key         string    `bson:"key"`   // ex: "sol", "pend", "prod"
	Label       string    `bson:"label"` // "Solicitações", etc
	Ordenacao   int       `bson:"ordenacao"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type KanbanCard struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	ColumnUUID  string    `bson:"column_uuid"`
	Titulo      string    `bson:"titulo"`
	Tipo        string    `bson:"tipo"`
	OwnerUUID   string    `bson:"owner_uuid"`
	Prazo       time.Time `bson:"prazo"`
	Prioridade  string    `bson:"prioridade"` // Alta/Média/Baixa
	Comentarios int       `bson:"comentarios"`
	Arquivos    int       `bson:"arquivos"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

