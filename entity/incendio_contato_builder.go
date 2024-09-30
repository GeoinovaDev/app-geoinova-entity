package entity

type IncendioContatoBuilder struct {
	IncendioContato *IncendioContato
}

func NewIncendioContatoBuilder() *IncendioContatoBuilder {
	return &IncendioContatoBuilder{
		IncendioContato: NewIncendioContato(),
	}
}

func (b *IncendioContatoBuilder) WithEmail(emails ...string) *IncendioContatoBuilder {
	b.IncendioContato.Emails = emails
	return b
}

func (b *IncendioContatoBuilder) WithTelefone(telefones ...*Telefone) *IncendioContatoBuilder {
	b.IncendioContato.Telefones = telefones
	return b
}

func (b *IncendioContatoBuilder) WithNumeroTelefone(numero ...string) *IncendioContatoBuilder {

	for _, n := range numero {
		b.IncendioContato.Telefones = append(b.IncendioContato.Telefones, ParseTelefone(n))
	}

	return b
}

func (b *IncendioContatoBuilder) Build() *IncendioContato {
	return b.IncendioContato
}
