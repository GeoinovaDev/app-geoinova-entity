package entity

type ResumoRelatorioDeteccao struct {
	TotalAnalises        uint
	TotalDeteccoes       uint
	TotalArea            float32
	TotalDeteccaoPorTipo map[DeteccaoClasse]uint
	TotalDeteccaoPorArea map[DeteccaoClasse]float32
}

func NewResumoRelatorioDeteccao(
	totalAnalises uint,
	totalDeteccoes uint,
	totalArea float32,
	totalTipoPorDeteccao map[DeteccaoClasse]uint,
	totalAreaPorDeteccao map[DeteccaoClasse]float32,
) *ResumoRelatorioDeteccao {
	return &ResumoRelatorioDeteccao{
		TotalAnalises:        totalAnalises,
		TotalDeteccoes:       totalDeteccoes,
		TotalArea:            totalArea,
		TotalDeteccaoPorTipo: totalTipoPorDeteccao,
		TotalDeteccaoPorArea: totalAreaPorDeteccao,
	}
}
