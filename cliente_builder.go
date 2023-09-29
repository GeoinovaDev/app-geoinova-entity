package entity

type ClienteBuilder struct {
	cliente *Cliente
}

func NewClienteBuilder(id uint) ClienteBuilder {
	return ClienteBuilder{NewCliente(id)}
}

func (b ClienteBuilder) WithNome(nome string) ClienteBuilder {
	b.cliente.Nome = nome
	return b
}

func (b ClienteBuilder) WithStatus(status ClienteStatus) ClienteBuilder {
	b.cliente.Status = status
	return b
}

func (b ClienteBuilder) WithDescricao(descricao string) ClienteBuilder {
	b.cliente.Descricao = descricao
	return b
}

func (b ClienteBuilder) WithUsuarios(usuarios []*Usuario) ClienteBuilder {
	b.cliente.Usuarios = usuarios
	return b
}

func (b ClienteBuilder) AddUsuario(usuario *Usuario) ClienteBuilder {
	b.cliente.Usuarios = append(b.cliente.Usuarios, usuario)
	return b
}

func (b ClienteBuilder) Build() *Cliente {
	return b.cliente
}
