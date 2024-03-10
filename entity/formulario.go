package entity

type Formulario struct {
	Id           uint
	Nome         string
	Propriedades []*FormularioPropriedade
	Cliente      *Cliente
	IsPadrao     bool
}

func NewFormulario(id uint) *Formulario {
	return &Formulario{
		Id:           id,
		Propriedades: make([]*FormularioPropriedade, 0),
	}
}
