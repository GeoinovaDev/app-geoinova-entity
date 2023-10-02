package entity

type AlertaBuilder struct {
	alerta *Alerta
}

func NewAlertaBuilder(id uint) AlertaBuilder {
	return AlertaBuilder{NewAlerta(id)}
}

func (b AlertaBuilder) WithPosicao(order uint) AlertaBuilder {
	b.alerta.Posicao = order
	return b
}

func (b AlertaBuilder) WithFases(fases []*AlertaFase) AlertaBuilder {
	b.alerta.Fases = fases
	return b
}

func (b AlertaBuilder) AddFase(fase *AlertaFase) AlertaBuilder {
	b.alerta.Fases = append(b.alerta.Fases, fase)
	return b
}

func (b AlertaBuilder) WithEventos(eventos []*AlertaEvento) AlertaBuilder {
	b.alerta.Eventos = eventos
	return b
}

func (b AlertaBuilder) AddEvento(evento *AlertaEvento) AlertaBuilder {
	b.alerta.Eventos = append(b.alerta.Eventos, evento)
	return b
}

func (b AlertaBuilder) WithObservacao(obs string) AlertaBuilder {
	b.alerta.Observacao = obs
	return b
}

func (b AlertaBuilder) WithResponsavel(responsavel string) AlertaBuilder {
	b.alerta.Responsavel = responsavel
	return b
}

func (b AlertaBuilder) WithJustificativa(justificativa string) AlertaBuilder {
	b.alerta.Justificativa = justificativa
	return b
}

func (b AlertaBuilder) WithStatus(status AlertaFaseStatus) AlertaBuilder {
	b.alerta.Status = status
	return b
}

func (b AlertaBuilder) WithDeteccao(deteccao *Deteccao) AlertaBuilder {
	b.alerta.Deteccao = deteccao
	return b
}

func (b AlertaBuilder) WithAnalise(analise *Analise) AlertaBuilder {
	b.alerta.Analise = analise
	return b
}

func (b AlertaBuilder) Build() *Alerta {
	return b.alerta
}
