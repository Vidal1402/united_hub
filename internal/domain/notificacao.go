package domain

import "time"

type Notificacao struct {
	UUID     string    `bson:"uuid" json:"uuid"`
	Titulo   string    `bson:"titulo" json:"titulo"`
	Target   string    `bson:"target" json:"target"`   // cliente específico ou "todos"
	Canal    string    `bson:"canal" json:"canal"`    // Email/Plataforma/WhatsApp
	Tipo     string    `bson:"tipo" json:"tipo"`      // sistema, financeiro, etc
	Lidas    int       `bson:"lidas" json:"lidas"`
	CriadoEm time.Time `bson:"criado_em" json:"criado_em"`
	Conteudo string    `bson:"conteudo" json:"conteudo"`
}

