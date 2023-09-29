package entity

type AnaliseStatus string

func (status AnaliseStatus) Equals(as AnaliseStatus) bool {
	return string(status) == string(as)
}
