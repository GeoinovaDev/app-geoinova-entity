package entity

type CamadaAtributo struct {
	Id    uint
	Nome  string
	Valor string
}

func NewCamadaAtributo(id uint, nome string, valor string) *CamadaAtributo {
	return &CamadaAtributo{
		Id:    id,
		Nome:  nome,
		Valor: valor,
	}
}
