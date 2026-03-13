package domain

import "time"

type Alerta struct {
	UUID        string     `bson:"uuid"`
	ClienteUUID string     `bson:"cliente_uuid,omitempty"`
	Titulo      string     `bson:"titulo"`
	Tipo        string     `bson:"tipo"`
	Prioridade  string     `bson:"prioridade"`
	Target      string     `bson:"target"`
	Status      string     `bson:"status"` // Ativo/Resolvido
	CreatedAt   time.Time  `bson:"created_at"`
	ResolvedAt  *time.Time `bson:"resolved_at,omitempty"`
}

