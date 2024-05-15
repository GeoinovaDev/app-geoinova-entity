package entity

type PermissaoGrupoPermissao struct {
	Id             uint
	Permissao      *Permissao
	GrupoPermissao *GrupoPermissao
	Cliente        *Cliente
}

func NewPermissaoGrupoPermissao(id uint) *PermissaoGrupoPermissao {
	return &PermissaoGrupoPermissao{
		Id: id,
	}
}
