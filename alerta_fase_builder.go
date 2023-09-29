package entity

import "time"

type AlertaFaseBuilder struct {
	fase *AlertaFase
}

func NewAlertaFaseBuilder(id uint, status AlertaFaseStatus) AlertaFaseBuilder {
	return AlertaFaseBuilder{NewAlertaFase(id, status)}
}

func (b AlertaFaseBuilder) WithCreatedAt(createdAt time.Time) AlertaFaseBuilder {
	b.fase.CreatedAt = createdAt
	return b
}

func (b AlertaFaseBuilder) WithAlerta(alerta *Alerta) AlertaFaseBuilder {
	b.fase.Alerta = alerta
	return b
}

func (b AlertaFaseBuilder) Build() *AlertaFase {
	return b.fase
}
