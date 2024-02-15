package entity

type FormularioPropriedadeBuilder struct {
	formularioPropriedade *FormularioPropriedade
}

func NewFormularioPropriedadeBuilder(id uint) *FormularioPropriedadeBuilder {
	return &FormularioPropriedadeBuilder{
		formularioPropriedade: NewFormularioPropriedade(id),
	}
}

func (b *FormularioPropriedadeBuilder) WithNome(nome string) *FormularioPropriedadeBuilder {
	b.formularioPropriedade.Nome = nome
	return b
}

func (b *FormularioPropriedadeBuilder) WithTipo(tipo FormularioPropriedadeTipo) *FormularioPropriedadeBuilder {
	b.formularioPropriedade.Tipo = tipo
	return b
}

func (b *FormularioPropriedadeBuilder) WithField(field string) *FormularioPropriedadeBuilder {
	b.formularioPropriedade.Field = field
	return b
}

func (b *FormularioPropriedadeBuilder) Build() *FormularioPropriedade {
	return b.formularioPropriedade
}
