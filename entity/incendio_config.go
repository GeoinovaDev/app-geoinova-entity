package entity

type IncendioConfig struct {
	Id       uint
	Cliente  *Cliente
	Nome     string
	Raio     float64
	Camadas  []*Camada
	Usuarios []*Usuario
	Contatos []*IncendioContato
}

func NewIncendioConfig(id uint) *IncendioConfig {
	return &IncendioConfig{
		Id: id,
	}
}
