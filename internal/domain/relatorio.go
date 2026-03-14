package domain

import "time"

type Relatorio struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid" json:"cliente_uuid"`
	Titulo      string    `bson:"titulo" json:"titulo"`
	Tipo        string    `bson:"tipo" json:"tipo"` // Mensal, Campanha, etc
	Periodo     string    `bson:"periodo" json:"periodo"`
	OwnerUUID   string    `bson:"owner_uuid" json:"owner_uuid"`
	Data        time.Time `bson:"data" json:"data"`
	Paginas     int       `bson:"paginas" json:"paginas"`
	FileURL     string    `bson:"file_url" json:"file_url"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

