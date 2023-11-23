package entity

type UsuarioDominio struct {
	Id      uint
	Dominio string
	Role    string
	Usuario *Usuario
	Checked bool
}

type UsuariosDominios []*UsuarioDominio

func NewUsuarioDominio(id uint) *UsuarioDominio {
	return &UsuarioDominio{
		Id: id,
	}
}
