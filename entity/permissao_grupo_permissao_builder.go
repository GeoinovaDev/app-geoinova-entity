package entity

type PermissaoGrupoPermissaoBuilder struct {
	permissaoGrupoPermissao *PermissaoGrupoPermissao
}

func NewPermissaoGrupoPermissaoBuilder(id uint) *PermissaoGrupoPermissaoBuilder {
	return &PermissaoGrupoPermissaoBuilder{
		permissaoGrupoPermissao: NewPermissaoGrupoPermissao(id),
	}
}

func (b *PermissaoGrupoPermissaoBuilder) WithEnabled(enabled bool) *PermissaoGrupoPermissaoBuilder {
	b.permissaoGrupoPermissao.Enabled = enabled
	return b
}

func (b *PermissaoGrupoPermissaoBuilder) WithPermissao(permissao *Permissao) *PermissaoGrupoPermissaoBuilder {
	b.permissaoGrupoPermissao.Permissao = permissao
	return b
}

func (b *PermissaoGrupoPermissaoBuilder) Build() *PermissaoGrupoPermissao {
	return b.permissaoGrupoPermissao
}
