package entity

type PermissaoGrupoPermissao struct {
	Id        uint
	Permissao *Permissao
	Enabled   bool
}

func NewPermissaoGrupoPermissao(id uint) *PermissaoGrupoPermissao {
	return &PermissaoGrupoPermissao{
		Id: id,
	}
}
