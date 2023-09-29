package entity

type Usuario struct {
	Id      uint
	Nome    string
	Email   string
	Senha   string
	Status  UsuarioStatus
	Cliente *Cliente
}

func NewUsuario(id uint) *Usuario {
	return &Usuario{
		Id: id,
	}
}
