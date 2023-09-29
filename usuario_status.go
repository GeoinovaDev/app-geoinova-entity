package entity

type UsuarioStatus string

func (u UsuarioStatus) Equals(status UsuarioStatus) bool {
	return string(u) == string(status)
}

const (
	USUARIO_STATUS_ATIVO   = UsuarioStatus("A")
	USUARIO_STATUS_INATIVO = UsuarioStatus("I")
)
