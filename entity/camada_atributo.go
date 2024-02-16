package entity

type CamadaAtributo struct {
	Id     uint
	Nome   string
	Valor  string
	Camada *Camada
}

func NewCamadaAtributo(id uint) *CamadaAtributo {
	return &CamadaAtributo{
		Id: id,
	}
}
