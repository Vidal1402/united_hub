package domain

import "time"

type Produto struct {
	UUID      string    `bson:"uuid"`
	Familia   string    `bson:"familia"` // marketing/food/ia/crm
	Slug      string    `bson:"slug"`
	Nome      string    `bson:"nome"`
	Preco     int64     `bson:"preco_centavos"`
	Descricao string    `bson:"descricao"`
	Features  []string  `bson:"features"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

