package domain

import "time"

type MaterialPasta struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	Nome        string    `bson:"nome"`
	Icone       string    `bson:"icone"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type MaterialArquivo struct {
	UUID        string    `bson:"uuid"`
	ClienteUUID string    `bson:"cliente_uuid"`
	PastaUUID   string    `bson:"pasta_uuid"`
	Nome        string    `bson:"nome"`
	Extensao    string    `bson:"extensao"`
	Tamanho     int64     `bson:"tamanho"`
	Data        time.Time `bson:"data"`
	URL         string    `bson:"url"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

