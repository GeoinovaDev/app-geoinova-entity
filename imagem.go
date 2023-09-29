package entity

type ImagemSource uint

const (
	IMAGEM_SOURCE_UPLOAD   = ImagemSource(0)
	IMAGEM_SOURCE_ANALISE  = ImagemSource(1)
	IMAGEM_SOURCE_FAVORITO = ImagemSource(2)
)

type Imagem struct {
	Id                 uint
	Uuid               string
	Data               string
	Bucket             string
	Resolucao          string
	TotalBandas        uint
	TaxaCoberturaNuvem float32
	Provedor           string
	ProvedorFonte      string
	Qualidade          *ImagemQualidade
	Ativo              *Ativo
	Source             ImagemSource
	UsuarioFavoritos   []*Usuario
}

func NewImagem(id uint) *Imagem {
	return &Imagem{
		Id:               id,
		Source:           IMAGEM_SOURCE_UPLOAD,
		UsuarioFavoritos: []*Usuario{},
	}
}

func (im *Imagem) AddUsuarioFavorito(usuario *Usuario) {
	im.UsuarioFavoritos = append(im.UsuarioFavoritos, usuario)
}
