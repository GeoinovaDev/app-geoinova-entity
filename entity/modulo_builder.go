package entity

type ModuloBuilder struct {
	modulo *Modulo
}

func NewModuloBuilder() *ModuloBuilder {
	return &ModuloBuilder{NewModulo(0)}
}

func (b *ModuloBuilder) WithNome(nome string) *ModuloBuilder {
	b.modulo.Nome = nome
	return b
}

func (b *ModuloBuilder) Build() *Modulo {
	return b.modulo
}
