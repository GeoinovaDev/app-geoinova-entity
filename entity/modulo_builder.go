package entity

type ModuloBuilder struct {
	modulo *Modulo
}

func NewModuloBuilder(id uint) *ModuloBuilder {
	return &ModuloBuilder{NewModulo(id)}
}

func (b *ModuloBuilder) WithNome(nome string) *ModuloBuilder {
	b.modulo.Nome = nome
	return b
}

func (b *ModuloBuilder) Build() *Modulo {
	return b.modulo
}
