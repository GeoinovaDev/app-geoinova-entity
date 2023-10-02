package entity

type DeteccaoBuilder struct {
	deteccao *Deteccao
}

func NewDeteccaoBuilder(id uint) DeteccaoBuilder {
	return DeteccaoBuilder{NewDeteccao(id)}
}

func (b DeteccaoBuilder) WithWkt(wkt string) DeteccaoBuilder {
	b.deteccao.Wkt = wkt
	return b
}

func (b DeteccaoBuilder) WithClasse(classe *DeteccaoClasse) DeteccaoBuilder {
	b.deteccao.Classe = classe
	return b
}

func (b DeteccaoBuilder) WithDescricao(descricao string) DeteccaoBuilder {
	b.deteccao.Descricao = descricao
	return b
}

func (b DeteccaoBuilder) WithImagemAntesUuid(uuid string) DeteccaoBuilder {
	b.deteccao.ImagemAntesUuid = uuid
	return b
}

func (b DeteccaoBuilder) WithImagemDepoisUuid(uuid string) DeteccaoBuilder {
	b.deteccao.ImagemDepoisUuid = uuid
	return b
}

func (b DeteccaoBuilder) WithImagemResultadoUuid(uuid string) DeteccaoBuilder {
	b.deteccao.ImagemResultadoUuid = uuid
	return b
}

func (b DeteccaoBuilder) WithCamadas(camadas []*Camada) DeteccaoBuilder {
	b.deteccao.Camadas = camadas
	return b
}

func (b DeteccaoBuilder) AddCamada(camada *Camada) DeteccaoBuilder {
	b.deteccao.Camadas = append(b.deteccao.Camadas, camada)
	return b
}

func (b DeteccaoBuilder) Build() *Deteccao {
	return b.deteccao
}
