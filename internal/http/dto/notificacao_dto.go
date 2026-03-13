package dto

import "time"

type NotificacaoOutput struct {
	UUID     string    `json:"uuid"`
	Titulo   string    `json:"titulo"`
	Target   string    `json:"target"`
	Canal    string    `json:"canal"`
	Tipo     string    `json:"tipo"`
	Lidas    int       `json:"lidas"`
	CriadoEm time.Time `json:"criado_em"`
	Conteudo string    `json:"conteudo"`
}

