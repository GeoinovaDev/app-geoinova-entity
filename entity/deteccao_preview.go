package entity

type DeteccaoPreview struct {
	Nome   string
	Tipo   string
	Base64 string
	Path   string
	Bucket string
}

func NewDeteccaoPreview() *DeteccaoPreview {
	return &DeteccaoPreview{}
}