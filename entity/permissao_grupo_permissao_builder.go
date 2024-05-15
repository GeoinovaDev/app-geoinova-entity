package entity

type PermissaoGrupoPermissaoBuilder struct {
	permissaoGrupoPermissao *PermissaoGrupoPermissao
}

func NewPermissaoGrupoPermissaoBuilder(id uint) *PermissaoGrupoPermissaoBuilder {
	return &PermissaoGrupoPermissaoBuilder{
		permissaoGrupoPermissao: NewPermissaoGrupoPermissao(id),
	}
}

func (b *PermissaoGrupoPermissaoBuilder) WithCliente(cliente *Cliente) *PermissaoGrupoPermissaoBuilder {
	b.permissaoGrupoPermissao.Cliente = cliente
	return b
}

func (b *PermissaoGrupoPermissaoBuilder) WithPermissao(permissao *Permissao) *PermissaoGrupoPermissaoBuilder {
	b.permissaoGrupoPermissao.Permissao = permissao
	return b
}

func (b *PermissaoGrupoPermissaoBuilder) WithGrupoPermissao(grupoPermissao *GrupoPermissao) *PermissaoGrupoPermissaoBuilder {
	b.permissaoGrupoPermissao.GrupoPermissao = grupoPermissao
	return b
}

func (b *PermissaoGrupoPermissaoBuilder) Build() *PermissaoGrupoPermissao {
	return b.permissaoGrupoPermissao
}
