package domain

import "time"

type Cliente struct {
	UUID                string                 `bson:"uuid" json:"uuid"`
	Nome                string                 `bson:"nome" json:"nome"`
	Email               string                 `bson:"email" json:"email"`
	Segmento            string                 `bson:"segmento" json:"segmento"`
	Plano               string                 `bson:"plano" json:"plano"`
	Status              string                 `bson:"status" json:"status"`
	Cidade              string                 `bson:"cidade" json:"cidade"`
	OwnerUUID           string                 `bson:"owner_uuid" json:"owner_uuid"`
	PerformanceChannels map[string]interface{} `bson:"performance_channels,omitempty" json:"performance_channels,omitempty"` // meta_ads: {gasto, leads, conversoes}, etc.
	CreatedAt           time.Time              `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time              `bson:"updated_at" json:"updated_at"`
}

