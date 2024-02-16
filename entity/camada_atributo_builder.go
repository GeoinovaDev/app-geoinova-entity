package entity

type CamadaAtributoBuilder struct {
	camadaAtributo *CamadaAtributo
}

func NewCamadaAtributoBuilder(id uint) *CamadaAtributoBuilder {
	return &CamadaAtributoBuilder{
		camadaAtributo: NewCamadaAtributo(id),
	}
}

func (b *CamadaAtributoBuilder) WithNome(nome string) *CamadaAtributoBuilder {
	b.camadaAtributo.Nome = nome
	return b
}

func (b *CamadaAtributoBuilder) WithValor(valor string) *CamadaAtributoBuilder {
	b.camadaAtributo.Valor = valor
	return b
}

func (b *CamadaAtributoBuilder) WithCamada(camada *Camada) *CamadaAtributoBuilder {
	b.camadaAtributo.Camada = camada
	return b
}

func (b *CamadaAtributoBuilder) Build() *CamadaAtributo {
	return b.camadaAtributo
}
