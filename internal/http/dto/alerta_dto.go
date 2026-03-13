package dto

import "time"

type AlertaOutput struct {
	UUID        string     `json:"uuid"`
	ClienteUUID string     `json:"cliente_uuid,omitempty"`
	Titulo      string     `json:"titulo"`
	Tipo        string     `json:"tipo"`
	Prioridade  string     `json:"prioridade"`
	Target      string     `json:"target"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
}

