package entity

type ImagemQualidadeBuilder struct {
	qualidade *ImagemQualidade
}

func NewImagemQualidadeBuilder(id uint) ImagemQualidadeBuilder {
	return ImagemQualidadeBuilder{NewImagemQualidade(id)}
}

func (b ImagemQualidadeBuilder) WithNome(nome string) ImagemQualidadeBuilder {
	b.qualidade.Nome = nome
	return b
}

func (b ImagemQualidadeBuilder) WithColor(color string) ImagemQualidadeBuilder {
	b.qualidade.Color = color
	return b
}

func (b ImagemQualidadeBuilder) WithBackground(background string) ImagemQualidadeBuilder {
	b.qualidade.Background = background
	return b
}

func (b ImagemQualidadeBuilder) WithTipo(tipo uint) ImagemQualidadeBuilder {
	b.qualidade.Tipo = tipo
	return b
}

func (b ImagemQualidadeBuilder) Build() *ImagemQualidade {
	return b.qualidade
}
