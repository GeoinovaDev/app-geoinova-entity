package entity

type CamadaAtributo struct {
	Id       uint
	Atributo string
	Valor    string
}

func NewCamadaAtributo(id uint, atributo string, valor string) *CamadaAtributo {
	return &CamadaAtributo{
		Id:       id,
		Atributo: atributo,
		Valor:    valor,
	}
}
