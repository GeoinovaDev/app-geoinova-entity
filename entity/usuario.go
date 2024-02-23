package entity

type Usuario struct {
	Id             uint
	Nome           string
	Email          string
	Senha          string
	Status         UsuarioStatus
	Cliente        *Cliente
	Permissoes     *Permissoes
	GrupoPermissao *GrupoPermissao
}

func NewUsuario(id uint) *Usuario {
	return &Usuario{
		Id:         id,
		Permissoes: &Permissoes{},
	}
}
