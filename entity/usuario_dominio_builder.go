package entity

type UsuarioDominioBuilder struct {
	usuarioDominio *UsuarioDominio
}

func NewUsuarioDominioBuilder(id uint) *UsuarioDominioBuilder {
	return &UsuarioDominioBuilder{
		usuarioDominio: NewUsuarioDominio(id),
	}
}

func (b *UsuarioDominioBuilder) WithDominio(dominio string) *UsuarioDominioBuilder {
	b.usuarioDominio.Dominio = dominio
	return b
}

func (b *UsuarioDominioBuilder) WithChecked(checked bool) *UsuarioDominioBuilder {
	b.usuarioDominio.Checked = checked
	return b
}

func (b *UsuarioDominioBuilder) WithRole(role string) *UsuarioDominioBuilder {
	b.usuarioDominio.Role = role
	return b
}

func (b *UsuarioDominioBuilder) WithUsuario(usuario *Usuario) *UsuarioDominioBuilder {
	b.usuarioDominio.Usuario = usuario
	return b
}

func (b *UsuarioDominioBuilder) Build() *UsuarioDominio {
	return b.usuarioDominio
}
