package entity

type DeteccaoPreviews struct {
	Antes  *DeteccaoPreview
	Depois *DeteccaoPreview
	Regiao *DeteccaoPreview
}

func NewDeteccaoPreviews() *DeteccaoPreviews {
	return &DeteccaoPreviews{
		Antes:  NewDeteccaoPreview(),
		Depois: NewDeteccaoPreview(),
		Regiao: NewDeteccaoPreview(),
	}
}
