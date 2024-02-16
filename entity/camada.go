package entity

type Camada struct {
	Id         uint
	Nome       string
	Wkt        string
	Detalhe    string
	Area       float32
	Ativo      *Ativo
	Categoria  *CamadaCategoria
	Formulario *Formulario
	Atributos  []*CamadaAtributo
}

func NewCamada(id uint) *Camada {
	return &Camada{
		Id:        id,
		Atributos: make([]*CamadaAtributo, 0),
	}
}
