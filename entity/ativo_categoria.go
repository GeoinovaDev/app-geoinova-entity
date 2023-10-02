package entity

type AtivoCategoria struct {
	Id   uint
	Nome string
}

func NewAtivoCategoria(id uint, nome string) *AtivoCategoria {
	return &AtivoCategoria{id, nome}
}
