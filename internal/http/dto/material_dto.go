package dto

import "time"

type MaterialPastaOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	Nome        string    `json:"nome"`
	Icone       string    `json:"icone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MaterialArquivoOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	PastaUUID   string    `json:"pasta_uuid"`
	Nome        string    `json:"nome"`
	Extensao    string    `json:"extensao"`
	Tamanho     int64     `json:"tamanho"`
	Data        time.Time `json:"data"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

