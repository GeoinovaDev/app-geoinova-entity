package entity

const (
	USUARIO_ROLE_ADMIN     = "admin"
	USUARIO_ROLE_USER      = "user"
	USUARIO_ROLE_MODERATOR = "moderador"
	USUARIO_ROLE_TESTER    = "tester"
	USUARIO_ROLE_ROOT      = "root"
)

const (
	USUARIO_DOMINIO_APP        = "app"
	USUARIO_DOMINIO_PAINEL     = "painel"
	USUARIO_DOMINIO_PREVIEW    = "preview"
	USUARIO_DOMINIO_INSIDERS   = "insiders"
	USUARIO_DOMINIO_DASHIBOARD = "dashbird"
)

type UsuarioDominio struct {
	Id         uint
	Dominio    string
	Role       string
	Usuario    *Usuario
	Checked    bool
	AppVersion string
	ApiVersion string
}

type UsuariosDominios []*UsuarioDominio

func NewUsuarioDominio(id uint) *UsuarioDominio {
	return &UsuarioDominio{
		Id: id,
	}
}

func (u UsuariosDominios) ExistRole(role string) bool {
	for _, _dominio := range u {
		if _dominio.Role == role {
			return true
		}
	}

	return false
}

func (u UsuariosDominios) ExistDominioAndRole(dominio string, role string) bool {
	for _, _dominio := range u {
		if _dominio.Role == role && _dominio.Dominio == dominio {
			return true
		}
	}

	return false
}

func (u UsuariosDominios) ExistDominio(dominio string) bool {
	for _, _dominio := range u {
		if _dominio.Dominio == dominio {
			return true
		}
	}

	return false
}

func (u UsuariosDominios) GetDominioDefault() string {
	for _, _dominio := range u {
		if _dominio.Checked {
			return _dominio.Dominio
		}
	}

	if len(u) > 0 {
		return u[0].Dominio
	}

	return ""
}
