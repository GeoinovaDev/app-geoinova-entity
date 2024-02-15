package entity

type CamadaCategoriaFormulario struct {
	Id         uint
	Categoria  *CamadaCategoria
	Formulario *Formulario
}

func NewCamadaCategoriaFormulario(id uint, categoria *CamadaCategoria, formulario *Formulario) *CamadaCategoriaFormulario {
	return &CamadaCategoriaFormulario{
		Id:         id,
		Categoria:  categoria,
		Formulario: formulario,
	}
}
