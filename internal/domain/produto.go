package domain

import "time"

type Produto struct {
	UUID      string    `bson:"uuid" json:"uuid"`
	Familia   string    `bson:"familia" json:"familia"` // marketing/food/ia/crm
	Slug      string    `bson:"slug" json:"slug"`
	Nome      string    `bson:"nome" json:"nome"`
	Preco     int64     `bson:"preco_centavos" json:"preco_centavos"`
	Badge     string    `bson:"badge" json:"badge,omitempty"`
	Descricao string    `bson:"descricao" json:"descricao"`
	Features  []string  `bson:"features" json:"features"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

