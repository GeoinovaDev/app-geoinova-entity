package entity

import "time"

const (
	ALERTA_EVENTO_CRIADO = AlertaEventoStatus("CR")
)

type AlertaEvento struct {
	Id        uint
	Status    AlertaEventoStatus
	Param     string
	Alerta    *Alerta
	CreatedAt time.Time
}

func NewAlertaEvento(id uint, status AlertaEventoStatus) *AlertaEvento {
	return &AlertaEvento{
		Id:     id,
		Status: status,
	}
}
