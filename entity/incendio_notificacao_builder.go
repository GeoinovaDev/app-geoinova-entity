package entity

type IncendioNotificacaoBuilder struct {
	incendioNotificacao *IncendioNotificacao
}

func NewIncendioNotificacaoBuilder(id uint) *IncendioNotificacaoBuilder {
	return &IncendioNotificacaoBuilder{
		incendioNotificacao: NewIncendioNotificacao(id),
	}
}

func (b *IncendioNotificacaoBuilder) WithIncendio(incendio *Incendio) *IncendioNotificacaoBuilder {
	b.incendioNotificacao.Incendio = incendio
	return b
}

func (b *IncendioNotificacaoBuilder) WithEmail(email ...string) *IncendioNotificacaoBuilder {
	b.incendioNotificacao.Emails = append(b.incendioNotificacao.Emails, email...)
	return b
}

func (b *IncendioNotificacaoBuilder) WithTelefone(telefone ...string) *IncendioNotificacaoBuilder {
	b.incendioNotificacao.Telefones = append(b.incendioNotificacao.Telefones, telefone...)
	return b
}

func (b *IncendioNotificacaoBuilder) Build() *IncendioNotificacao {
	return b.incendioNotificacao
}
