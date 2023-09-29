package entity

type DeteccaoCamada struct {
	Id       uint
	Deteccao *Deteccao
	Camada   *Camada
}

func NewDeteccaoCamada(id uint, deteccao *Deteccao, camada *Camada) *DeteccaoCamada {
	return &DeteccaoCamada{id, deteccao, camada}
}
