package entity

type EnderecoEmail struct {
	Nome  string
	Email string
}

func NewEnderecoEmail(nome string, email string) *EnderecoEmail {
	return &EnderecoEmail{nome, email}
}
