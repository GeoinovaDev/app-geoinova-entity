package entity

type UsuarioBuilder struct {
	usuario *Usuario
}

func NewUsuarioBuilder(id uint) UsuarioBuilder {
	return UsuarioBuilder{NewUsuario(id)}
}

func (b UsuarioBuilder) WithNome(nome string) UsuarioBuilder {
	b.usuario.Nome = nome
	return b
}

func (b UsuarioBuilder) AddPermissao(permissao *Permissao) UsuarioBuilder {
	b.usuario.Permissoes.Add(permissao)
	return b
}

func (b UsuarioBuilder) WithGrupoPermissao(grupo *GrupoPermissao) UsuarioBuilder {
	b.usuario.GrupoPermissao = grupo
	return b
}

func (b UsuarioBuilder) WithEmail(email string) UsuarioBuilder {
	b.usuario.Email = email
	return b
}

func (b UsuarioBuilder) WithStatus(status UsuarioStatus) UsuarioBuilder {
	b.usuario.Status = status
	return b
}

func (b UsuarioBuilder) WithSenha(senha string) UsuarioBuilder {
	b.usuario.Senha = senha
	return b
}

func (b UsuarioBuilder) WithCliente(cliente *Cliente) UsuarioBuilder {
	b.usuario.Cliente = cliente
	return b
}

func (b UsuarioBuilder) Build() *Usuario {
	return b.usuario
}
