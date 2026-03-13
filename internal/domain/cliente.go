package domain

import "time"

type Cliente struct {
	UUID      string    `bson:"uuid"` // UUID v4
	Nome      string    `bson:"nome"`
	Email     string    `bson:"email"`
	Segmento  string    `bson:"segmento"`
	Plano     string    `bson:"plano"` // Starter/Growth/Pro/Scale etc
	Status    string    `bson:"status"`
	Cidade    string    `bson:"cidade"`
	OwnerUUID string    `bson:"owner_uuid"` // Colaborador responsável
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

