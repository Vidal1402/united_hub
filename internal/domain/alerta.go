package domain

import "time"

type Alerta struct {
	UUID        string     `bson:"uuid" json:"uuid"`
	ClienteUUID string     `bson:"cliente_uuid,omitempty" json:"cliente_uuid,omitempty"`
	Titulo      string     `bson:"titulo" json:"titulo"`
	Tipo        string     `bson:"tipo" json:"tipo"`
	Prioridade  string     `bson:"prioridade" json:"prioridade"`
	Target      string     `bson:"target" json:"target"`
	Status      string     `bson:"status" json:"status"` // Ativo/Resolvido
	CreatedAt   time.Time  `bson:"created_at" json:"created_at"`
	ResolvedAt  *time.Time `bson:"resolved_at,omitempty" json:"resolved_at,omitempty"`
}

