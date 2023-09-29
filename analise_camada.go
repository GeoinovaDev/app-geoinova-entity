package entity

type AnaliseCamada struct {
	Id     uint    `json:"id"`
	Camada *Camada `json:"camada"`
	Total  uint    `json:"total"`
}

func NewAnaliseCamada(id uint, camada *Camada, total uint) *AnaliseCamada {
	return &AnaliseCamada{id, camada, total}
}
