package entity

type AlertaEnvioNotificacaoBuider struct {
	notificacao *AlertaEnvioNotificacao
}

func NewAlertaNotificacaoBuilder() AlertaEnvioNotificacaoBuider {
	return AlertaEnvioNotificacaoBuider{
		notificacao: NewAlertaNotificacao(),
	}
}

func (b AlertaEnvioNotificacaoBuider) WithAtivoNome(ativoNome string) AlertaEnvioNotificacaoBuider {
	b.notificacao.AtivoNome = ativoNome
	return b
}
func (b AlertaEnvioNotificacaoBuider) WithDataImagemAntes(dataImagemAntes string) AlertaEnvioNotificacaoBuider {
	b.notificacao.DataImagemAntes = dataImagemAntes
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithDataImagemDepois(dataImagemDepois string) AlertaEnvioNotificacaoBuider {
	b.notificacao.DataImagemDepois = dataImagemDepois
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithDeteccaoClasseNome(deteccaoClasseNome string) AlertaEnvioNotificacaoBuider {
	b.notificacao.DeteccaoClasseNome = deteccaoClasseNome
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithDeteccaoDescricao(deteccaoDescricao string) AlertaEnvioNotificacaoBuider {
	b.notificacao.DeteccaoDescricao = deteccaoDescricao
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithAnaliseAutomatica(analiseAutomatica bool) AlertaEnvioNotificacaoBuider {
	b.notificacao.AnaliseAutomatica = analiseAutomatica
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithPosicao(posicao uint) AlertaEnvioNotificacaoBuider {
	b.notificacao.Posicao = posicao
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithEmails(emails []*EnderecoEmail) AlertaEnvioNotificacaoBuider {
	b.notificacao.Emails = emails
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithTelefones(telefones []*Telefone) AlertaEnvioNotificacaoBuider {
	b.notificacao.Telefones = telefones
	return b
}

func (b AlertaEnvioNotificacaoBuider) WithParametros(parametros map[string]string) AlertaEnvioNotificacaoBuider {
	b.notificacao.Parametros = parametros
	return b
}

func (b AlertaEnvioNotificacaoBuider) Build() *AlertaEnvioNotificacao {
	return b.notificacao
}
