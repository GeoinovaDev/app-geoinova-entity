package entity

type IncendioContato struct {
	Emails    []string
	Telefones []*Telefone
}

func NewIncendioContato() *IncendioContato {
	return &IncendioContato{}
}
