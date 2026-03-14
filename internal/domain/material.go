package domain

import "time"

type MaterialPasta struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid" json:"cliente_uuid"`
	Nome        string    `bson:"nome" json:"nome"`
	Icone       string    `bson:"icone" json:"icone"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type MaterialArquivo struct {
	UUID        string    `bson:"uuid" json:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid" json:"cliente_uuid"`
	PastaUUID   string    `bson:"pasta_uuid" json:"pasta_uuid"`
	Nome        string    `bson:"nome" json:"nome"`
	Extensao    string    `bson:"extensao" json:"extensao"`
	Tamanho     int64     `bson:"tamanho" json:"tamanho"`
	Data        time.Time `bson:"data" json:"data"`
	URL         string    `bson:"url" json:"url"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

