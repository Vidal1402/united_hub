package domain

import "time"

type Colaborador struct {
	UUID      string    `bson:"uuid" json:"uuid"`
	Nome      string    `bson:"nome" json:"nome"`
	Email     string    `bson:"email" json:"email"`
	Cargo     string    `bson:"cargo" json:"cargo"`
	Role      string    `bson:"role" json:"role"` // Colaborador, Gestor, Admin
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

