package entity

type Formulario struct {
	Id           uint
	Nome         string
	Propriedades []*FormularioPropriedade
}

func NewFormulario(id uint) *Formulario {
	return &Formulario{
		Id:           id,
		Propriedades: make([]*FormularioPropriedade, 0),
	}
}
