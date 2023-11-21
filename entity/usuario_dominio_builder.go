package entity

type UsuarioDominioBuilder struct {
	usuarioDominio *UsuarioDominio
}

func NewUsuarioDominioBuilder() *UsuarioDominioBuilder {
	return &UsuarioDominioBuilder{
		usuarioDominio: &UsuarioDominio{},
	}
}

func (b *UsuarioDominioBuilder) WithDominio(dominio string) *UsuarioDominioBuilder {
	b.usuarioDominio.Dominio = dominio
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
