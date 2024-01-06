package entity

type Deteccao struct {
	Id                  uint
	Wkt                 string
	Descricao           string
	ImagemAntesUuid     string
	ImagemDepoisUuid    string
	ImagemResultadoUuid string
	AreaHa              float64
	Classe              *DeteccaoClasse
	Alerta              *Alerta
	Cliente             *Cliente
	Camadas             []*Camada
}

func NewDeteccao(id uint) *Deteccao {
	return &Deteccao{
		Id:      id,
		Camadas: []*Camada{},
	}
}
