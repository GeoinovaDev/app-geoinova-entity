package entity

type CamadaCategoriaBuilder struct {
	categoria *CamadaCategoria
}

func NewCamadaCategoriaBuilder(id uint) CamadaCategoriaBuilder {
	return CamadaCategoriaBuilder{NewCamadaCategoria(id)}
}

func (b CamadaCategoriaBuilder) WithNome(nome string) CamadaCategoriaBuilder {
	b.categoria.Nome = nome
	return b
}

func (b CamadaCategoriaBuilder) WithBorda(borda string) CamadaCategoriaBuilder {
	b.categoria.Borda = borda
	return b
}

func (b CamadaCategoriaBuilder) WithCor(cor string) CamadaCategoriaBuilder {
	b.categoria.Cor = cor
	return b
}

func (b CamadaCategoriaBuilder) WithCategoria(categoria *CamadaCategoria) CamadaCategoriaBuilder {
	b.categoria.Categoria = categoria
	return b
}

func (b CamadaCategoriaBuilder) Build() *CamadaCategoria {
	return b.categoria
}
