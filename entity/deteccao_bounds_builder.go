package entity

type DeteccaoBoundsBuilder struct {
	bounds *DeteccaoBounds
}

func NewDeteccaoBoundsBuilder() *DeteccaoBoundsBuilder {
	return &DeteccaoBoundsBuilder{
		bounds: NewDeteccaoBounds(),
	}
}

func (b *DeteccaoBoundsBuilder) WithAntes(antes []float32) *DeteccaoBoundsBuilder {
	b.bounds.Antes = antes
	return b
}

func (b *DeteccaoBoundsBuilder) WithDepois(depois []float32) *DeteccaoBoundsBuilder {
	b.bounds.Depois = depois
	return b
}

func (b *DeteccaoBoundsBuilder) WithRegiao(regiao []float32) *DeteccaoBoundsBuilder {
	b.bounds.Regiao = regiao
	return b
}

func (b *DeteccaoBoundsBuilder) Build() *DeteccaoBounds {
	return b.bounds
}
