package domain

import "time"

type KanbanColumn struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid" json:"cliente_uuid"`
	Key         string    `bson:"key" json:"key"`   // ex: backlog, doing, review, done
	Label       string    `bson:"label" json:"label"`
	Ordenacao   int       `bson:"ordenacao" json:"ordenacao"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type KanbanCard struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid" json:"cliente_uuid"`
	ColumnID    string    `bson:"column_id" json:"column_id"` // backlog, doing, review, done
	ColumnUUID  string    `bson:"column_uuid" json:"column_uuid"`
	Titulo      string    `bson:"titulo" json:"titulo"`
	Tipo        string    `bson:"tipo" json:"tipo"`
	Prioridade  string    `bson:"prioridade" json:"prioridade"`
	Descricao   string    `bson:"descricao" json:"descricao"`
	OwnerUUID   string    `bson:"owner_uuid" json:"owner_uuid"`
	OwnerNome   string    `bson:"owner_nome" json:"owner_nome"`
	Prazo       time.Time `bson:"prazo" json:"prazo"`
	Comentarios int       `bson:"comentarios" json:"comentarios"`
	Arquivos    int       `bson:"arquivos" json:"arquivos"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

