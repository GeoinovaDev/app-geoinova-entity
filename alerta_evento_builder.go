package entity

import "time"

type AlertaEventoBuilder struct {
	evento *AlertaEvento
}

func NewAlertaEventoBuilder(id uint, status AlertaEventoStatus) AlertaEventoBuilder {
	return AlertaEventoBuilder{NewAlertaEvento(id, status)}
}

func (b AlertaEventoBuilder) WithCreatedAt(createdAt time.Time) AlertaEventoBuilder {
	b.evento.CreatedAt = createdAt
	return b
}

func (b AlertaEventoBuilder) WithParam(param string) AlertaEventoBuilder {
	b.evento.Param = param
	return b
}

func (b AlertaEventoBuilder) WithAlerta(alerta *Alerta) AlertaEventoBuilder {
	b.evento.Alerta = alerta
	return b
}

func (b AlertaEventoBuilder) Build() *AlertaEvento {
	return b.evento
}
