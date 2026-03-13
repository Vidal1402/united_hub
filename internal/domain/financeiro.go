package domain

import "time"

type Fatura struct {
	UUID          string     `bson:"uuid"`
	ClienteUUID   string     `bson:"cliente_uuid"`
	Periodo       string     `bson:"periodo"`
	Valor         int64      `bson:"valor_centavos"`
	Vencimento    time.Time  `bson:"vencimento"`
	Status        string     `bson:"status"` // Pago/Pendente/Vencido
	DataPagamento *time.Time `bson:"data_pagamento,omitempty"`
	CreatedAt     time.Time  `bson:"created_at"`
	UpdatedAt     time.Time  `bson:"updated_at"`
}

// Recebível no ADM
type Recebivel struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	Descricao   string    `bson:"descricao"`
	Valor       int64     `bson:"valor_centavos"`
	Vencimento  time.Time `bson:"vencimento"`
	Status      string    `bson:"status"`
	Plano       string    `bson:"plano"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type Pagamento struct {
	UUID       string    `bson:"uuid"`
	Descricao  string    `bson:"descricao"`
	Valor      int64     `bson:"valor_centavos"`
	Vencimento time.Time `bson:"vencimento"`
	Status     string    `bson:"status"` // Pago/Pendente
	Categoria  string    `bson:"categoria"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

