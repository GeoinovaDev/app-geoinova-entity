package entity

type Usuario struct {
	Id              uint
	Nome            string
	Email           string
	Senha           string
	Status          UsuarioStatus
	Cliente         *Cliente
	PermissaoAcesso []*Ativo
	GrupoPermissao  *GrupoPermissao
}

func NewUsuario(id uint) *Usuario {
	return &Usuario{
		Id:              id,
		PermissaoAcesso: []*Ativo{},
	}
}
