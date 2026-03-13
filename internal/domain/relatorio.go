package domain

import "time"

type Relatorio struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	Titulo      string    `bson:"titulo"`
	Tipo        string    `bson:"tipo"` // Mensal, Campanha, etc
	Periodo     string    `bson:"periodo"`
	OwnerUUID   string    `bson:"owner_uuid"`
	Data        time.Time `bson:"data"`
	Paginas     int       `bson:"paginas"`
	FileURL     string    `bson:"file_url"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

