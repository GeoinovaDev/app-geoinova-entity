package entity

type PermissaoGrupoPermissao struct {
	Id             uint
	Permissao      *Permissao
	GrupoPermissao *GrupoPermissao
	Enabled        bool
}

func NewPermissaoGrupoPermissao(id uint) *PermissaoGrupoPermissao {
	return &PermissaoGrupoPermissao{
		Id: id,
	}
}
