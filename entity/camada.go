package entity

type Camada struct {
	Id         uint
	Nome       string
	Wkt        string
	Detalhe    string
	Ativo      *Ativo
	Categoria  *CamadaCategoria
	Formulario *Formulario
	Atributos  map[string]string
}

func NewCamada(id uint) *Camada {
	return &Camada{
		Id:        id,
		Atributos: make(map[string]string),
	}
}
