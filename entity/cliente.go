package entity

type Cliente struct {
	Id        uint
	Nome      string
	Descricao string
	Status    ClienteStatus
	Usuarios  []*Usuario
	Modulos   []*Modulo
}

func NewCliente(id uint) *Cliente {
	return &Cliente{
		Id:       id,
		Status:   CLIENTE_STATUS_ATIVO,
		Usuarios: []*Usuario{},
		Modulos:  []*Modulo{},
	}
}

func (b *Cliente) AddModulo(modulo *Modulo) {
	b.Modulos = append(b.Modulos, modulo)
}

func (b *Cliente) RemoveModulo(modulo *Modulo) {
	for i, m := range b.Modulos {
		if m.Id == modulo.Id {
			b.Modulos = append(b.Modulos[:i], b.Modulos[i+1:]...)
		}
	}
}

func (b *Cliente) ExistModulo(modulo *Modulo) bool {
	for _, m := range b.Modulos {
		if m.Id == modulo.Id {
			return true
		}

		if m.Nome == modulo.Nome {
			return true
		}
	}

	return false
}

func (b *Cliente) ClearModulos() {
	b.Modulos = []*Modulo{}
}
