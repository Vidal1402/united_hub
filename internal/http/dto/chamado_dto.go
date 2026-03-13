package dto

import "time"

type ChamadoOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	Categoria   string    `json:"categoria"`
	Titulo      string    `json:"titulo"`
	Status      string    `json:"status"`
	CriadoEm    time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
	Descricao   string    `json:"descricao"`
}

