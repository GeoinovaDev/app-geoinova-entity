package entity

type UsuarioStatus string

const (
	USUARIO_STATUS_ATIVO   = UsuarioStatus("A")
	USUARIO_STATUS_INATIVO = UsuarioStatus("I")
)

func (u UsuarioStatus) Equals(status UsuarioStatus) bool {
	return string(u) == string(status)
}

func NewUsuarioStatus(status string) UsuarioStatus {
	return UsuarioStatus(status)
}
