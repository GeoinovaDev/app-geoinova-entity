package entity

type DeteccaoClasse struct {
	Id    uint
	Nome  string
	Tipo  string
	Color string
}

func NewDeteccaoClasse(id uint) *DeteccaoClasse {
	return &DeteccaoClasse{Id: id}
}
