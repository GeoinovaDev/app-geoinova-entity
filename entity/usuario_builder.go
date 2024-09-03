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

func (b UsuarioBuilder) WithPapel(papel UsuarioPapel) UsuarioBuilder {
	b.usuario.Papel = papel
	return b
}

func (b UsuarioBuilder) WithTelefone(telefone Telefone) UsuarioBuilder {
	b.usuario.Telefone = telefone
	return b
}

func (b UsuarioBuilder) WithDescricao(descricao string) UsuarioBuilder {
	b.usuario.Descricao = descricao
	return b
}

func (b UsuarioBuilder) AddPermissaoAcesso(ativo *Ativo) UsuarioBuilder {
	b.usuario.PermissaoAcesso = append(b.usuario.PermissaoAcesso, ativo)
	return b
}

func (b UsuarioBuilder) WithPermissaoAcesso(permissaoAcesso []*Ativo) UsuarioBuilder {
	b.usuario.PermissaoAcesso = permissaoAcesso
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
