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

func (b *GrupoPermissaoBuilder) AddPermissao(permissao *Permissao) *GrupoPermissaoBuilder {
	b.grupo.Permissao = append(b.grupo.Permissao, permissao)
	return b
}

func (b *GrupoPermissaoBuilder) WithPermissao(permissao ...*Permissao) *GrupoPermissaoBuilder {
	b.grupo.Permissao = append(b.grupo.Permissao, permissao...)
	return b
}

func (b *GrupoPermissaoBuilder) Build() *GrupoPermissao {
	return b.grupo
}
