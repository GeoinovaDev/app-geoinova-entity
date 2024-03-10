package entity

type FormularioBuilder struct {
	formulario *Formulario
}

func NewFormularioBuilder(id uint) *FormularioBuilder {
	return &FormularioBuilder{
		formulario: NewFormulario(id),
	}
}

func (b *FormularioBuilder) WithNome(nome string) *FormularioBuilder {
	b.formulario.Nome = nome
	return b
}

func (b *FormularioBuilder) WithPropriedade(propriedades []*FormularioPropriedade) *FormularioBuilder {
	b.formulario.Propriedades = propriedades
	return b
}

func (b *FormularioBuilder) WithIsPadrao(isPadrao bool) *FormularioBuilder {
	b.formulario.IsPadrao = isPadrao
	return b
}

func (b *FormularioBuilder) WithCliente(cliente *Cliente) *FormularioBuilder {
	b.formulario.Cliente = cliente
	return b
}

func (b *FormularioBuilder) AddPropriedade(propriedade ...*FormularioPropriedade) *FormularioBuilder {
	b.formulario.Propriedades = append(b.formulario.Propriedades, propriedade...)
	return b
}

func (b *FormularioBuilder) Build() *Formulario {
	return b.formulario
}
