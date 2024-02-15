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

func (b *FormularioPropriedadeBuilder) WithAtributos(atributos []string) *FormularioPropriedadeBuilder {
	b.formularioPropriedade.Atributos = atributos
	return b
}

func (b *FormularioPropriedadeBuilder) AddAtributo(atributo ...string) *FormularioPropriedadeBuilder {
	b.formularioPropriedade.Atributos = append(b.formularioPropriedade.Atributos, atributo...)
	return b
}

func (b *FormularioPropriedadeBuilder) Build() *FormularioPropriedade {
	return b.formularioPropriedade
}
