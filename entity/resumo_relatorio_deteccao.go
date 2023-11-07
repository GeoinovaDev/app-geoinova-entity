package entity

type ResumoRelatorioDeteccaoClasseTotal struct {
	Classe *DeteccaoClasse
	Total  uint
}

type ResumoRelatorioDeteccaoClasseArea struct {
	Classe *DeteccaoClasse
	Total  float32
}

type ResumoRelatorioDeteccao struct {
	TotalAnalises  uint
	TotalDeteccoes uint
	TotalArea      float32
	ClassesTotal   []*ResumoRelatorioDeteccaoClasseTotal
	ClassesArea    []*ResumoRelatorioDeteccaoClasseArea
}

func NewResumoRelatorioDeteccaoClasseTotal(classe *DeteccaoClasse, total uint) *ResumoRelatorioDeteccaoClasseTotal {
	return &ResumoRelatorioDeteccaoClasseTotal{
		Classe: classe,
		Total:  total,
	}
}

func NewResumoRelatorioDeteccaoClasseArea(classe *DeteccaoClasse, total float32) *ResumoRelatorioDeteccaoClasseArea {
	return &ResumoRelatorioDeteccaoClasseArea{
		Classe: classe,
		Total:  total,
	}
}

func NewResumoRelatorioDeteccao(
	totalAnalises uint,
	totalDeteccoes uint,
	totalArea float32,
	classesTotal []*ResumoRelatorioDeteccaoClasseTotal,
	classesArea []*ResumoRelatorioDeteccaoClasseArea,
) *ResumoRelatorioDeteccao {
	return &ResumoRelatorioDeteccao{
		TotalAnalises:  totalAnalises,
		TotalDeteccoes: totalDeteccoes,
		TotalArea:      totalArea,
		ClassesTotal:   classesTotal,
		ClassesArea:    classesArea,
	}
}
