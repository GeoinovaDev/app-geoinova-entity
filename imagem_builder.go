package entity

type ImagemBuilder struct {
	imagem *Imagem
}

func NewImagemBuilder(id uint) ImagemBuilder {
	return ImagemBuilder{NewImagem(id)}
}

func (b ImagemBuilder) WithTotalBanda(totalBandas uint) ImagemBuilder {
	b.imagem.TotalBandas = totalBandas
	return b
}

func (b ImagemBuilder) WithUuid(uuid string) ImagemBuilder {
	b.imagem.Uuid = uuid
	return b
}

func (b ImagemBuilder) WithData(data string) ImagemBuilder {
	b.imagem.Data = data
	return b
}

func (b ImagemBuilder) WithBucket(bucket string) ImagemBuilder {
	b.imagem.Bucket = bucket
	return b
}

func (b ImagemBuilder) WithResolucao(resolucao string) ImagemBuilder {
	b.imagem.Resolucao = resolucao
	return b
}

func (b ImagemBuilder) WithProvedor(provedor string) ImagemBuilder {
	b.imagem.Provedor = provedor
	return b
}

func (b ImagemBuilder) WithSource(source ImagemSource) ImagemBuilder {
	b.imagem.Source = source
	return b
}

func (b ImagemBuilder) WithTaxaCoberturaNuvem(tx float32) ImagemBuilder {
	b.imagem.TaxaCoberturaNuvem = tx
	return b
}

func (b ImagemBuilder) WithProvedorFonte(fonte string) ImagemBuilder {
	b.imagem.ProvedorFonte = fonte
	return b
}

func (b ImagemBuilder) WithQualidade(qualidade *ImagemQualidade) ImagemBuilder {
	b.imagem.Qualidade = qualidade
	return b
}

func (b ImagemBuilder) WithAtivo(ativo *Ativo) ImagemBuilder {
	b.imagem.Ativo = ativo
	return b
}

func (b ImagemBuilder) WithUsuariosFavoritagem(usuarios []*Usuario) ImagemBuilder {
	b.imagem.UsuarioFavoritos = usuarios
	return b
}

func (b ImagemBuilder) AddUsuarioFavorito(usuario *Usuario) ImagemBuilder {
	b.imagem.UsuarioFavoritos = append(b.imagem.UsuarioFavoritos, usuario)
	return b
}

func (b ImagemBuilder) Build() *Imagem {
	return b.imagem
}
