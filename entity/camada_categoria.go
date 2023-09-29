package entity

type CamadaCategoria struct {
	Id        uint
	Nome      string
	Cor       string
	Borda     string
	Categoria *CamadaCategoria
}

func NewCamadaCategoria(id uint) *CamadaCategoria {
	return &CamadaCategoria{Id: id}
}
