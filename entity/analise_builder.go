package entity

type AnaliseBuilder struct {
	analise *Analise
}

func NewAnaliseBuilder(id uint) AnaliseBuilder {
	return AnaliseBuilder{NewAnalise(id)}
}

func (b AnaliseBuilder) WithTotalDeteccoes(totalDeteccoes uint) AnaliseBuilder {
	b.analise.TotalDeteccoes = totalDeteccoes
	return b
}

func (b AnaliseBuilder) WithWkt(wkt string) AnaliseBuilder {
	b.analise.Wkt = wkt
	return b
}

func (b AnaliseBuilder) WithUuid(uuid string) AnaliseBuilder {
	b.analise.Uuid = uuid
	return b
}

func (b AnaliseBuilder) WithUsuario(usuario *Usuario) AnaliseBuilder {
	b.analise.Usuario = usuario
	return b
}

func (b AnaliseBuilder) WithStatus(status AnaliseStatus) AnaliseBuilder {
	b.analise.Status = status
	return b
}

func (b AnaliseBuilder) WithTotalArea(totalArea float32) AnaliseBuilder {
	b.analise.TotalArea = totalArea
	return b
}

func (b AnaliseBuilder) WithTipo(tipo AnaliseTipo) AnaliseBuilder {
	b.analise.Tipo = tipo
	return b
}

func (b AnaliseBuilder) WithImagemAntes(imagemAntes *Imagem) AnaliseBuilder {
	b.analise.ImagemAntes = imagemAntes
	return b
}

func (b AnaliseBuilder) WithImagemDepois(imagemDepois *Imagem) AnaliseBuilder {
	b.analise.ImagemDepois = imagemDepois
	return b
}

func (b AnaliseBuilder) WithAtivo(ativo *Ativo) AnaliseBuilder {
	b.analise.Ativo = ativo
	return b
}

func (b AnaliseBuilder) WithCamadas(camadas []*AnaliseCamada) AnaliseBuilder {
	b.analise.Camadas = camadas
	return b
}

func (b AnaliseBuilder) AddCamada(camada *AnaliseCamada) AnaliseBuilder {
	b.analise.Camadas = append(b.analise.Camadas, camada)
	return b
}

func (b AnaliseBuilder) WithAlertas(alertas []*Alerta) AnaliseBuilder {
	b.analise.Alertas = alertas
	return b
}

func (b AnaliseBuilder) AddAlerta(alerta *Alerta) AnaliseBuilder {
	b.analise.Alertas = append(b.analise.Alertas, alerta)
	return b
}

func (b AnaliseBuilder) Build() *Analise {
	return b.analise
}
