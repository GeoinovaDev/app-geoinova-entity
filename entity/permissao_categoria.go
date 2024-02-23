package entity

type PermissaoCategoria struct {
	Id   uint
	Nome string
}

func NewPermissaoCategoria(id uint) *PermissaoCategoria {
	return &PermissaoCategoria{
		Id: id,
	}
}
