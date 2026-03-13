package dto

import "time"

type ProdutoOutput struct {
	UUID      string    `json:"uuid"`
	Familia   string    `json:"familia"`
	Slug      string    `json:"slug"`
	Nome      string    `json:"nome"`
	Preco     string    `json:"preco"` // formatado
	Descricao string    `json:"descricao"`
	Features  []string  `json:"features"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

