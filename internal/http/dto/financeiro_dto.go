package dto

import "time"

type FaturaOutput struct {
	UUID        string     `json:"uuid"`
	ClienteUUID string     `json:"cliente_uuid"`
	Periodo     string     `json:"periodo"`
	Valor       string     `json:"valor"` // formatado ex.: "R$ 4.800,00"
	Vencimento  time.Time  `json:"vencimento"`
	Status      string     `json:"status"`
	DataPgto    *time.Time `json:"data_pagamento,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type RecebivelOutput struct {
	UUID        string    `json:"uuid"`
	ClienteUUID string    `json:"cliente_uuid"`
	Descricao   string    `json:"descricao"`
	Valor       string    `json:"valor"` // formatado
	Vencimento  time.Time `json:"vencimento"`
	Status      string    `json:"status"`
	Plano       string    `json:"plano"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PagamentoOutput struct {
	UUID       string    `json:"uuid"`
	Descricao  string    `json:"descricao"`
	Valor      string    `json:"valor"` // formatado
	Vencimento time.Time `json:"vencimento"`
	Status     string    `json:"status"`
	Categoria  string    `json:"categoria"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

