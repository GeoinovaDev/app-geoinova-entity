package entity

import "time"

const (
	ALERTA_FASE_ABERTA  = AlertaFaseStatus("AB")
	ALERTA_FASE_FECHADA = AlertaFaseStatus("FC")
	ALERTA_FASE_ANALISE = AlertaFaseStatus("AN")
)

type AlertaFase struct {
	Id        uint
	Status    AlertaFaseStatus
	Alerta    *Alerta
	CreatedAt time.Time
}

func NewAlertaFase(id uint, status AlertaFaseStatus) *AlertaFase {
	return &AlertaFase{
		Id:     id,
		Status: status,
	}
}
