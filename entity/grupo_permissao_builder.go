package entity

type GrupoPermissaoBuilder struct {
	grupo *GrupoPermissao
}

func NewGrupoPermissaoBuilder(id uint) *GrupoPermissaoBuilder {
	return &GrupoPermissaoBuilder{
		grupo: NewGrupoPermissao(id),
	}
}

func (b *GrupoPermissaoBuilder) WithNome(nome string) *GrupoPermissaoBuilder {
	b.grupo.Nome = nome
	return b
}

func (b *GrupoPermissaoBuilder) WithCliente(cliente *Cliente) *GrupoPermissaoBuilder {
	b.grupo.Cliente = cliente
	return b
}

func (b *GrupoPermissaoBuilder) AddPermissao(permissao *Permissao) *GrupoPermissaoBuilder {
	b.grupo.Permissoes = append(b.grupo.Permissoes, permissao)
	return b
}

func (b *GrupoPermissaoBuilder) WithPermissao(permissao ...*Permissao) *GrupoPermissaoBuilder {
	b.grupo.Permissoes = append(b.grupo.Permissoes, permissao...)
	return b
}

func (b *GrupoPermissaoBuilder) Build() *GrupoPermissao {
	return b.grupo
}
