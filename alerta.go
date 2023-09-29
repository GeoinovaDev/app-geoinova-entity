package entity

type Alerta struct {
	Id            uint
	Posicao       uint
	Observacao    string
	Responsavel   string
	Justificativa string
	Status        AlertaFaseStatus
	Deteccao      *Deteccao
	Analise       *Analise
	Eventos       []*AlertaEvento
	Fases         []*AlertaFase
}

func NewAlerta(id uint) *Alerta {
	return &Alerta{
		Id:      id,
		Status:  ALERTA_FASE_ABERTA,
		Eventos: []*AlertaEvento{},
		Fases:   []*AlertaFase{},
	}
}
