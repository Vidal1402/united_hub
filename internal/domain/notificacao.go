package domain

import "time"

type Notificacao struct {
	UUID     string    `bson:"uuid"`
	Titulo   string    `bson:"titulo"`
	Target   string    `bson:"target"` // cliente específico ou "todos"
	Canal    string    `bson:"canal"`  // Email/Plataforma/WhatsApp
	Tipo     string    `bson:"tipo"`   // sistema, financeiro, etc
	Lidas    int       `bson:"lidas"`
	CriadoEm time.Time `bson:"criado_em"`
	Conteudo string    `bson:"conteudo"`
}

