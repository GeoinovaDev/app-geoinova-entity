package entity

type PermissaoBuilder struct {
	permissao *Permissao
}

func NewPermissaoBuilder(id uint) *PermissaoBuilder {
	return &PermissaoBuilder{
		permissao: NewPermissao(id),
	}
}

func (b *PermissaoBuilder) WithNome(nome string) *PermissaoBuilder {
	b.permissao.Nome = nome
	return b
}

func (b *PermissaoBuilder) WithAction(action string) *PermissaoBuilder {
	b.permissao.Action = action
	return b
}

func (b *PermissaoBuilder) WithLeitura(leitura bool) *PermissaoBuilder {
	b.permissao.Leitura = leitura
	return b
}

func (b *PermissaoBuilder) WithEscrita(escrita bool) *PermissaoBuilder {
	b.permissao.Escrita = escrita
	return b
}

func (b *PermissaoBuilder) WithExclusao(exclusao bool) *PermissaoBuilder {
	b.permissao.Exclusao = exclusao
	return b
}

func (b *PermissaoBuilder) Build() *Permissao {
	return b.permissao
}
