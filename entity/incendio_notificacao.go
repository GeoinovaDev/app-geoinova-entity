package entity

type IncendioNotificacao struct {
	Id        uint
	Incendio  *Incendio
	Emails    []string
	Telefones []string
}

func NewIncendioNotificacao(id uint) *IncendioNotificacao {
	return &IncendioNotificacao{
		Id: id,
	}
}
