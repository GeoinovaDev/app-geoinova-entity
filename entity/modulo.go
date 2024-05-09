package entity

type Modulo struct {
	Id   uint
	Nome string
}

func NewModulo(id uint) *Modulo {
	return &Modulo{Id: id}
}
