package entity

type DeteccaoBounds struct {
	Antes  []float32
	Depois []float32
	Regiao []float32
}

func NewDeteccaoBounds() *DeteccaoBounds {
	return &DeteccaoBounds{
		Antes:  []float32{},
		Depois: []float32{},
		Regiao: []float32{},
	}
}
