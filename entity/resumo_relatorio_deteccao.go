package entity

type ResumoRelatorioDeteccao struct {
	TotalAnalises        uint
	TotalDeteccoes       uint
	TotalTipoPorDeteccao map[DeteccaoClasse]uint
	TotalAreaPorDeteccao map[DeteccaoClasse]float32
}

func NewResumoRelatorioDeteccao(
	totalAnalises uint,
	totalDeteccoes uint,
	totalTipoPorDeteccao map[DeteccaoClasse]uint,
	totalAreaPorDeteccao map[DeteccaoClasse]float32,
) *ResumoRelatorioDeteccao {
	return &ResumoRelatorioDeteccao{
		TotalAnalises:        totalAnalises,
		TotalDeteccoes:       totalDeteccoes,
		TotalTipoPorDeteccao: totalTipoPorDeteccao,
		TotalAreaPorDeteccao: totalAreaPorDeteccao,
	}
}
