package entity

type UsuarioPapel string

const (
	USUARIO_PAPEL_ADMIN = "admin"
	USUARIO_PAPEL_USER  = "user"
)

func (u UsuarioPapel) isEquals(papel UsuarioPapel) bool {
	return string(u) == string(papel)
}
