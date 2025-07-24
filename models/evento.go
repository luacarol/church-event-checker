package models

import "time"

type Evento struct {
	ID   uint      `json:"id" gorm:"primaryKey"`
	Tipo string    `json:"tipo"`
	Data time.Time `json:"data"`

	AberturaTexto           string `json:"abertura_texto"`
	AberturaResponsavel     string `json:"abertura_responsavel"`
	AberturaTipoResponsavel string `json:"abertura_tipo_responsavel"`

	PalavraTexto           string `json:"palavra_texto"`
	PalavraResponsavel     string `json:"palavra_responsavel"`
	PalavraTipoResponsavel string `json:"palavra_tipo_responsavel"`
	PossuiTema             bool   `json:"possui_tema"`
	Tema                   string `json:"tema"`

	Louvores  string `json:"louvores"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
