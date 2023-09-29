package entity

type DeteccaoClasseBuilder struct {
	classe *DeteccaoClasse
}

func NewDeteccaoClasseBuilder(id uint) DeteccaoClasseBuilder {
	return DeteccaoClasseBuilder{NewDeteccaoClasse(id)}
}

func (b DeteccaoClasseBuilder) WithNome(nome string) DeteccaoClasseBuilder {
	b.classe.Nome = nome
	return b
}

func (b DeteccaoClasseBuilder) WithTipo(tipo string) DeteccaoClasseBuilder {
	b.classe.Tipo = tipo
	return b
}

func (b DeteccaoClasseBuilder) WithColor(color string) DeteccaoClasseBuilder {
	b.classe.Color = color
	return b
}

func (b DeteccaoClasseBuilder) Build() *DeteccaoClasse {
	return b.classe
}
