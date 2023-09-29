package entity

type ImagemFavoritada struct {
	Id      uint
	Usuario *Usuario
	Imagem  *Imagem
}

func NewImagemFavoritada(id uint, usuario *Usuario, imagem *Imagem) *ImagemFavoritada {
	return &ImagemFavoritada{id, usuario, imagem}
}
