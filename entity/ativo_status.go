package entity

type AtivoStatus string

func (a AtivoStatus) Equals(status AtivoStatus) bool {
	return string(a) == string(status)
}

func NewAtivoStatus(status string) AtivoStatus {
	return AtivoStatus(status)
}
