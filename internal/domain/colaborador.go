package domain

import "time"

type Colaborador struct {
	UUID      string    `bson:"uuid"`
	Nome      string    `bson:"nome"`
	Email     string    `bson:"email"`
	Cargo     string    `bson:"cargo"`
	Role      string    `bson:"role"` // "admin" ou "user"
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

