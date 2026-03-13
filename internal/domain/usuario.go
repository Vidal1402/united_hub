package domain

import "time"

type Usuario struct {
	UUID           string    `bson:"uuid"`
	ClienteUUID    string    `bson:"cliente_uuid,omitempty"`
	Email          string    `bson:"email"`
	SenhaHash      string    `bson:"senha_hash"`
	Role           string    `bson:"role"` // "client" | "admin"
	CanProducao    bool      `bson:"can_producao"`
	CanPerformance bool      `bson:"can_performance"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
}

