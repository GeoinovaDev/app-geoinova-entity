package entity

import "time"

type IncendioBuilder struct {
	incendio *Incendio
}

func NewIncendioBuilder(id uint) *IncendioBuilder {
	return &IncendioBuilder{
		incendio: NewIncendio(id),
	}
}

func (b *IncendioBuilder) WithWkt(wkt string) *IncendioBuilder {
	b.incendio.Wkt = wkt
	return b
}

func (b *IncendioBuilder) WithAreaHa(areaHa float64) *IncendioBuilder {
	b.incendio.AreaHa = areaHa
	return b
}

func (b *IncendioBuilder) WithCamada(camada *Camada) *IncendioBuilder {
	b.incendio.Camada = camada
	return b
}

func (b *IncendioBuilder) WithDateTime(dateTime *time.Time) *IncendioBuilder {
	b.incendio.DateTime = dateTime
	return b
}

func (b *IncendioBuilder) WithTitulo(titulo string) *IncendioBuilder {
	b.incendio.Titulo = titulo
	return b
}

func (b *IncendioBuilder) WithCriticidade(criticidade string) *IncendioBuilder {
	b.incendio.Criticidade = criticidade
	return b
}

func (b *IncendioBuilder) WithConfianca(confianca string) *IncendioBuilder {
	b.incendio.Confianca = confianca
	return b
}

func (b *IncendioBuilder) WithDescricao(descricao string) *IncendioBuilder {
	b.incendio.Descricao = descricao
	return b
}

func (b *IncendioBuilder) WithTemperatura(temperatura int) *IncendioBuilder {
	b.incendio.Temperatura = temperatura
	return b
}

func (b *IncendioBuilder) WithTotalFocos(totalFocos uint) *IncendioBuilder {
	b.incendio.TotalFocos = totalFocos
	return b
}

func (b *IncendioBuilder) WithSatelites(satelites []string) *IncendioBuilder {
	b.incendio.Satelites = satelites
	return b
}

func (b *IncendioBuilder) Build() *Incendio {
	return b.incendio
}
