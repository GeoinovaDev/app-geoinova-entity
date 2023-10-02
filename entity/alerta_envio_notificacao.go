package entity

type AlertaEnvioNotificacao struct {
	Uuid               string
	AtivoNome          string
	DataImagemAntes    string
	DataImagemDepois   string
	DeteccaoClasseNome string
	DeteccaoDescricao  string
	AnaliseAutomatica  bool
	Posicao            uint
	Emails             []*EnderecoEmail
	Telefones          []*Telefone
	Parametros         map[string]string
}

func NewAlertaNotificacao() *AlertaEnvioNotificacao {
	return &AlertaEnvioNotificacao{}
}
