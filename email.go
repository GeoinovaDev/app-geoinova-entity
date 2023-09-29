package entity

type Email struct {
	De       *EnderecoEmail
	Para     *EnderecoEmail
	Assunto  string
	Conteudo string
}

func NewEmail(de *EnderecoEmail, para *EnderecoEmail, assunto string, conteudo string) *Email {
	return &Email{de, para, assunto, conteudo}
}
