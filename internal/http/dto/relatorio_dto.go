package dto

import "time"

type RelatorioOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	Titulo      string    `json:"titulo"`
	Tipo        string    `json:"tipo"`
	Periodo     string    `json:"periodo"`
	OwnerUUID   string    `json:"owner_uuid"`
	Data        time.Time `json:"data"`
	Paginas     int       `json:"paginas"`
	FileURL     string    `json:"file_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

