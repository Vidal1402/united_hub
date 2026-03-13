package dto

import "time"

type KanbanColumnOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	Key         string    `json:"key"`
	Label       string    `json:"label"`
	Ordenacao   int       `json:"ordenacao"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type KanbanCardOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	ColumnUUID  string    `json:"column_uuid"`
	Titulo      string    `json:"titulo"`
	Tipo        string    `json:"tipo"`
	OwnerUUID   string    `json:"owner_uuid"`
	Prazo       time.Time `json:"prazo"`
	Prioridade  string    `json:"prioridade"`
	Comentarios int       `json:"comentarios"`
	Arquivos    int       `json:"arquivos"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

