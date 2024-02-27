package entity

type ClienteStatus string

const (
	CLIENTE_STATUS_ATIVO   = ClienteStatus("A")
	CLIENTE_STATUS_INATIVO = ClienteStatus("I")
)

func (cs ClienteStatus) Equals(status ClienteStatus) bool {
	return string(cs) == string(status)
}

func NewClienteStatus(status string) ClienteStatus {
	return ClienteStatus(status)
}
