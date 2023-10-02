package entity

type ImagemQualidade struct {
	Id         uint
	Nome       string
	Color      string
	Background string
	Tipo       uint
}

func NewImagemQualidade(id uint) *ImagemQualidade {
	return &ImagemQualidade{Id: id}
}
