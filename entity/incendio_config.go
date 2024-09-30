package entity

type IncendioConfig struct {
	Id       uint
	Cliente  *Cliente
	Nome     string
	Raio     float64
	Contatos *IncendioContato
	Camadas  []*Camada
	Usuarios []*Usuario
}

func NewIncendioConfig(id uint) *IncendioConfig {
	return &IncendioConfig{
		Id: id,
	}
}
