package entity

type CamadaBuilder struct {
	camada *Camada
}

func NewCamadaBuilder(id uint) CamadaBuilder {
	return CamadaBuilder{NewCamada(id)}
}

func (b CamadaBuilder) WithNome(nome string) CamadaBuilder {
	b.camada.Nome = nome
	return b
}

func (b CamadaBuilder) WithDetalhe(detalhe string) CamadaBuilder {
	b.camada.Detalhe = detalhe
	return b
}

func (b CamadaBuilder) WithWkt(wkt string) CamadaBuilder {
	b.camada.Wkt = wkt
	return b
}

func (b CamadaBuilder) WithAtivo(ativo *Ativo) CamadaBuilder {
	b.camada.Ativo = ativo
	return b
}

func (b CamadaBuilder) WithArea(area float32) CamadaBuilder {
	b.camada.Area = area
	return b
}

func (b CamadaBuilder) WithCategoria(categoria *CamadaCategoria) CamadaBuilder {
	b.camada.Categoria = categoria
	return b
}

func (b CamadaBuilder) WithAtributos(atributos []*CamadaAtributo) CamadaBuilder {
	b.camada.Atributos = atributos
	return b
}

func (b CamadaBuilder) AddAtributo(atributo ...*CamadaAtributo) CamadaBuilder {
	b.camada.Atributos = append(b.camada.Atributos, atributo...)
	return b
}

func (b CamadaBuilder) Build() *Camada {
	return b.camada
}
