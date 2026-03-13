package dto

import "time"

type ReuniaoOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	Titulo      string    `json:"titulo"`
	DataHora    time.Time `json:"data_hora"`
	Via         string    `json:"via"`
	OwnerUUID   string    `json:"owner_uuid"`
	Pauta       []string  `json:"pauta"`
	Status      string    `json:"status"`
	DuracaoMin  int       `json:"duracao_min"`
	TemGravacao bool      `json:"tem_gravacao"`
	TemAta      bool      `json:"tem_ata"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

