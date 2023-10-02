package entity

type Cliente struct {
	Id        uint
	Nome      string
	Descricao string
	Status    ClienteStatus
	Usuarios  []*Usuario
}

func NewCliente(id uint) *Cliente {
	return &Cliente{
		Id:       id,
		Status:   CLIENTE_STATUS_ATIVO,
		Usuarios: []*Usuario{},
	}
}
