package entity

type Deteccao struct {
	Id                  uint
	Wkt                 string
	Descricao           string
	ImagemAntesUuid     string
	ImagemDepoisUuid    string
	ImagemResultadoUuid string
	Classe              *DeteccaoClasse
	Alerta              *Alerta
	Camadas             []*Camada
}

func NewDeteccao(id uint) *Deteccao {
	return &Deteccao{
		Id:      id,
		Camadas: []*Camada{},
	}
}
