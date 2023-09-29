package entity

type AtivoBuilder struct {
	ativo *Ativo
}

func NewAtivoBuilder(id uint) AtivoBuilder {
	return AtivoBuilder{NewAtivo(id)}
}

func (b AtivoBuilder) WithNome(nome string) AtivoBuilder {
	b.ativo.Nome = nome
	return b
}

func (b AtivoBuilder) WithWkt(wkt string) AtivoBuilder {
	b.ativo.Wkt = wkt
	return b
}

func (b AtivoBuilder) WithCliente(cliente *Cliente) AtivoBuilder {
	b.ativo.Cliente = cliente
	return b
}

func (b AtivoBuilder) WithCategoria(categoria *AtivoCategoria) AtivoBuilder {
	b.ativo.Categoria = categoria
	return b
}

func (b AtivoBuilder) WithStatus(status AtivoStatus) AtivoBuilder {
	b.ativo.Status = status
	return b
}

func (b AtivoBuilder) Build() *Ativo {
	return b.ativo
}
