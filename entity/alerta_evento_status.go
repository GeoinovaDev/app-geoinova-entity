package entity

type AlertaEventoStatus string

func (status AlertaEventoStatus) Equals(aes AlertaEventoStatus) bool {
	return string(status) == string(aes)
}

func NewAlertaEventoStatus(status string) AlertaEventoStatus {
	return AlertaEventoStatus(status)
}
