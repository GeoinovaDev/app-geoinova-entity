package entity

type PermissaoCategoriaBuilder struct {
	categoria *PermissaoCategoria
}

func NewPermissaoCategoriaBuilder(id uint) PermissaoCategoriaBuilder {
	return PermissaoCategoriaBuilder{
		categoria: NewPermissaoCategoria(id),
	}
}

func (b PermissaoCategoriaBuilder) WithNome(nome string) PermissaoCategoriaBuilder {
	b.categoria.Nome = nome
	return b
}

func (b PermissaoCategoriaBuilder) Build() *PermissaoCategoria {
	return b.categoria
}
