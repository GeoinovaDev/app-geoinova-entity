package entity

type AlertaStatus string

func (status AlertaStatus) Equals(as AlertaStatus) bool {
	return string(status) == string(as)
}
