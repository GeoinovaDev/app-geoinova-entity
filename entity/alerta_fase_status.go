package entity

type AlertaFaseStatus string

func (status AlertaFaseStatus) Equals(afs AlertaFaseStatus) bool {
	return string(status) == string(afs)
}
