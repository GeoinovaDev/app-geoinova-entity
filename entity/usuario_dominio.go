package entity

const (
	USUARIO_ROLE_ADMIN     = "admin"
	USUARIO_ROLE_USER      = "user"
	USUARIO_ROLE_MODERATOR = "moderador"
	USUARIO_ROLE_TESTER    = "tester"
)

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

func (u UsuariosDominios) ExistDominio(dominio string) bool {
	for _, _dominio := range u {
		if _dominio.Dominio == dominio {
			return true
		}
	}

	return false
}
